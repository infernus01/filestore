package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestWordCountCommand(t *testing.T) {
	// Create and add a temporary file
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	content := "Go is an open-source programming language."
	tmpFile.WriteString(content)
	tmpFile.Close()

	// Add the file to the store
	rootCmd.SetArgs([]string{"add", tmpFile.Name()})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Add command failed: %v", err)
	}

	// Capture the output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)

	// Execute the wc command
	rootCmd.SetArgs([]string{"wc"})
	err = rootCmd.Execute()
	if err != nil {
		t.Errorf("Word count command failed: %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Errorf("Expected output, got empty string")
	}
}
