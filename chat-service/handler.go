package main

import (
	"context"
	"encoding/json"
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
	userProto "github.com/jackson6/gcic-service/user-service/proto/user"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func newChatService(session *mgo.Session, hub *Hub, userClient userProto.UserServiceClient) *service{
	return &service{session, hub, userClient}
}

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

	online, err := s.userClient.GetUsers(context.Background(), &userProto.Id{
		Id: clients,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "could not get clients", 500)
	}
	respondJSON(w, http.StatusOK, online.Users)
}