// user-service/datastore.go

package main

import (
	"fmt"
	"golang.org/x/net/context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

// CreateSession creates the main session to our mongodb instance
func CreateSession() (*gorm.DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}

func CreateFrirebase()(*firebase.App, error){
	opt := option.WithCredentialsFile("./invest-ff3f4-firebase-adminsdk-zgkg5-ae79e82ab1.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
