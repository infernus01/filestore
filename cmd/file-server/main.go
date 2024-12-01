package main

import (
	"log"
	"net/http"

	filehandler "github.com/infernus01/FileServer/pkg/fileHandler"
)

func main() {
	http.HandleFunc("/list", filehandler.HandleListFile)
	http.HandleFunc("/add", filehandler.HandleAddFile)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
