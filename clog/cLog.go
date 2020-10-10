package clog

import (
	"log"
	"os"
)

const (
	fatalPrefix = "[FATAL] "
	errorPrefix = "[ERROR] "
	warnPrefix  = "[WARN] "
	infoPrefix  = "[INFO] "
	debugPrefix = "[DEBUG] "
)

var FatalLogger *log.Logger
var ErrorLogger *log.Logger
var WarnLogger *log.Logger
var InfoLogger *log.Logger
var DebugLogger *log.Logger

func InitLog(file string) {
	err := os.Remove(file)
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		panic(err)
	}
	FatalLogger = log.New(logFile, fatalPrefix, log.LstdFlags|log.Lshortfile|log.LUTC)
	ErrorLogger = log.New(logFile, errorPrefix, log.LstdFlags|log.Lshortfile|log.LUTC)
	WarnLogger = log.New(logFile, warnPrefix, log.LstdFlags|log.Lshortfile|log.LUTC)
	InfoLogger = log.New(logFile, infoPrefix, log.LstdFlags|log.Lshortfile|log.LUTC)
	DebugLogger = log.New(logFile, debugPrefix, log.LstdFlags|log.Lshortfile|log.LUTC)
	InfoLogger.Println("Logger is ready!")
}
