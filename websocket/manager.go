// websocket/manager.go

package websocket

import "log"

type ClientPool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan ClientMessage
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan ClientMessage),
	}
}

func (pool *ClientPool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("New client connected")
			log.Println("Size of connection pool: ", len(pool.Clients))
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Client disconnected")
			log.Println("Size of connection pool: ", len(pool.Clients))
		case message := <-pool.Broadcast:
			log.Printf("Broadcasting message: %s", message.Body)
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
