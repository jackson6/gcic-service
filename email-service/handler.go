package main

import (
	"context"
	"net/smtp"
	pb "github.com/jackson6/gcic-service/email-service/proto/email"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	client *smtp.Client
}

func (s *service) GetRepo() Repository {
	return &EmailRepository{s.client}
}

func (s *service) Send(ctx context.Context, req *pb.User, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Save our user
	err := repo.Send(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}