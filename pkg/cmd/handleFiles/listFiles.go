package handlefiles

import (
	"github.com/infernus01/fileService/pkg/clients"
	"github.com/spf13/cobra"
)

func ListFiles() *cobra.Command {
	listFileCmd := &cobra.Command{

		Use:   "list",
		Short: "lists the files in the store",
		Long: `longer version
j89u
3
			87u8
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			clients.ListFiles()
		},
	}

	// listFileCmd.Flags().StringVar(&myFlags.help, "help", "h", "anything", "gives help about the subcommand")

	return listFileCmd
}
