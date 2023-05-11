package log

import (
	"log"
	"os"
)

func SetupLogger(filePath string) {
	// filePath为日志文件存放地址
	logFileLocation, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(logFileLocation)
	log.Println("hello world!")
}
