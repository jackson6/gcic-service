// chat-service/main.go
package main

import (
	"github.com/micro/go-micro"
	 "github.com/micro/go-plugins/client/http"
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

	srv := micro.NewService(
		micro.Name("go.micro.srv.chat"),
		micro.Version("latest"),
		micro.Client(http.NewClient()),
	)

	srv.Init()


	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	hub := newHub()

	go hub.run()

	http.HandleFunc("/pusher", hub.handleWebSocket)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}