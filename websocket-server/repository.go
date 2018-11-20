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
	Get(bson.ObjectId) ([]*Message, error)
	GetById(string) (*Message, error)
	All () ([]*Message, error)
	Save(Message) error
	Update(*Message) error
	ContactList([]*MessageResp) ([]*ContactResp, error)
	Close()
}

type ContactResp struct {
	From string `json:"from"`
	LastMessage *Message `json:"messages"`
	Unread int64 `json:"unread"`
}

type MessageResp struct {
	NextTime int64 `json:"next_time"`
	NextId string `json:"next_id"`
	To       string        `bson:"to"json:"to,omitempty"`
	From     string        `bson:"from"json:"from,omitempty"`
	Messages []*Message `json:"messages"`
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

func (repo *ChatRepository) ContactList(contacts []*MessageResp) ([]*ContactResp, error) {
	var contactData []*ContactResp

	for _, contact := range contacts {
		message := new(Message)
		count, err := repo.collection().Find(bson.M{"from": contact.From, "to": contact.To, "seen": false}).Count()
		if err != nil {
			return nil, err
		}
		err = repo.collection().Find(bson.M{"$or": []bson.M{bson.M{"from": contact.From, "to": contact.To}, bson.M{"from": contact.To, "to": contact.From}}}).Sort("-time", "-_id").One(&message)
		if err != nil {
			return nil, err
		}
		contactData = append(contactData, &ContactResp{Unread: int64(count), From: contact.From, LastMessage:message})
	}

	return contactData, nil
}

func (repo *ChatRepository) All() ([]*Message, error) {
	var messages []*Message

	err := repo.collection().Find(nil).Sort("-time").Limit(50).All(&messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (repo *ChatRepository) GetById(id bson.ObjectId) (*Message, error) {
	message := new(Message)
	err := repo.collection().FindId(id).One(&message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (repo *ChatRepository) Get(message *MessageResp) ([]*Message, error) {
	var messages []*Message

	if message.NextTime == 0 {
		err := repo.collection().Find(bson.M{"$or": []bson.M{bson.M{"from": message.From, "to": message.To}, bson.M{"from": message.To, "to": message.From}}}).Sort("-time", "-_id").Limit(6).All(&messages)
		if err != nil {
			return nil, err
		}
	} else {
		err := repo.collection().Find(bson.M{"$or": []bson.M{bson.M{"$or": []bson.M{bson.M{"from": message.From, "to": message.To, "time": bson.M{"$lt": message.NextTime}},
			bson.M{"from": message.To, "to": message.From, "time": bson.M{"$lt": message.NextTime}}}},
			bson.M{"$or": []bson.M{bson.M{"from": message.From, "to": message.To, "time": bson.M{"$lt": message.NextTime}},
			bson.M{"from": message.To, "to": message.From, "time": message.NextTime, "_id": bson.M{"$lt": message.NextId}}}}}}).Sort("-_id").Limit(6).All(&messages)
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
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