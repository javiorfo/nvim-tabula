package table

import (
	"os"
	"testing"
	"strings"
	"fmt"
)

func TestAddSpaces(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"test", 10, "test      "},
		{"test", 4, "test"},
		{"", 5, "     "},
	}

	for _, test := range tests {
		result := addSpaces(test.input, test.length)
		if result != test.expected {
			t.Errorf("addSpaces(%q, %d) = %q; expected %q", test.input, test.length, result, test.expected)
		}
	}
}

func TestHighlighting(t *testing.T) {
	headers := map[int]Header{
		1: {Name: "Header1", Length: 10},
		2: {Name: "Header2", Length: 15},
	}

	expected := "syn match header1 'Header1' | hi link header1 style |syn match header2 'Header2' | hi link header2 style |"
	result := highlighting(headers, "style")

	if result != expected {
		t.Errorf("highlighting() = %q; expected %q", result, expected)
	}
}

func TestCreateDBeerFileFormat(t *testing.T) {
	destFolder := "./test"
	os.MkdirAll(destFolder, os.ModePerm)
	defer os.RemoveAll(destFolder)

	result := CreateDBeerFileFormat(destFolder)
	if !strings.HasPrefix(result, destFolder) || !strings.HasSuffix(result, ".dbeer") {
		t.Errorf("CreateDBeerFileFormat() = %q; expected to start with %q and end with .dbeer", result, destFolder)
	}
}

func TestWriteToFile(t *testing.T) {
	destFolder := "./test"
	os.MkdirAll(destFolder, os.ModePerm)
	defer os.RemoveAll(destFolder)

	filePath := fmt.Sprintf("%s/testfile.txt", destFolder)
	WriteToFile(filePath, "Hello, World!\n", "This is a test.\n")

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expected := "Hello, World!\nThis is a test.\n"
	if string(content) != expected {
		t.Errorf("WriteToFile() content = %q; expected %q", string(content), expected)
	}
}

func TestGenerate(t *testing.T) {
	destFolder := "./test"
	os.MkdirAll(destFolder, os.ModePerm)
	defer os.RemoveAll(destFolder)

	dbeer := DBeer{
		DestFolder:      destFolder,
		HeaderStyleLink: "style",
		BorderStyle:     1,
		Headers: map[int]Header{
			1: {Name: "Header1", Length: 10},
			2: {Name: "Header2", Length: 15},
		},
		Rows: [][]string{
			{"Row1Col1", "Row1Col2"},
			{"Row2Col1", "Row2Col2"},
		},
	}

	dbeer.Generate()

	filePath := CreateDBeerFileFormat(destFolder)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Expected file %q to be created, but it was not", filePath)
	}
}

