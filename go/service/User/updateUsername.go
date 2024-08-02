package User

import (
	"ChatSpace/initialize"
	"ChatSpace/middleware"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// UpdateUsername godoc
// @Summary Change user username
// @Description Endpoint to change username
// @Tags user
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Param new_name formData string true "new_name"
// @Success 200 {object} string "Username updated successfully"
// @Failure 400 {object} string "Authentication failed"
// @Failure 400 {object} string "Failed to update username"
// @Router /user/update_username [post]
func UpdateUsername(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")
	newName := c.Request.FormValue("new_name")

	Id, err := middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": strconv.Itoa(int(Id)) + "Authentication failed"})
		log.Println("Authentication failed")
		return
	}
	result := initialize.DB.Table(`users`).Model(&modules.User{}).Where("(id = ?)", Id).Update("name", newName)
	if result.Error != nil {
		log.Println("Failed to update username")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update username"})
	} else {
		log.Println("username updated successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Username updated successfully"})
	}
	return
}
