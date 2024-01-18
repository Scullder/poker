package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients ClientList
	sync.RWMutex
	// handlers are functions that are used to handle Events
	handlers map[string]EventHandler

	// GAMES MAP
}

func NewManager(handlers map[string]EventHandler) *Manager {
	return &Manager{
		clients:  make(ClientList),
		handlers: handlers,
	}
}

func (m *Manager) ServerWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	log.Println("New connection")

	// Begin by upgrading the HTTP request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// GET FREE GAME FROM MAP OR CREATE NEW

	client := NewClient(conn, m)
	// SET CLIENT GAME POINTER
	m.AddClient(client)

	// Start the read / write processes
	go client.ReadMessages()
	go client.WriteMessages()
}

func (m *Manager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	// search hadler by event name
	handler, ok := m.handlers[event.Type]

	if !ok {
		return fmt.Errorf("%v event not registered", event.Type)
	}

	if err := handler(event, c); err != nil {
		return err
	}

	return nil
}
