/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-21 17:21:50
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 18:14:45
 * @FilePath: /go-mall/pkg/util/logger.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	// 获取输出的工作目录
	src, _ := setOutPutFile()
	if LogrusObj != nil {
		LogrusObj.Out = src
		return

	}
	// 为空就实例化
	logger := logrus.New()
	logger.Out = src                   // 设置输出
	logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	LogrusObj = logger

}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	// os.Getwd()获取当前工作目录 go-mall
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}

	// logFilePath是否存在，不存在就创建
	// os.Stat(logFilePath) 返回一个 os.FileInfo 对象和一个错误值 err
	// 如果路径存在，err 将为 nil；如果路径不存在，err 将包含一个错误
	// 通常是 os.ErrNotExist
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// fileName是否存在，不存在就创建
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(fileName, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return src, nil
}
