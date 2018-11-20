package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/mgo.v2"
	"os"
)

// CreateSession creates the main session to our mongodb instance
func CreateSession() (*mgo.Session, error) {
	host := os.Getenv("DB_HOST")
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}