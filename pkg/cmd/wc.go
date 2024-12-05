package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "Count the number of words in all files stored on the server",
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/files"

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Failed to fetch file list: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to retrieve file list: %s", resp.Status)
			return
		}

		var fileList []string
		if err := json.NewDecoder(resp.Body).Decode(&fileList); err != nil {
			log.Printf("Failed to parse file list: %v", err)
			return
		}

		totalWords := 0

		for _, file := range fileList {
			fileURL := fmt.Sprintf("http://localhost:8080/files/%s", file)
			resp, err := http.Get(fileURL)
			if err != nil {
				log.Printf("Failed to fetch file %s: %v", file, err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Failed to retrieve file %s: %s", file, resp.Status)
				continue
			}

			content, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read content of file %s: %v", file, err)
				continue
			}

			// Calculate words in the file
			wordCount := len(strings.Fields(string(content)))
			totalWords += wordCount
		}

		fmt.Printf("Total word count in all files: %d\n", totalWords)
	},
}

func init() {
	rootCmd.AddCommand(wcCmd)
}
