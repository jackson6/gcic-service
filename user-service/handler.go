// user-service/handler.go

package main

import (
	"encoding/json"
	"firebase.google.com/go"
	pb "github.com/jackson6/gcic-service/user-service/proto/user"
	paymentProto "github.com/jackson6/gcic-service/payment-service/proto/payment"
	"github.com/micro/go-micro/broker"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"log"
)

const topic = "user.created"

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session *mgo.Session
	firebase *firebase.App
	PubSub broker.Broker
	paymentClient paymentProto.PaymentServiceClient
}

func (s *service) GetRepo() Repository {
	return &UserRepository{s.session.Clone()}
}

// CreateUser - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) Create(ctx context.Context, req *pb.Request, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	charge := &paymentProto.Charge{
		Amount: 30000,
		Description: "Plan Description",
		Currency: "jmd",
	}

	if req.SaveCard {
		customerResponse, err := s.paymentClient.CreateCustomer(context.Background(), &paymentProto.Customer{
			Email: req.User.Email,
		})
		if err != nil {
			return err
		}
		charge.Customer = customerResponse.Id
	}

	paymentResponse, err := s.paymentClient.CreateCharge(context.Background(), charge)
	if err != nil {
		return err
	}

	if paymentResponse != nil {
		// Save our user
		user, err := repo.Create(req.User)
		if err != nil {
			return err
		}
		res.User = user
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	user, err := repo.Get(req)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *service) GetAll(ctx context.Context,req *pb.Request, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	users, err := repo.GetAll()
	if err !=nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	repo := s.GetRepo()
	defer repo.Close()

	client, err := s.firebase.Auth(context.Background())
	if err != nil {
		return err
	}
	_, err = client.VerifyIDToken(context.Background(), req.Token)
	if err != nil {
		return err
	}

	res.Valid = true

	return nil
}

func (s *service) publishEvent(user *pb.User) error {
	// Marshal to JSON string
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Create a broker message
	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	// Publish message to broker
	if err := s.PubSub.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	}

	return nil
}
