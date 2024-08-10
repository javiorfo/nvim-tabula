package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

const DATE_FORMAT = "2006/01/02 15:04:05"

func Initialize(logFileName string) {
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error with %s, %v", logFileName, err)
	}
	defer file.Close()

	infoLogger = log.New(file, "", 0)
	errorLogger = log.New(file, "", 0)
}

func Info(message string) {
	timestamp := time.Now().Format(DATE_FORMAT)
	infoLogger.Println(fmt.Sprintf("[INFO] [%s] %s", timestamp, message))
}

func Error(message string) {
	timestamp := time.Now().Format(DATE_FORMAT)
	errorLogger.Println(fmt.Sprintf("[ERROR] [%s] %s", timestamp, message))
}

func Errorf(format string, a ...any) {
	finalFormat := "[ERROR] [%s] " + format
	timestamp := time.Now().Format(DATE_FORMAT)
	errorLogger.Println(fmt.Sprintf(finalFormat, timestamp, a))
}

func Infof(format string, a ...any) {
	finalFormat := "[INFO] [%s] " + format
	timestamp := time.Now().Format(DATE_FORMAT)
	errorLogger.Println(fmt.Sprintf(finalFormat, timestamp, a))
}
