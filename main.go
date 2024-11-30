package main

import (
	"net/http"

	"github.com/infernus01/fileService/filehandler"
)

func main() {
	http.HandleFunc("/list", filehandler.HandleListFile)

	// handler:=http.HandlerFunc(PlayerServer)
	// log.Fatal(http.ListenAndServe(":5000",handler))
	http.ListenAndServe(":8080", nil)
}
