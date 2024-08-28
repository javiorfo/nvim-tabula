package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var logFile *os.File

const DATE_FORMAT = "2006/01/02 15:04:05"

func Initialize(logFileName string) {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error with %s, %v", logFileName, err)
	}

	logger = log.New(logFile, "", 0)
}

func Debug(message string) {
	timestamp := time.Now().Format(DATE_FORMAT)
	logger.Println(fmt.Sprintf("[DEBUG] [%s] %s", timestamp, message))
	defer logFile.Close()
}

func Debugf(format string, a ...any) {
	finalFormat := "[DEBUG] [%s] " + format
	timestamp := time.Now().Format(DATE_FORMAT)
	logger.Println(fmt.Sprintf(finalFormat, timestamp, a))
	defer logFile.Close()
}

func Error(message string) {
	timestamp := time.Now().Format(DATE_FORMAT)
	logger.Println(fmt.Sprintf("[ERROR] [%s] %s", timestamp, message))
	defer logFile.Close()
}

func Errorf(format string, a ...any) {
	finalFormat := "[ERROR] [%s] " + format
	timestamp := time.Now().Format(DATE_FORMAT)
	logger.Println(fmt.Sprintf(finalFormat, timestamp, a))
	defer logFile.Close()
}
