package WebSocket

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Type     string `json:"type"`
	Time     string `json:"time"`
	Sender   string `json:"senderID"`
	Receiver string `json:"receiverID"`
	MsgId    string `json:"msgID"`
	Content  string `json:"content"`
}

func Handler(manager *ClientManager, client *Client) error {
	var errReceive, errSend error
	err := msgReturn(client)
	if err != nil {
		return err
	}
	// 使用匿名函数启动两个协程
	go func() {
		for {
			errReceive = receiveMessage(manager, client)
			//log.Println("listening")
			if errReceive != nil {
				return
			}
		}
	}()

	go func() {
		for {
			errSend = sendMessage(client)
			//log.Println("ready to send")
			if errSend != nil {
				return
			}
		}
	}()

	// 在这里等待两个协程结束
	// 你可能需要根据实际情况来处理协程的结束条件
	// 这里只是简单示例
	for {
		if errReceive != nil {
			return errReceive
		}
		if errSend != nil {
			return errSend
		}
	}

}

func receiveMessage(manager *ClientManager, client *Client) error {
	_, p, err := client.Conn.ReadMessage()
	log.Println("Message Get")
	if err != nil {
		return err
	}
	input := string(p)

	var msg Message
	err = json.Unmarshal([]byte(input), &msg)
	if err != nil {
		fmt.Println("解析 JSON 出错:", err)
		return err
	}

	switch msg.Type {
	case "message", "video", "image", "file", "audio":
		{
			messageR(msg.Type, client.ID, msg.Receiver, msg.Content, manager)
		}
	default:
		log.Println("Invalid format")
	}
	return nil
}

func sendMessage(client *Client) error {
	mc := <-client.Send
	if mc.User == client.ID {
		err := messageS(mc.Id, client)
		return err
	}
	return nil
}
