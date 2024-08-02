package WebSocket

import (
	"ChatSpace/initialize"
	"ChatSpace/modules"
	"encoding/json"
	"log"
	"strconv"
)

func messageS(id uint, client *Client) error {
	log.Println("mc get")
	message := modules.Message{}
	result := initialize.DB.Table(`messages`).Model(&modules.Message{}).Where("id = ?", id).Update("state", 1)
	if result.Error != nil {
		log.Println("database err")
	}
	log.Println("update successful")

	result = initialize.DB.Table(`messages`).Model(&modules.Message{}).Where("id = ?", id).Find(&message)
	if result.Error != nil {
		log.Println("database err")
	}
	log.Println("get message ")

	time := message.CreatedAt.Format("2006-01-02 15:04:05")
	sender := strconv.Itoa(int(message.User))
	number := strconv.Itoa(int(message.ID))
	receiver := strconv.Itoa(int(message.Friend))
	msg := Message{message.Type, time, sender, receiver, number, message.Message}
	marshal, err := json.Marshal(msg)
	if err := client.Conn.WriteMessage(1, marshal); err != nil {
		return err
	}
	return err
}

func msgReturn(client *Client) error {
	var message []modules.Message
	result := initialize.DB.Table(`messages`).Model(&modules.Message{}).Where("user2_id = ? AND state = 0", client.ID).Update("state", 1)
	result = initialize.DB.Table(`messages`).Model(&modules.Message{}).Where("user2_id = ? AND state = 0", client.ID).Order("id ASC").Find(&message)

	if result.Error != nil {
		log.Println("database err")
	}
	var msg []Message
	for _, source := range message {
		time := source.CreatedAt.Format("2006-01-02 15:04:05")
		sender := strconv.Itoa(int(source.User))
		number := strconv.Itoa(int(source.ID))
		receiver := strconv.Itoa(int(source.Friend))
		target := Message{
			Type:     source.Type,
			Time:     time,
			Sender:   sender,
			Receiver: receiver,
			MsgId:    number,
			Content:  source.Message,
		}
		msg = append(msg, target)
	}
	marshal, err := json.Marshal(msg)
	if err := client.Conn.WriteMessage(1, marshal); err != nil {
		return err
	}
	return err
}
