package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	eventChan  chan Event
	// GAME POINTER
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		eventChan:  make(chan Event),
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		c.manager.RemoveClient(c)
	}()

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}

			break
		}

		// Decode incoming data into a Event struct (data from frontend coming as Event type object)
		var event Event

		if err := json.Unmarshal(payload, &event); err != nil {
			log.Printf("error marshalling message: %v", err)

			break
		}

		// Route the Event
		if err := c.manager.routeEvent(event, c); err != nil {
			log.Println("error handeling Message: ", err)
		}
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		c.manager.RemoveClient(c)
	}()

	for {
		select {
		case message, ok := <-c.eventChan:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}

				return
			}

			encodedData, err := json.Marshal(message)

			if err != nil {
				log.Println(err)

				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, encodedData); err != nil {
				log.Println(err)
			}

			log.Println("sent message")
		}
	}
}
