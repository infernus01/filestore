package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in the store",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/files")
		if err != nil {
			fmt.Printf("Failed to communicate with server: %v\n", err)
			return
		}
		defer resp.Body.Close()
		var files []string
		if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
			fmt.Printf("Failed to parse response: %v\n", err)
			return
		}
		fmt.Println("Files in the store:")
		for _, file := range files {
			fmt.Println("-", file)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
