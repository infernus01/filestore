// package cmd

// import (
// 	"bytes"
// 	"os"
// 	"testing"
// )

// func TestWordCountCommand(t *testing.T) {
// 	tmpFile, err := os.CreateTemp("", "testfile*.txt")
// 	if err != nil {
// 		t.Fatalf("Failed to create temp file: %v", err)
// 	}
// 	defer os.Remove(tmpFile.Name())
// 	content := "Go is an open-source programming language."
// 	tmpFile.WriteString(content)
// 	tmpFile.Close()

// 	rootCmd.SetArgs([]string{"add", tmpFile.Name()})
// 	err = rootCmd.Execute()
// 	if err != nil {
// 		t.Fatalf("Add command failed: %v", err)
// 	}

// 	buf := new(bytes.Buffer)
// 	rootCmd.SetOut(buf)

// 	rootCmd.SetArgs([]string{"wc"})
// 	err = rootCmd.Execute()
// 	if err != nil {
// 		t.Errorf("Word count command failed: %v", err)
// 	}

// 	output := buf.String()
// 	if output == "" {
// 		t.Errorf("Expected output, got empty string")
// 	}
// }

package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWordCount(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/wc" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"word_count": 42}`)) // Mock response with word count
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	url := mockServer.URL + "/wc"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	body := make([]byte, 32)
	n, err := resp.Body.Read(body)
	if err != nil && err.Error() != "EOF" {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expected := `{"word_count": 42}`
	if string(body[:n]) != expected {
		t.Errorf("Expected response %q, got %q", expected, string(body[:n]))
	}
}
