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

var addCmd = &cobra.Command{
	Use:   "add [files]",
	Short: "Add files to the store",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, filename := range args {
			filePath := filepath.Clean(filename)
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filename, err)
				continue
			}
			url := fmt.Sprintf("http://localhost:8080/files/%s", filepath.Base(filename))
			req, err := http.NewRequest("POST", url, bytes.NewReader(content))
			if err != nil {
				log.Printf("Failed to create request for file %s: %v", filename, err)
				continue
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Failed to send request for file %s: %v", filename, err)
				continue
			}
			if resp.StatusCode == http.StatusCreated {
				fmt.Printf("Added file %s to the store.\n", filename)
			} else if resp.StatusCode == http.StatusConflict {
				fmt.Printf("File %s already exists in the store.\n", filename)
			} else {
				fmt.Printf("Failed to add file %s: %s\n", filename, resp.Status)
			}
			resp.Body.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
