// user-service/handler.go

package main

import (
	"encoding/json"
	"firebase.google.com/go"
	paymentProto "github.com/jackson6/gcic-service/payment-service/proto/payment"
	planProto "github.com/jackson6/gcic-service/plan-service/proto/plan"
	pb "github.com/jackson6/gcic-service/user-service/proto/user"
	"github.com/micro/go-micro/broker"
	"golang.org/x/net/context"
	"log"
	"time"
)

const topic = "user.created"

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo Repository
	firebase *firebase.App
	PubSub broker.Broker
	paymentClient paymentProto.PaymentServiceClient
	planClient planProto.PlanServiceClient
}

// CreateUser - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) Create(ctx context.Context, req *pb.Request, res *pb.Response) error {
	charge := &paymentProto.Charge{
		Amount: 30000,
		Description: "Plan Description",
		Currency: "jmd",
		Token: req.Token,
		UserId: req.User.UserId,
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

	date := time.Now()
	expDate := date.AddDate(0, 1, 0)

	req.User.MembershipExp = expDate.Unix()

	if paymentResponse != nil {
		// Save our user
		err := s.repo.Create(req.User)
		if err != nil {
			return err
		}
	}
	res.User = req.User
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Update(ctx context.Context, req *pb.User, res *pb.Response) error {
	err := s.repo.Update(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(ctx context.Context, req *pb.User, res *pb.Response) error {
	err := s.repo.Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		log.Println("error", err)
		return err
	}
	res.User = user
	return nil
}

func (s *service) GetUserReferral(ctx context.Context, req *pb.User, res *pb.Response) error {
		users, err := s.repo.GetReferrals(req.ReferralCode)
	if err != nil {
		log.Println("error", err)
		return err
	}
	res.Users = users
	return nil
}

func (s *service) GetByEmail(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *service) GetByMemberId (ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.GetByMemberId(req.MemberId)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *service) GetUsers(ctx context.Context, req *pb.Id, res *pb.Response) error {
	users, err := s.repo.GetUsers(req.Id)
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *service) All(ctx context.Context,req *pb.Request, res *pb.Response) error {
	users, err := s.repo.All()
	if err !=nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
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

