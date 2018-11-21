package main

import (
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "invest"
	chatCollection = "messages"
)

type Repository interface {
	Messages(*pb.MessageReq) ([]*Message, error)
	Contacts([]*userProto.User, string) ([]*pb.ContactResp, error)
	Close()
}

type ChatRepository struct {
	session *mgo.Session

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

func (repo *ChatRepository) Contacts(users []*userProto.User, to string) ([]*pb.ContactResp, error) {
	var contactData []*pb.ContactResp

	for _, user := range users {
		message := new(pb.Message)
		count, err := repo.collection().Find(bson.M{"from": user.Id, "to": to, "seen": false}).Count()
		if err != nil && err.Error() != "not found" {
			return nil, err
		}
		err = repo.collection().Find(bson.M{"$or": []bson.M{bson.M{"from": user.Id, "to": to}, bson.M{"from": to, "to": user.Id}}}).Sort("-time", "-_id").One(&message)
		if err != nil && err.Error() != "not found" {
			return nil, err
		}
		contactData = append(contactData, &pb.ContactResp{FirstName: user.FirstName, LastName: user.LastName,
				ProfilePic: user.ProfilePic, Unread: int64(count), Id: user.Id, LastMessage: message})
	}

	return contactData, nil
}

func (repo *ChatRepository) Messages(req *pb.MessageReq) ([]*Message, error) {
	var messages []*Message

	if req.NextTime == 0 {
		err := repo.collection().Find(bson.M{"$or": []bson.M{bson.M{"from": req.From, "to": req.To}, bson.M{"from": req.To, "to": req.From}}}).Sort("-time", "-_id").Limit(6).All(&messages)
		if err != nil {
			return nil, err
		}
	} else {
		err := repo.collection().Find(bson.M{"$or": []bson.M{bson.M{"$or": []bson.M{bson.M{"from": req.From, "to": req.To, "time": bson.M{"$lt": req.NextTime}},
			bson.M{"from": req.To, "to": req.From, "time": bson.M{"$lt": req.NextTime}}}},
			bson.M{"$or": []bson.M{bson.M{"from": req.From, "to": req.To, "time": bson.M{"$lt": req.NextTime}},
				bson.M{"from": req.To, "to": req.From, "time": req.NextTime, "_id": bson.M{"$lt": req.NextId}}}}}}).Sort("-_id").Limit(6).All(&messages)
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
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