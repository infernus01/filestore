package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [filename]",
	Short: "Remove a file from the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		url := fmt.Sprintf("http://localhost:8080/files/%s", filename)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Failed to communicate with server: %v\n", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Removed file %s from the store.\n", filename)
		} else {
			fmt.Printf("Failed to remove file %s: %s\n", filename, resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
