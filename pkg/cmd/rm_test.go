package cmd

import (
	"os"
	"testing"
)

func TestRemoveCommand(t *testing.T) {
	// Create and add a temporary file
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("Test content")
	tmpFile.Close()

	// Add the file to the store
	rootCmd.SetArgs([]string{"add", tmpFile.Name()})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Add command failed: %v", err)
	}

	// Execute the rm command
	rootCmd.SetArgs([]string{"rm", tmpFile.Name()})
	err = rootCmd.Execute()
	if err != nil {
		t.Errorf("Remove command failed: %v", err)
	}
}
