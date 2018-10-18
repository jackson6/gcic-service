package main

import (
	pb "github.com/jackson6/gcic-service/partner-service/proto/partner"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"log"
)

func main() {

	// Create new greeter client
	client := pb.NewPartnerServiceClient("gcic.partner", microclient.DefaultClient)


	r, err := client.All(context.Background(), nil)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Partners)

	for _, v := range r.Partners {
		log.Println(v)
	}
}