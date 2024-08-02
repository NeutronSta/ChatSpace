package WebSocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type MC struct {
	Id   uint //message的id
	User uint //用户的id
	//Type string
}

type Client struct {
	ID   uint
	Conn *websocket.Conn
	Send chan MC // 发送消息的channel
}

// ClientManager 管理所有的客户端连接
type ClientManager struct {
	Clients    map[uint]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan MC
	lock       sync.Mutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:    make(map[uint]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan MC),
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn.ID] = conn
			fmt.Println("Added new connection!")
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn.ID]; ok {
				close(conn.Send)
				delete(manager.Clients, conn.ID)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.Broadcast:
			for id, conn := range manager.Clients {
				select {
				case conn.Send <- message:
					fmt.Printf("Sent message to %d\n", id)
				default:
					close(conn.Send)
					delete(manager.Clients, id)
				}
			}
		}
	}
}
