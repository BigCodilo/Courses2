package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func SetLoggers() error{
	fileInfo, err := os.OpenFile("logger/LogInfo.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		Info.SetOutput(os.Stdout)
		return err
	}
	Info = log.New(fileInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	fileError, err := os.OpenFile("logger/LogError.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		Error.SetOutput(os.Stdout)
		return err
	}
	Error = log.New(fileError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	fileDebug, err := os.OpenFile("logger/LogGebug.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil{
		Debug.SetOutput(os.Stdout)
		return err
	}
	Debug = log.New(fileDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}
