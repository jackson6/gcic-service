package main

import (
	"context"
	"gopkg.in/mgo.v2"
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session *mgo.Session
	userClient userProto.UserServiceClient
}

func (s *service) GetRepo() Repository {
	return &ChatRepository{s.session.Clone()}
}

func (s *service) Messages(ctx context.Context, req *pb.MessageReq, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	data, err := repo.Messages(req)
	if err != nil {
		return err
	}
	res.Messages = data
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Contacts(ctx context.Context, user *pb.User, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	response, err := s.userClient.GetUserReferral(context.Background(), &userProto.User{ReferralCode:user.ReferralCode})
	if err != nil {
		return err
	}

	data, err := repo.Contacts(response.Users, user.Id)
	if err != nil {
		return err
	}

	res.Contacts = data
	return nil
}