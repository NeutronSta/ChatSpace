package Avatar

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var src = "./Files/Default/Avatar.png"

func NewAvatar(id uint) error {
	dest := "./Files/Avatar/" + strconv.Itoa(int(id)) + ".png"
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("无法打开源文件:", err)
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(srcFile)

	destFile, err := os.Create(dest)
	if err != nil {
		log.Println("无法创建目标文件:", err)
		return err
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(destFile)

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		log.Println("无法复制文件内容:", err)
		return err
	}
	return nil
}

/*
func UpdateAvatar(id uint, pic *multipart.FileHeader) error {
	dest := "./Files/Avatar/" + strconv.Itoa(int(id)) + ".png"
	// 打开上传的文件
	file, err := pic.Open()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// 创建目标文件
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Println(err)
		}
	}(out)

	// 将上传的文件内容写入目标文件
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
*/

func GenerateAvatarURL(id uint) string {
	baseURL := "http://114.116.204.137/ChatSpace"
	filePath := "/Avatar/" + strconv.Itoa(int(id)) + ".png"
	return baseURL + filePath
}
