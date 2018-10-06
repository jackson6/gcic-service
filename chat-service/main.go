// chat-service/main.go
package main

import (
	userService "github.com/jackson6/gcic-service/user-service/proto/user"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"log"
	"os"
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

	srv := web.NewService(
		web.Name("gcic.chat"),
		web.Version("latest"),
	)

	srv.Init()

	hub := newHub()

	userClient := userService.NewUserServiceClient("gcic.user", client.DefaultClient)
	service := service{session, hub, userClient}

	go hub.run()

	// Register Handler
	srv.HandleFunc("/chat", hub.handleWebSocket)
	srv.HandleFunc("/online", service.Online)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}