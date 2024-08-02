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

type FriendRequest struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

// FriendRequestList godoc
// @Summary Get friend request
// @Description Provide every name of user request to be your friend
// @Tags Relationship
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Success 200 {object} string "Get Request List successful"
// @Failure 401 {object} string "without Authorization"
// @Failure 401 {object} string "unexpect Authorization"
// @Router /relationship/friend_request_list [GET]
func FriendRequestList(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")
	Id, err := middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	var user1IDs []uint
	result := initialize.DB.Table(`relationships`).Model(&modules.FriendShip{}).Where("user2_id = ? AND state = ?", Id, 0).Pluck("user1_id", &user1IDs)
	if result.Error != nil {
		log.Println("database err")
	}

	FQL := make([]FriendRequest, 0)
	//userNames := make(map[uint]string)
	for i := range user1IDs {
		name := modules.UserName(user1IDs[i])
		FQL = append(FQL, FriendRequest{Id: user1IDs[i], Username: name, Url: Avatar.GenerateAvatarURL(user1IDs[i])})
		//userNames[user1IDs[id]] = name
	}

	/*
		jsonData, err := json.Marshal(FQL)
		if err != nil {
			fmt.Println("JSON encoding error:", err)
			return
		}

	*/
	c.JSON(http.StatusOK, FQL)
}
