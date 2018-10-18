package main

import (
	"context"
	"encoding/json"
	emailService "github.com/jackson6/gcic-service/email-service/proto/email"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-plugins/broker/nats"
	"log"
)

const topic = "user.created"

func main() {


	// Connect to the remote SMTP server.
	smptClient, err := CreateClient("smtp.gmail.com:465")
	if err != nil {
		log.Fatal(err)
	}

	srv := micro.NewService(
		micro.Name("gcic.email"),
		micro.Version("latest"),
	)

	srv.Init()

	// Get the broker instance using our environment variables
	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to messages on the broker
	_, err = pubsub.Subscribe(topic, func(p broker.Publication) error {
		var user *emailService.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}
		log.Println(user)
		emailClient := emailService.NewEmailServiceClient("go.micro.srv.email", client.DefaultClient)
		_, err := emailClient.Send(context.Background(), user)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	emailService.RegisterEmailServiceHandler(srv.Server(), &service{smptClient})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
