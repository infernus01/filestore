package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "Get total word count of all files in the store",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/wc")
		if err != nil {
			fmt.Printf("Failed to communicate with server: %v\n", err)
			return
		}
		defer resp.Body.Close()
		var result map[string]int
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Printf("Failed to parse response: %v\n", err)
			return
		}
		fmt.Printf("Total word count: %d\n", result["word_count"])
	},
}

func init() {
	rootCmd.AddCommand(wcCmd)
}
