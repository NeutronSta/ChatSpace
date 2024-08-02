package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有CORS请求
	},
}

// Client代表一个连接的客户端
type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte // 发送消息的channel
}

// ClientManager管理所有的客户端连接
type ClientManager struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	lock       sync.Mutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (manager *ClientManager) start() {
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
					fmt.Printf("Sent message to %s\n", id)
				default:
					close(conn.Send)
					delete(manager.Clients, id)
				}
			}
		}
	}
}

func handleConnections(manager *ClientManager, w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrade.Upgrade(w, r, nil)
	client := &Client{ID: "someUniqueID", Conn: conn, Send: make(chan []byte)}

	manager.Register <- client

	// 在这里处理客户端发送的消息等
}

func main() {
	manager := NewClientManager()
	go manager.start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(manager, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
