package logger

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testlog-*.log")
	if err != nil {
		t.Fatalf("Failed to create temp log file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	cleanup := Initialize(tempFile.Name(), true)
	defer cleanup()

	if logger == nil {
		t.Fatal("Logger was not initialized")
	}
}

func TestDebugDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

    file := "test.log"
	cleanup := Initialize(file, false)
	defer cleanup()
	defer os.Remove(file)

	Debug("This debug message should not appear")

	if buf.Len() > 0 {
		t.Error("Expected no log output, but got some")
	}
}

