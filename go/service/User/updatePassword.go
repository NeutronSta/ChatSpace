package User

import (
	"ChatSpace/middleware"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type U struct {
	//Token       string `json:"token"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// UpdatePassword godoc
// @Summary Change user password
// @Description Endpoint to change user password
// @Tags user
// @Accept json
// @Produce json
// @Param Token header string true "Authentication token"
// @Param old_password formData string true "Old password"
// @Param new_password formData string true "New password"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Please try again"
// @Failure 400 {object} string "Password Incorrect"
// @Failure 400 {object} string "Unknown error"
// @Router /user/update_password [post]
func UpdatePassword(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")

	json := U{}
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again"})
		log.Println(err)
		return
	}

	Id, err := middleware.GetID(Authorization)
	/*
		if Authorization == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Without Authorization"})
			return
		}

		parts := strings.Fields(Authorization)
		if len(parts) != 2 || parts[0] != "Bearer" {
			// 处理 Authorization 头部不符合预期的情况
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unexpect Authorization"})
		}

		// 提取 token 值
		Token := parts[1]

		//log.Println(Token)

		//log.Println(json.Token, json.NewPassword, json.OldPassword)

		//name := c.Request.FormValue("name")
		//password := c.Request.FormValue("password")

		name, err := middleware.ParseToken(Token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	*/
	//log.Println(name)
	// 验证密码
	if modules.CheckPassword(Id, json.OldPassword) {
		err := modules.ChangePassword(Id, json.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, "Update Successful")
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password Incorrect"})
	}

}
