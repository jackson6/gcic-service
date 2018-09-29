package main

import (
	"context"
	pb "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"gopkg.in/mgo.v2"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &PaymentRepository{s.session}
}

func (s *service) CreateCharge(ctx context.Context, req *pb.Charge, res *pb.Charge) error {
	repo := s.GetRepo()

	// Save our user
	charge, err := repo.CreateCharge(req)
	if err != nil {
		return err
	}

	res.Id = charge.ID

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) CreateCustomer(ctx context.Context, req *pb.Customer, res *pb.Customer) error {
	repo := s.GetRepo()

	// Save our user
	customer, err := repo.CreateCustomer(req)
	if err != nil {
		return err
	}

	res.Id = customer.ID
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}