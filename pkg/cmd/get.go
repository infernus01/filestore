package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [file]",
	Short: "Retrieve the content of a file from the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := filepath.Clean(args[0])
		url := fmt.Sprintf("http://localhost:8080/files/%s", filepath.Base(filename))

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Failed to send GET request for file %s: %v", filename, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			content, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read response body for file %s: %v", filename, err)
				return
			}
			fmt.Printf("Content of file \"%s\":\n%s\n", filename, string(content))
		} else if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("File \"%s\" not found in the store.\n", filename)
		} else {
			fmt.Printf("Failed to retrieve file %s: %s\n", filename, resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
