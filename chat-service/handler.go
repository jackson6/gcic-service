package main

import (
	"context"
	"encoding/json"
	"github.com/jackson6/gcic-service/chat-service/_vendor-20181124171755/gopkg.in/mgo.v2"
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
	"gopkg.in/mgo.v2"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session *mgo.Session
	userClient userProto.UserServiceClient
}

func converter(data interface{}, dataType int) (interface{}, error){
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if dataType == 0 {
		message := new(pb.Message)
		err = json.Unmarshal(jsonData, &message)
		if err != nil {
			return nil, err
		}
		return message, nil
	} else {
		messages := make([]*pb.Message, 0)
		err = json.Unmarshal(jsonData, &messages)
		if err != nil {
			return nil, err
		}
		return messages, nil
	}
}

func (s *service) GetRepo() Repository {
	return &ChatRepository{s.session.Clone()}
}

func (s *service) Messages(ctx context.Context, req *pb.MessageReq, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	data := new(pb.MessageResp)

	messages, err := repo.Messages(req)
	if err != nil {
		return err
	}

	if len(messages) > 0 {
		data.NextTime = messages[len(messages) - 1].Time
		data.NextId = messages[len(messages) - 1].Id.Hex()
	}

	result, err := converter(messages, 2)
	if err != nil {
		return err
	}

	data.Messages = result.([]*pb.Message)

	res.Messages = data
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return nil
}

func (s *service) Contacts(ctx context.Context, user *pb.User, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	response, err := s.userClient.GetUserReferral(context.Background(), &userProto.User{Id: user.Id, ReferralCode:user.ReferralCode})
	if err != nil {
		return err
	}

	response1, err := s.userClient.GetReferred(context.Background(), &userProto.User{Id: user.Id, SponsorId :user.SponsorId})
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