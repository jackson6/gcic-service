// user-service/datastore.go

package main

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

// CreateSession creates the main session to our mongodb instance
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

func CreateFrirebase()(*firebase.App, error){
	opt := option.WithCredentialsFile("./invest-ff3f4-firebase-adminsdk-zgkg5-ae79e82ab1.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
