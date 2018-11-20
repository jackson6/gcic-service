// websocket-server/main.go
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const (
	defaultHost = ":9090"
)

var addr = flag.String("addr", ":9090", "http service address")

type CORSRouterDecorator struct {
	R *http.ServeMux
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}

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
		log.Println("Could not connect to datastore with host %s - %v", host, err)
	}

	repo := &ChatRepository{session}

	hub := newHub(repo)

	go hub.run()

	mux := http.NewServeMux()

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.handleWebSocket(w, r)
	})

	err = http.ListenAndServe(*addr, &CORSRouterDecorator{mux})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}