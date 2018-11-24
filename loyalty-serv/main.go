// benefit-service/main.go
package main

import (
	"github.com/micro/go-micro"
	"log"
	"net/http"
	"os"

	// Import the generated protobuf code
	loyaltyService "github.com/jackson6/gcic-service/loyalty-serv/proto/loyalty"
)

func main() {

	namespace := os.Getenv("NAME_SPACE")

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("gcic.loyalty"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	client := &http.Client{}

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	loyaltyService.RegisterLoyaltyServiceHandler(srv.Server(), &service{client, namespace})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

