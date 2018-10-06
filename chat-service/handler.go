package main

import (
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
	"gopkg.in/mgo.v2"
	"net/http"
)

type service struct {
	session *mgo.Session
	hub *Hub
	userClient userProto.UserServiceClient
}

func (s *service) GetRepo() Repository {
	return &ChatRepository{s.session.Clone(), s.hub}
}

// SaveMessage - we created just one method on our service
func (s *service) Save(msg *pb.Message) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Save our partner
	_, err := repo.Save(msg)
	if err != nil {
		return err
	}

	return nil
}

// SaveMessage - we created just one method on our service
func (s *service) Online(w http.ResponseWriter, r *http.Request) {
	repo := s.GetRepo()
	defer repo.Close()

	// Save our partner
	clients := repo.Online()

	online := s.userClient.GetUsers(context.Background(), &userProtoId{
		id: clients,
	})

	online.

	return nil
}