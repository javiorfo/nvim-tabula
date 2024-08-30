package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var logFile *os.File
var logDebug bool

const DATE_FORMAT = "2006/01/02 15:04:05"

func Initialize(logFileName string, logDebugEnabled bool) func() {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error with %s, %v", logFileName, err)
	}

	logger = log.New(logFile, "", 0)
	logDebug = logDebugEnabled

	return func() {
		defer logFile.Close()
	}
}

func Debug(message string) {
	if logDebug {
		timestamp := time.Now().Format(DATE_FORMAT)
		logger.Println(fmt.Sprintf("[DEBUG] [%s] [GO] %s", timestamp, message))
	}
}

func Debugf(format string, a ...any) {
	if logDebug {
		finalFormat := "[DEBUG] [%s] [GO] " + format
		timestamp := time.Now().Format(DATE_FORMAT)
		logger.Println(fmt.Sprintf(finalFormat, timestamp, a))
	}
}

func Error(message string) {
	timestamp := time.Now().Format(DATE_FORMAT)
	logger.Println(fmt.Sprintf("[ERROR] [%s] [GO] %s", timestamp, message))
}

func Errorf(format string, a ...any) {
	finalFormat := "[ERROR] [%s] [GO] " + format
	timestamp := time.Now().Format(DATE_FORMAT)
	logger.Println(fmt.Sprintf(finalFormat, timestamp, a))
}
