package ReletionShip

import (
	"ChatSpace/Files/Avatar"
	"ChatSpace/initialize"
	"ChatSpace/middleware"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Friend struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

// FriendList godoc
// @Summary Get FriendList
// @Description return all your friend and their information
// @Tags Relationship
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Success 200 {object} string "Get FriendList List successful"
// @Failure 401 {object} string "without Authorization"
// @Failure 401 {object} string "unexpect Authorization"
// @Router /relationship/friend_list [GET]
func FriendList(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")
	Id, err := middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	var user1IDs []uint
	var user2IDs []uint
	result := initialize.DB.Table(`relationships`).Model(&modules.FriendShip{}).Where("user1_id = ? AND state = ?", Id, 1).Pluck("user2_id", &user2IDs)
	if result.Error != nil {
		log.Println("database err")
	}
	result = initialize.DB.Table(`relationships`).Model(&modules.FriendShip{}).Where("user2_id = ? AND state = ?", Id, 1).Pluck("user1_id", &user1IDs)
	if result.Error != nil {
		log.Println("database err")
	}
	friendIDs := append(user1IDs, user2IDs...)

	/*
		var friendIDs [] uint
		result := initialize.DB.Table(`relationships`).Model(&modules.FriendShip{}).Where("(user1_id = ? OR user2_id = ?) AND (state = ?)", Id, Id, 1).Pluck("user2_id", &friendIDs)
		if result.Error != nil {
			log.Println("database err")
		}
	*/

	FL := make([]Friend, 0)
	//userNames := make(map[uint]string)
	for id := range friendIDs {
		name := modules.UserName(friendIDs[id])
		FL = append(FL, Friend{Id: friendIDs[id], Username: name, Url: Avatar.GenerateAvatarURL(friendIDs[id])})
		//userNames[user1IDs[id]] = name
	}

	/*
		jsonData, err := json.Marshal(FL)
		if err != nil {
			fmt.Println("JSON encoding error:", err)
			return
		}

	*/
	c.JSON(http.StatusOK, FL)
}
