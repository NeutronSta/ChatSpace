package WebSocket

import (
	"ChatSpace/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 在这里可以编写自定义的逻辑来检查请求的来源是否合法
		// 返回 true 表示允许连接，返回 false 表示拒绝连接
		// 例如，可以检查请求的 Origin 头部信息
		// 这里是一个简单的示例，允许所有的来源
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(manager *ClientManager, c *gin.Context) {

	token := c.Query("token")
	Authorization := "Bearer " + token
	//log.Println(Authorization)
	Id, err := middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		return
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &Client{ID: Id, Conn: conn, Send: make(chan MC)}
	manager.Register <- client

	//Channel = make(chan MC)
	//log.Println("channel established")
	//GetMessages(Id)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
		}
		manager.Unregister <- client
	}(conn)

	err = Handler(manager, client) //传入要传达的用户id
}

/*
	for {

		messageR, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if err := conn.WriteMessage(messageR, p); err != nil {
			return
		}
	}
*/
