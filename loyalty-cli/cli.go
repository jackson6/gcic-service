package loyalty_cli

import (
	pb "github.com/jackson6/gcic-service/loyalty-service/proto/loyalty"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"log"
)

func main() {

	// Create new greeter client
	client := pb.NewLoyaltyServiceClient("gcic.loyalty", microclient.DefaultClient)


	r, err := client.ListPartners(context.Background(), nil)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Partners)

	for _, v := range r.Partners {
		log.Println(v)
	}
}