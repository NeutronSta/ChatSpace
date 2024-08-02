package User

import (
	"ChatSpace/Files/Avatar"
	"ChatSpace/initialize"
	"ChatSpace/middleware"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
type Response struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
*/

// MyInfo godoc
// @Summary Get my_info
// @Description Provide current user's information
// @Tags user
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Failure 400 {object} string "Can not find your information"
// @Failure 400 {object} string "without Authorization"
// @Failure 400 {object} string "unexpect Authorization"
// @Router /user/my_info [GET]
func MyInfo(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")
	Id, err := middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	u := modules.User{}
	if err := initialize.DB.Where("id = ?", Id).First(&u).Error; err != nil {
		log.Println("database error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not find your information"})
	}

	url := Avatar.GenerateAvatarURL(Id)
	c.JSON(http.StatusOK, gin.H{
		"id":     Id,
		"name":   u.Name,
		"avatar": url,
	})
}
