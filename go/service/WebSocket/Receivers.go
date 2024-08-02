package WebSocket

import (
	"ChatSpace/initialize"
	"ChatSpace/modules"
	"log"
	"strconv"
)

func messageR(Type string, ID uint, name, text string, manager *ClientManager) {
	var message modules.Message
	id, _ := strconv.ParseUint(name, 10, 0)
	message.Friend = uint(id)
	message.User = ID
	message.Message = text
	message.Type = Type
	result := initialize.DB.Table(`messages`).Create(&message)
	log.Println("Message stored")
	if result.Error != nil {
		log.Println("message written failed")
		//return result.Error
	} else {
		//log.Println("register successful")
		mc := MC{message.ID, message.Friend}
		manager.Broadcast <- mc
		//return nil
	}
}
