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

// AddFriend godoc
// @Summary AddFriend
// @Description request to add a new friend
// @Tags Relationship
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Param Friend_id formData int true "Friend_id"
// @Success 200 {object} string "request send successfully"
// @Failure 400 {object} string "Authentication failed"
// @Failure 400 {object} string "Can not find user"
// @Failure 400 {object} string "Request failed"
// @Router /relationship/add_friend [post]
func AddFriend(c *gin.Context) {
	f := modules.FriendShip{}
	Authorization := c.GetHeader("Authorization")

	friendID, err := strconv.Atoi(c.Request.FormValue("Friend_id"))
	if err != nil {
		// 处理转换错误
		c.JSON(400, gin.H{"error": "Invalid Friend_id"})
		return
	}
	f.Friend = uint(friendID)
	if f.Friend == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not find user"})
		return
	}

	f.User, err = middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		return
	}

	f.State = false

	var count int64
	result := initialize.DB.Table("relationships").
		Where("(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)", f.User, f.Friend, f.Friend, f.User).Count(&count)
	if result.Error != nil {
		// 处理查询错误
		log.Println("database error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request failed"})
	} else {
		if count == 1 {
			// 符合条件的记录存在
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request already sent"})
			log.Println("Request already sent")
			return
		} else if count == 0 {
			// 符合条件的记录不存在，创建请求
			result = initialize.DB.Table(`relationships`).Create(&f)
			if result.Error != nil {
				log.Println("database creat failed")
				c.JSON(http.StatusBadRequest, gin.H{"error": "Request failed"})
			} else {
				log.Println("database creat successful")
				c.JSON(http.StatusOK, "Request send Successful")
			}
			return
		} else {
			log.Println("Data duplication")
			return
		}
	}

}
