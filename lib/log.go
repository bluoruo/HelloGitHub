package lib

import (
	"log"
	"os"
)

var LogFile = "./run.log"

var Logger *log.Logger

func init() {
	file, err := os.OpenFile(LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常: ", err)
	}
	Logger = log.New(file, "[GitHost]", log.Ldate|log.Ltime)
}
