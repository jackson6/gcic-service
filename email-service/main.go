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
	"os"
)

const topic = "user.created"

type SmtpServer struct {
	host string
	port string
	password string
	addr string
}

func main() {

	host := os.Getenv("SMPT_HOST")
	port := os.Getenv("SMTP_PORT")
	addr := os.Getenv("SMTP_ADDR")
	password := os.Getenv("SMTP_PASS")

	if host == "" || port == "" || addr == "" || password == ""{
		log.Fatal("missing email configuration")
	}

	smtpServer := SmtpServer{host:host, port:port, password:password, addr:addr}

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
	_, err := pubsub.Subscribe(topic, func(p broker.Publication) error {
		var user *emailService.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}
		emailClient := emailService.NewEmailServiceClient("gcic.email", client.DefaultClient)
		_, err := emailClient.Welcome(context.Background(), user)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	emailService.RegisterEmailServiceHandler(srv.Server(), &service{&smtpServer})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
