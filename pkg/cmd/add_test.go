package cmd

import (
	"os"
	"testing"
)

func TestAddCommand(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("Hello World")
	tmpFile.Close()

	// Execute the add command
	rootCmd.SetArgs([]string{"add", tmpFile.Name()})
	err = rootCmd.Execute()
	if err != nil {
		t.Errorf("Add command failed: %v", err)
	}
}
