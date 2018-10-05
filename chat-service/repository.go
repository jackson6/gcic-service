package main

import (
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	"gopkg.in/mgo.v2"
)

const (
	dbName = "invest"
	chatCollection = "messages"
)

type Repository interface {
	Save(message *pb.Message) (*pb.Message, error)
	Close()
}

type ChatRepository struct {
	session *mgo.Session
}

// Create a new user
func (repo *ChatRepository) Save(message *pb.Message)  (*pb.Message, error) {
	err := repo.collection().Insert(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// Close closes the database session after each query has ran.
// Mgo creates a 'master' session on start-up, it's then good practice
// to copy a new session for each request that's made. This means that
// each request has its own database session. This is safer and more efficient,
// as under the hood each session has its own database socket and error handling.
// Using one main database socket means requests having to wait for that session.
// I.e this approach avoids locking and allows for requests to be processed concurrently. Nice!
// But... it does mean we need to ensure each session is closed on completion. Otherwise
// you'll likely build up loads of dud connections and hit a connection limit. Not nice!
func (repo *ChatRepository) Close() {
	repo.session.Close()
}

func (repo *ChatRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(chatCollection)
}