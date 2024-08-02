package User

import (
	"ChatSpace/middleware"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary Login
// @Description User login
// @Tags user
// @Accept json
// @Produce json
// @Param name formData string true "name"
// @Param password formData string true "Password"
// @Success 200 {string} string "Login successful"
// @Failure 400 {string} string "User does not exist"
// @Failure 400 {string} string "Please enter username"
// @Failure 400 {string} string "Please enter password"
// @Failure 401 {string} string "Authentication failed"
// @Router /user/login [post]
func Login(c *gin.Context) {
	// 解析请求中的用户名和密码
	//name := c.PostForm("name")
	//password := c.PostForm("password")

	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")

	if name == "" {
		c.JSON(http.StatusBadRequest, "Please enter username")
		return
	}

	if password == "" {
		c.JSON(http.StatusBadRequest, "Please enter password")
		return
	}

	// 检查用户id是否存在
	id := modules.UserId(name)
	if !modules.CheckUserExists(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	// 验证密码
	if modules.CheckPassword(id, password) {
		id := modules.UserId(name)
		token := middleware.NewToken(id)
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}
