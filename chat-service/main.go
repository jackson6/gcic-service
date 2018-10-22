// chat-service/main.go
package main

import (
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

	repo := &ChatRepository{session}

	srv := web.NewService(
		web.Name("gcic.chat"),
		web.Version("latest"),
	)

	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}

	hub := newHub(repo)

	go hub.run()

	// Register Handler
	srv.HandleFunc("/chat", hub.handleWebSocket)
	srv.HandleFunc("/online", hub.online)
	srv.HandleFunc("/message", hub.messages)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}