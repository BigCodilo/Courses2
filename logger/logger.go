package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func SetLoggers() {
	fileInfo, err := os.OpenFile("logger/infoLogger.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("333333333", err)
	}
	Info = log.New(fileInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	fileError, err := os.OpenFile("logger/infoError.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("222222222222222", err)
	}
	Error = log.New(fileError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	fileDebug, err := os.OpenFile("logger/infoDebug.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("1111111111111", err)
	}
	Debug = log.New(fileDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	
}
