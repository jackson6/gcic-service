package main

import (
	"context"
	"fmt"
	"log"
	pb "github.com/jackson6/gcic-service/email-service/proto/email"
	"strings"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	smtpServer *SmtpServer
}

type Mail struct {
	to []string
	from string
	message string
}

func (s *service) GetRepo() Repository {
	// Connect to the remote SMTP server.
	smptClient, err := CreateClient(*s.smtpServer)
	if err != nil {
		log.Fatal(err)
	}
	return &EmailRepository{smptClient}
}

func (s *service) BuildMessage(mail *pb.ContactUs) *Mail {
	data := new(Mail)
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.Email)
	message += fmt.Sprintf("To: %s\r\n", strings.Join([]string{s.smtpServer.addr}, ";"))

	message += fmt.Sprintf("Subject: %s\r\n", "Contact Us")
	message += "\r\n" + "Name: " + mail.Name + "\nEmail: " + mail.Email + "\nCompany: " + mail.Company + "\nPhone: " + mail.Phone + "\n" + mail.Message
	data.message = message
	data.to = []string{s.smtpServer.addr}
	data.from = mail.Email
	return data
}

func (s *service) Contact(ctx context.Context, req *pb.ContactUs, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	message := s.BuildMessage(req)

	// Save our user
	err := repo.Send(message)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Welcome(ctx context.Context, req *pb.User, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}