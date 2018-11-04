// user-service/main.go
package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"os"

	paymentService "github.com/jackson6/gcic-service/payment-service/proto/payment"
	planService "github.com/jackson6/gcic-service/plan-service/proto/plan"
	userService "github.com/jackson6/gcic-service/user-service/proto/user"
)

func main() {

	db, err := CreateSession()


	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer db.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore - %v", err)
	}

	repo := &UserRepository{db}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&userService.User{})

	firebase, err := CreateFrirebase()
	if err != nil {

		// We're wrapping the error returned from our CreateFirebase
		// here to add some context to the error.
		log.Panic("Could not connect to firebase", err)
	}

	// Create a new service. Optionally include some options here.
	srv := k8s.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("user"),
		micro.Version("latest"),
	)

	paymentClient := paymentService.NewPaymentServiceClient("payment", srv.Client())
	planClient := planService.NewPlanServiceClient("plan", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	pubsub := srv.Server().Options().Broker

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	userService.RegisterUserServiceHandler(srv.Server(), &service{repo, firebase, pubsub, paymentClient, planClient})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {

		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserServiceClient("gcic.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

