package Message

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// UploadVideos godoc
// @Summary upload videos
// @Description upload videos type
// @Tags message
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param file formData file true "video to upload"
// @Success 200 {object} map[string]interface{} "id: User ID, message: Upload status message, path: URL to the uploaded video"
// @Failure 400 {object} map[string]string "error: Can not receive the video"
// @Failure 400 {object} map[string]string "error: Failed to create directory"
// @Failure 500 {object} map[string]string "error: Error message for internal server error"
// @Router /upload_videos [post]
func UploadVideos(c *gin.Context) {
	file, err := c.FormFile("video") // 这里的 "file" 是前端通过表单上传文件时的字段名
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not receive the file"})
		return
	}


	// 获取文件大小
	fileSize := file.Size
	// 生成UUID作为文件名
	newFileName := uuid.New().String() + filepath.Ext(file.Filename) // 保留原始文件扩展名

	// 构建保存路径
	savePath := "./Files/Videos/"

	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		// 使用MkdirAll而不是Mkdir，以便创建所有必需的父目录
		if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			log.Println("Failed to create directory")
			return
		}
	}
	Path := savePath + newFileName
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("can not save file")
		return
	}

	// 文件上传成功后，返回文件的保存路径和类型
	c.JSON(http.StatusOK, gin.H{
		"size":     fileSize,
		"filename": newFileName,
		"url":      "http://114.116.204.137/ChatSpace" + Path,
	})
}
