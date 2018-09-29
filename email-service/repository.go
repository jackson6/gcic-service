package main

import (
	pb "github.com/jackson6/gcic-service/email-service/proto/email"
	"log"
	"net/smtp"
)

type Repository interface {
	Send(*pb.User) error
	Close()
}

type EmailRepository struct {
	smpt *smtp.Client
}

func (repo *EmailRepository) Send(user *pb.User) error {
	log.Println("Sending email to:", user.FirstName, user.LastName)
	return nil
}

func (repo *EmailRepository) Close() {
	repo.smpt.Close()
}