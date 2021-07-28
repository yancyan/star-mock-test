package requests

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

const (
	LogFilePath = "logs"
	LogFileName = "kingdee-test.log"
)

var Logger = logrus.New()

func init() {
	// 设置日志格式为json格式
	Logger.SetFormatter(&logrus.TextFormatter{})
	Logger.SetLevel(logrus.InfoLevel)

	Logger.SetOutput(os.Stdout)
	fileName := path.Join(LogFilePath, LogFileName)

	Logger.Infof("open-file right %s", os.FileMode(0644).String())
	//// If the file doesn't exist, create it, or append to the file
	//f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		Logger.Fatal("log init err", err)
	}

	// 创建记录日志的文件
	//gin.DefaultWriter = io.MultiWriter(src, os.Stdout)
	Logger.SetOutput(io.MultiWriter(src, os.Stdout))

	//Logger.WithFields(logrus.Fields{
	//	"omg":    true,
	//	"number": 122,
	//}).Warn("The group's number increased tremendously!")

	//Logger.AddHook(&CustomHook{})

}

type CustomHook struct {
	//	type Hook interface {
	//	Levels() []Level
	//	Fire(*Entry) error
	//}
}

func (ch *CustomHook) Fire(entry *logrus.Entry) error {
	Logger.Info(entry.Message, "<<< -------- customHook.")
	return nil
}
func (ch *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
