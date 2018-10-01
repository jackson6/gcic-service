package main

import (
	paymentService "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"log"
	"os"
)

const (
	defaultHost = "mongo:27017"
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

	srv := k8s.NewService(
		micro.Name("gcic.payment"),
		micro.Version("latest"),
	)

	srv.Init()

	paymentService.RegisterPaymentServiceHandler(srv.Server(), &service{ session})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
