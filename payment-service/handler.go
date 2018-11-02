package main

import (
	"context"
	pb "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"log"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo Repository
}

func (s *service) CreateCharge(ctx context.Context, req *pb.Charge, res *pb.Response) error {
	log.Println(req)

	// Save our user
	charge, err := s.repo.CreateCharge(req)
	if err != nil {
		return err
	}

	res.Charge.Id = charge.ID

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) CreateCustomer(ctx context.Context, req *pb.Customer, res *pb.Response) error {
	// Save our user
	customer, err := s.repo.CreateCustomer(req)
	if err != nil {
		return err
	}

	res.Customer.Id = customer.ID
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}


func (s *service) History(ctx context.Context, req *pb.Transaction, res *pb.Response) error {
	// Save our userCustomer
	transactions, err := s.repo.GetHistory(req)
	if err != nil {
		return err
	}

	res.Transactions = transactions
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}