/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-18 11:25:55
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-18 16:02:04
 * @FilePath: /go-mall/service/upload.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"go-mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarTolocalStatic(file multipart.File, userId uint, userName string) (filepath string, err error) {
	bId := strconv.Itoa(int(userId)) // 路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" // todo: 把file的后缀提取出来
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}

	return "user" + bId + "/" + userName + ".jpg", err
}

// 判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
