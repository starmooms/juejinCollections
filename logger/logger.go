package logger

import (
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

const isDebug = true

var Logger *logrus.Logger

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

	src, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	Logger = logrus.New()
	formatConfig := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	if isDebug {
		Logger.SetOutput(os.Stdout)
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.Out = src
		Logger.SetLevel(logrus.InfoLevel)
		formatConfig.PrettyPrint = true
		// Logger.AddHook(Hook)
	}
	Logger.SetFormatter(formatConfig)

	Logger.Info("Logger Start")
	// Logger.Fatal("close")
	// Logger.Panic("panic close")
}
