package logger

import (
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

const isDebug = true

var Logger *logrus.Logger
var src *os.File

func GetLog() *logrus.Logger {
	return Logger
}

func GetFile() *os.File {
	return src
}

func init() {
	now := time.Now()
	logFilePath := ""

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logFilePath = path.Join(dir, "logs")

	// 生成目录
	err = os.MkdirAll(logFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logFilePath = path.Join(logFilePath, now.Format("2006-01-02")+".log")

	// // 检查生成文件
	// if _, err = os.Stat(logFilePath); err != nil {
	// 	if _, err = os.Create(logFilePath); err != nil {
	// 		panic(err)
	// 	}
	// }

	src, err = os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	Logger = logrus.New()
	formatConfig := &MyFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		enableColor:     false,
	}
	if isDebug {
		Logger.SetOutput(os.Stdout)
		Logger.SetLevel(logrus.DebugLevel)
		formatConfig.enableColor = true
		formatConfig.errorWrite = src
	} else {
		Logger.Out = src
		Logger.SetLevel(logrus.InfoLevel)
	}
	Logger.SetFormatter(formatConfig)

	Logger.Info("Logger Start")
}
