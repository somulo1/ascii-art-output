package main

import (
	"ascii-art-output/functions"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {

	outputFileName := "output.txt" // Specify the default value for the flag as specified in main.go

	// Check if the file exists
	_, err := os.Stat(outputFileName)
	if err != nil {
		t.Fatalf("File %s does not exist", outputFileName)
	}

	//Read the file if it exists
	expectedContent, err := os.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("Failed to read expected output file '%s': %v", outputFileName, err)
	}

	expectedContentStr := string(expectedContent)

	//Read the default banner file specified in main.go
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	var fileLine []string

	fileLine = strings.Split(string(file), "\n")

	//Specify a string input that can be entered by the user
	input := "hello"

	result := functions.AsciiArt(input, fileLine)
	if result != expectedContentStr {
		t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
	}
}
