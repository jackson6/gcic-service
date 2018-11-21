package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "invest"
	chatCollection = "messages"
)

type Repository interface {
	GetById(string) (*Message, error)
	Save(Message) error
	Update(*Message) error
	Close()
}

type Message struct {
	Id       bson.ObjectId `bson:"_id"json:"id"`
	Event 	 string 	   `bson:"event"json:"event"`
	To       string        `bson:"to"json:"to"`
	From     string        `bson:"from"json:"from"`
	Time     int64         `bson:"time"json:"time"`
	Text     string        `bson:"text"json:"text"`
	Seen     bool          `bson:"seen"json:"seen"`
	Received bool          `bson:"received"json:"received"`
}

type ChatRepository struct {
	session *mgo.Session
}

func (repo *ChatRepository) GetById(id bson.ObjectId) (*Message, error) {
	message := new(Message)
	err := repo.collection().FindId(id).One(&message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (repo *ChatRepository) Save(message *Message) error {
	err := repo.collection().Insert(message)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ChatRepository) Update(message *Message) error {
	err := repo.collection().UpdateId(message.Id, message)
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