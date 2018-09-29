// user-service/main.go
package main

import (
	"context"
	"errors"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"os"

	// Import the generated protobuf code
	userService "github.com/jackson6/gcic-service/user-service/proto/user"
	paymentService "github.com/jackson6/gcic-service/payment-service/proto/payment"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)


	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	firebase, err := CreateFrirebase()
	if err != nil {

		// We're wrapping the error returned from our CreateFirebase
		// here to add some context to the error.
		log.Panic("Could not connect to firebase", err)
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("gcic.user"),
		micro.Version("latest"),
	)

	paymentClient := paymentService.NewPaymentServiceClient("gcic.payment", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	pubsub := srv.Server().Options().Broker

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	userService.RegisterUserServiceHandler(srv.Server(), &service{session, firebase, pubsub, paymentClient})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
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

