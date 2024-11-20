package main

import (
	"fmt"
	"io"
	"os"
)

// TODO: Refactor to take multpile inputs message args
// TODO: file content into empy struct to see!
type io_input struct {
	outpuleFile []string
	message     []string
}

type file_conent struct{}

func main() {
	// Ensure sufficient command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <output_filename> <message>")
		os.Exit(1)
	}

	// Get the filename and message from command-line arguments
	outputFile := os.Args[1]
	message := os.Args[2]

	// Create the file with the given name and write the message
	if err := createFile(outputFile, message); err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	// Read the newly created file and print its contents to standard output
	f, err := os.Open(outputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer f.Close() // Ensure the file is closed after usage

	// Copy the file contents to standard output
	if _, err := io.Copy(os.Stdout, f); err != nil {
		fmt.Println("Error copying file contents:", err)
		os.Exit(1)
	}
}

// createFile creates a file with the specified name and writes the message into it
func createFile(filename, message string) error {
	// Create a new file with the given filename
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write the message to the file
	if _, err := file.WriteString(message); err != nil {
		return err
	}
	return nil
}
