package AddFile_test

import (
	"testing"

	AddFile "github.com/infernus01/FileServer/pkg/cmd/AddFile"
	"github.com/stretchr/testify/assert"
)

func TestAddFilesCommand(t *testing.T) {
	cmd := AddFile.AddFiles()

	// Check if the command is correctly initialized
	assert.Equal(t, "add [files...]", cmd.Use, "Command 'Use' should be 'add [files...]'")
	assert.Equal(t, "Adds files to the store", cmd.Short, "Command 'Short' description mismatch")
	assert.NotNil(t, cmd.Run, "Command 'Run' function should not be nil")

	// Set sample arguments for the command
	cmd.SetArgs([]string{"file1.txt", "file2.txt"})

	// Execute the command
	err := cmd.Execute()
	assert.NoError(t, err, "Command execution should not produce an error")
}
