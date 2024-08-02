package ReletionShip

import (
	"ChatSpace/initialize"
	"ChatSpace/middleware"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FriendVerify godoc
// @Summary Verify request
// @Description agree your friend request
// @Tags Relationship
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Param Friend_id formData int true "Friend_id"
// @Param Verify(true/false) formData string true "Verify"
// @Success 200 {object} string "Record updated successfully"
// @Failure 401 {object} string "Authentication failed"
// @Failure 400 {object} string "Can not find user"
// @Failure 400 {object} string "Failed to update record"
// @Router /relationship/verify [post]
func FriendVerify(c *gin.Context) {
	f := modules.FriendShip{}
	Authorization := c.GetHeader("Authorization")
	friendID, _ := strconv.Atoi(c.Request.FormValue("Friend_id"))
	f.User = uint(friendID)

	if f.User == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not find user"})
		log.Println("Can not find user")
		return
	}
	verify, err := strconv.ParseBool(c.Request.FormValue("Verify"))

	f.Friend, err = middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		log.Println("Authentication failed")
		return
	}

	if !verify {
		result := initialize.DB.Table(`relationships`).Where("(user1_id = ? AND user2_id = ?)", f.User, f.Friend).Delete(&modules.FriendShip{})
		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete"})
		} else {
			log.Println("The request was denied")
			c.JSON(http.StatusOK, gin.H{"message": "Record Deleted"})
		}
		return
	}

	//f.State = true

	result := initialize.DB.Table(`relationships`).Model(&modules.FriendShip{}).Where("(user1_id = ? AND user2_id = ?)", f.User, f.Friend).Update("state", 1)

	if result.Error != nil {
		log.Println("Failed to update record")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update record"})
	} else {
		log.Println("Record updated successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
	}
	return
}
