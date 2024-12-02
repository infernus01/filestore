package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	limit int
	order string
)

var freqWordsCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "Get most or least frequent words",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("http://localhost:8080/freq-words?limit=%d&order=%s", limit, order)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to communicate with server: %v\n", err)
			return
		}
		defer resp.Body.Close()
		var words []struct {
			Word  string `json:"word"`
			Count int    `json:"count"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&words); err != nil {
			fmt.Printf("Failed to parse response: %v\n", err)
			return
		}
		for _, wc := range words {
			fmt.Printf("%s: %d\n", wc.Word, wc.Count)
		}
	},
}

func init() {
	freqWordsCmd.Flags().IntVarP(&limit, "limit", "n", 10, "Number of words to display")
	freqWordsCmd.Flags().StringVar(&order, "order", "dsc", "Order of frequency (asc|dsc)")
	rootCmd.AddCommand(freqWordsCmd)
}
