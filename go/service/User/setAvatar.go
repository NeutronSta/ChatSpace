package User

import (
	"ChatSpace/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SetAvatar godoc
// @Summary Change user avatar
// @Description upload user avatar in png form
// @Tags user
// @Accept  multipart/form-data
// @Produce json
// @Param Token header string true "Authentication token"
// @Param avatar formData file true "picture"
// @Success 200 {object} string "Update avatar successfully"
// @Failure 400 {object} string "Can not receive the file"
// @Failure 400 {object} string "without Authorization"
// @Failure 400 {object} string "unexpect Authorization"
// @Failure 400 {object} string "Update avatar failed"
// @Router /user/avatar [post]
func SetAvatar(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not receive the file"})
		return
	}
	Id, err := middleware.GetID(Authorization)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//err = Avatar.UpdateAvatar(Id, file)
	dst := "./Files/Avatar/" + strconv.Itoa(int(Id)) + ".png"
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Upload file failed"})
		return
	}
	// 保存文件到指定路径

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功"})
	return
}
