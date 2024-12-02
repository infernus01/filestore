package main

import (
	"log"
	"net/http"

	handler "github.com/infernus01/FileServer/pkg/fileHandler"
)

func main() {
	http.HandleFunc("/files", handler.FilesHandler)
	http.HandleFunc("/files/", handler.FileHandler)
	// http.HandleFunc("/wc", wordCountHandler)
	// http.HandleFunc("/freq-words", frequentWordsHandler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
