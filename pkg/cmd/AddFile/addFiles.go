package AddFile

import (
	"fmt"

	"github.com/infernus01/FileServer/pkg/clients"
	"github.com/spf13/cobra"
)

func AddFiles() *cobra.Command {
	var files []string

	addFileCmd := &cobra.Command{
		Use:   "add [files...]",
		Short: "Adds files to the store",
		Run: func(cmd *cobra.Command, args []string) {
			allFiles := append(args, files...)
			if len(allFiles) == 0 {
				fmt.Println("Please provide at least one file either as an argument or using the --files flag")
				return
			}
			clients.AddFiles_c(allFiles)
		},
	}

	addFileCmd.Flags().StringSliceVarP(&files, "files", "f", nil, "Files to add")

	return addFileCmd
}
