package cmd

import (
	"os"
	"testing"
)

func TestUpdateCommand(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("Initial content")
	tmpFile.Close()

	// Add the file to the store
	rootCmd.SetArgs([]string{"add", tmpFile.Name()})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Add command failed: %v", err)
	}

	// Update the file content
	os.WriteFile(tmpFile.Name(), []byte("Updated content"), 0644)

	// Execute the update command
	rootCmd.SetArgs([]string{"update", tmpFile.Name()})
	err = rootCmd.Execute()
	if err != nil {
		t.Errorf("Update command failed: %v", err)
	}
}
