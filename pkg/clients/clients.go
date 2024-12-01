package clients

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const serverURL = "http://localhost:8080"

func AddFiles(files []string) {
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}
		hash := generateHash(content)

		resp, err := http.Get(serverURL + "/has?hash=" + hash)
		if err != nil {
			fmt.Printf("Error checking hash for file %s: %v\n", file, err)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("File %s already exists on server.", file)
			continue
		}

		resp, err = http.Post(serverURL+"/add?file="+file, "application/octet-stream", bytes.NewReader(content))
		if err != nil {
			fmt.Printf("Error uploading file %s: %v\n", file, err)
			continue
		}
		resp.Body.Close()
		fmt.Printf("File %s uploaded  successfully.\n", file)
	}
}

// Lists all files in the store
func ListFiles() {
	resp, err := http.Get(serverURL + "/ls")
	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Files in the store:")
	fmt.Println(string(body))
}

// Removes a file from the store
func RemoveFile(file string) {
	req, err := http.NewRequest("DELETE", serverURL+"/rm?file="+file, nil)
	if err != nil {
		fmt.Printf("Error creating delete request: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending delete request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", string(body))
}

// Word count of all files in the store
func WordCount() {
	resp, err := http.Get(serverURL + "/wc")
	if err != nil {
		fmt.Printf("Error getting word count: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Word count: %s\n", string(body))
}

// Most or least frequent words in the store
func FreqWords() {
	resp, err := http.Get(serverURL + "/freq-words")
	if err != nil {
		fmt.Printf("Error getting frequent words: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Frequent words:")
	fmt.Println(string(body))
}

// Helper function to generate a file's hash
func generateHash(content []byte) string {
	hash := md5.Sum(content)
	return hex.EncodeToString(hash[:])
}
