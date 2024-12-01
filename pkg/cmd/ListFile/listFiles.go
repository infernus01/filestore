package ListFiles

import (
	"github.com/infernus01/FileServer/pkg/clients"
	"github.com/spf13/cobra"
)

func ListFiles() *cobra.Command {
	listFileCmd := &cobra.Command{

		Use:   "list",
		Short: "lists the files in the store",
		Run: func(cmd *cobra.Command, args []string) {
			clients.ListFiles_c()
		},
	}

	// listFileCmd.Flags().StringVar(&myFlags.help, "help", "h", "anything", "gives help about the subcommand")

	return listFileCmd
}
