package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	pb "github.com/jackson6/gcic-service/chat-service/proto/chat"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    map[string]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
	session	   *mgo.Session
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (s *Hub) GetRepo() Repository {
	return &ChatRepository{s.session.Clone()}
}

// SaveMessage - we created just one method on our service
func (s *Hub) save(msg *pb.Message) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Save our partner
	_, err := repo.Save(msg)
	if err != nil {
		return err
	}

	return nil
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		case message := <-hub.broadcast:
			hub.onMessage(message)
		}
	}
}

func (hub *Hub) broadcastMessage(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)
	for _, c := range hub.clients {
		if c != ignore {
			c.outbound <- data
		}
	}
}

func (hub *Hub) send(message interface{}, client *Client) {
	data, _ := json.Marshal(message)
	client.outbound <- data
}

func (hub *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	id := r.URL.Query().Get("id")
	if err != nil {
		log.Println(err)
		http.Error(w, "could not upgrade", http.StatusInternalServerError)
		return
	}
	client := newClient(hub, socket, id)
	hub.register <- client

	go client.write()
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("client connected: ", client.socket.RemoteAddr())

	// Make new client
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	hub.clients[client.id] = client
}

func (hub *Hub) onMessage(message []byte) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	var msg pb.Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	log.Println("new message to: ", msg.To)

	to := hub.clients[msg.To]
	hub.send(msg, to)
	go hub.save(&msg)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.socket.RemoteAddr())

	client.close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	// Delete client from list
	delete(hub.clients, client.id)
}