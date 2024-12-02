package cmd

import (
	"bytes"
	"testing"
)

func TestListCommand(t *testing.T) {
	// Capture the output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)

	// Execute the ls command
	rootCmd.SetArgs([]string{"ls"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("List command failed: %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Errorf("Expected output, got empty string")
	}
}
