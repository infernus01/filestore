package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [filename]",
	Short: "Update a file in the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		filePath := filepath.Clean(filename)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read file %s: %v", filename, err)
			return
		}
		url := fmt.Sprintf("http://localhost:8080/files/%s", filepath.Base(filename))
		req, err := http.NewRequest("PUT", url, bytes.NewReader(content))
		if err != nil {
			log.Printf("Failed to create request for file %s: %v", filename, err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Failed to send request for file %s: %v", filename, err)
			return
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Updated file %s in the store.\n", filename)
		} else {
			fmt.Printf("Failed to update file %s: %s\n", filename, resp.Status)
		}
		resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
