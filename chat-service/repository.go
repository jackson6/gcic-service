package main

import (
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "invest"
	chatCollection = "messages"
)

type Repository interface {
	Get(*pb.User) (*pb.Message, error)
	Save(*pb.Message) error
	Close()
}

type ChatRepository struct {
	session *mgo.Session
}



func (repo *ChatRepository) Get(message *pb.Message) ([]*pb.Message, error) {
	var messages []*pb.Message
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(bson.M{"from": message.From, "to": message.To}).All(&messages)
	return messages, err
}

func (repo *ChatRepository) Save(messagge *pb.Message) error {
	err := repo.collection().Insert(messagge)
	if err != nil {
		return err
	}
	return nil
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