// benefit-service/main.go
package main

import (
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"log"
	"os"

	// Import the generated protobuf code
	benefitService "github.com/jackson6/gcic-service/benefit-service/proto/benefit"
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

	// Create a new service. Optionally include some options here.
	srv := k8s.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("benefit"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	benefitService.RegisterBenefitServiceHandler(srv.Server(), &service{session})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

