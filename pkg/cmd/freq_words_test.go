package cmd

import (
    "bytes"
    "os"
    "testing"
)

func TestFreqWordsCommand(t *testing.T) {
    // Create and add a temporary file
    tmpFile, err := os.CreateTemp("", "testfile*.txt")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer os.Remove(tmpFile.Name())
    content := "Go Go Go Golang"
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

    // Execute the freq-words command
    rootCmd.SetArgs([]string{"freq-words", "-n", "2", "--order", "dsc"})
    err = rootCmd.Execute()
    if err != nil {
        t.Errorf("Freq-words command failed: %v", err)
    }

    output := buf.String()
    if output == "" {
        t.Errorf("Expected output, got empty string")
    }
}
