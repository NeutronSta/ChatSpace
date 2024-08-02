package User

import (
	"ChatSpace/Files/Avatar"
	"ChatSpace/modules"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// RegisterUser godoc
// @Summary register user
// @Description register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param name formData string true "name"
// @Param password formData string true "Password"
// @Success 200 {string} string "User registered successfully"
// @Failure 400 {string} string "Please enter username"
// @Failure 400 {string} string "Please enter password"
// @Failure 400 {string} string "User already exists"
// @Router /user/register [post]
func RegisterUser(c *gin.Context) {

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

	// 检查当前请求中的用户id是否已存在
	id := modules.UserId(name)
	if modules.CheckUserExists(id) {
		c.JSON(http.StatusBadRequest, "User already exists")
		return
	}

	// 用户不存在，创建用户
	if modules.CreateUser(name, password) {
		c.JSON(http.StatusOK, "User registered successfully")
		err := Avatar.NewAvatar(id)
		if err != nil {
			log.Println("default avatar failed")
		} else {
			log.Println("avatar create successfully")
		}
	} else {
		c.JSON(http.StatusBadRequest, "Pleaase try again")
	}
}
