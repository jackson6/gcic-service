package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (hub *Hub) send(message *pb.Message, client *Client) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	client.outbound <- data
	go hub.repo.Save(message)
}

func (hub *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	img := r.URL.Query().Get("img")
	client := newClient(hub, socket, id, name, img)
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
	hub.broadcastMessage("connected:"+strconv.Itoa(len(hub.clients)), nil)
}

func (hub *Hub) onMessage(message []byte) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	log.Println(string(message))

	var msg pb.Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	log.Println("new message to: ", msg.To)

	to := hub.clients[msg.To]
	hub.send(&msg, to)
}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("client disconnected: ", client.socket.RemoteAddr())

	client.close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	// Delete client from list
	delete(hub.clients, client.id)
}

func (hub *Hub) online(w http.ResponseWriter, r *http.Request) {
	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: hub.clients,
	}
	log.Println(response)
	RespondJSON(w, http.StatusOK, response)
}

func (hub *Hub) messages(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	messages, err := hub.repo.Get(&pb.Message{From:from, To:to})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	response := HttpResponse{
		ResultCode: 200,
		CodeContent: "Success",
		Data: messages,
	}
	RespondJSON(w, http.StatusOK, response)
}