package main

import (
	chatService "github.com/jackson6/gcic-service/chat-service/proto/chat"
	userService "github.com/jackson6/gcic-service/user-service/proto/user"
	k8s "github.com/micro/kubernetes/go/micro"
	"github.com/micro/go-micro"
	"log"
)

func main() {

	session, err := CreateSession()


	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore - %v", err)
	}

	srv := k8s.NewService(
		micro.Name("payment"),
		micro.Version("latest"),
	)

	userClient := userService.NewUserServiceClient("payment", srv.Client())

	srv.Init()

	chatService.RegisterChatServiceHandler(srv.Server(), &service{ session, userClient})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
