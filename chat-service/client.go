package main

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	outbound chan []byte
}

func newClient(hub *Hub, socket *websocket.Conn, id string) *Client {
	return &Client{
		hub:      hub,
		id:       id,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (client *Client) read() {
	defer func() {
		client.close()
	}()
	client.socket.SetReadLimit(maxMessageSize)
	client.socket.SetReadDeadline(time.Now().Add(pongWait))
	client.socket.SetPongHandler(func(string) error { client.socket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := client.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		client.hub.broadcast <- message
	}
}

func (client *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.close()
	}()
	for {
		select {
		case data, ok := <-client.outbound:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := client.socket.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(data)

			// Add queued chat messages to the current websocket message.
			n := len(client.outbound)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.outbound)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client Client) close() {
	client.hub.unregister <- &client
	client.socket.Close()
	close(client.outbound)
}