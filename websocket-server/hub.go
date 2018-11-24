package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"reflect"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func updateBuilder(old, new interface{}) interface{} {
	oldVal := reflect.ValueOf(old).Elem()
	newVal := reflect.ValueOf(new).Elem()

	for i := 0; i < oldVal.NumField(); i++ {
		for j := 0; j < newVal.NumField(); j++ {
			if oldVal.Type().Field(i).Name == newVal.Type().Field(j).Name {
				if newVal.Field(j).Interface() != nil {
					oldVal.Field(i).Set(newVal.Field(j))
				}
			}
		}
	}
	return oldVal.Interface()
}

type Hub struct {
	clients    map[string]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	repo 	   *ChatRepository
	mutex      *sync.Mutex
}

func newHub(repo *ChatRepository) *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		repo: 		repo,
		mutex:      &sync.Mutex{},
	}
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

func (hub *Hub) send(message *Message, client *Client) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	if client != nil {
		client.outbound <- data
	}
}

func (hub *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}

	id := r.URL.Query().Get("id")

	client := newClient(hub, socket, id)

	hub.register <- client

	go client.write()
	go client.read()
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("client connected: ", client.socket.RemoteAddr())

	// Make new client
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	hub.clients[client.id] = client
	hub.send(&Message{Event: "Connected", Text:"user connected"}, client)
}

func (hub *Hub) onMessage(message []byte) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	var msg Message

	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	switch msg.Event {
	case "Message":
		msg.Id = bson.NewObjectId()
		to := hub.clients[msg.To]
		hub.send(&msg, to)
		go hub.repo.Save(&msg)
		break
	case "Seen":
		data, err := hub.repo.GetById(msg.Id)
		if err != nil {
			log.Println(err)
			break
		}
		data.Seen = true
		to := hub.clients[msg.From]
		hub.send(&msg, to)
		go hub.repo.Update(data)
		break
	case "Received":
		data, err := hub.repo.GetById(msg.Id)
		if err != nil {
			log.Println(err)
			break
		}
		data.Received = true
		to := hub.clients[msg.From]
		hub.send(&msg, to)
		go hub.repo.Update(data)
	case "Typing":
		to := hub.clients[msg.To]
		hub.send(&msg, to)
		break
	}
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	// Delete client from list
	delete(hub.clients, client.id)
	close(client.outbound)
}