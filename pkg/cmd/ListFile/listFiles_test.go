package ListFiles_test

import (
	"testing"

	ListFiles "github.com/infernus01/FileServer/pkg/cmd/ListFile"
	"github.com/stretchr/testify/assert"
)

func TestListFilesCommand(t *testing.T) {
	cmd := ListFiles.ListFiles()

	// Check if the command is correctly initialized
	assert.Equal(t, "list", cmd.Use, "Command 'Use' should be 'list'")
	assert.Equal(t, "lists the files in the store", cmd.Short, "Command 'Short' description mismatch")
	assert.NotNil(t, cmd.Run, "Command 'Run' function should not be nil")

	// Execute the command
	err := cmd.Execute()
	assert.NoError(t, err, "Command execution should not produce an error")

}
