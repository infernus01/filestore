package cmd

import (
	"testing"
)

func TestRootCommand(t *testing.T) {
	rootCmd.SetArgs([]string{})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Root command failed: %v", err)
	}
}
