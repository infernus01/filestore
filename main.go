package main

import (
	"fmt"
	"os"

	"github.com/infernus01/fileService/pkg/clients"
	// "github.com/infernus01/fileService/filehandler"
)

func main() {
	// http.HandleFunc("/list", filehandler.HandleListFile)
	// http.HandleFunc("/add", filehandler.HandleAddFile)

	// http.ListenAndServe(":8080", nil)

	if len(os.Args) < 2 {
		fmt.Println("Usage: store <command> [args...]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: store add <file1> <file2> ...")
			return
		}
		clients.AddFiles(os.Args[2:])
	case "ls":
		clients.ListFiles()
	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Usage: store rm <fileName>")
			return
		}
		clients.RemoveFile(os.Args[2])
	case "wc":
		clients.WordCount()
	case "freq-words":
		clients.FreqWords()
	default:
		fmt.Println("Unknown command. Use 'add', 'ls','rm','wc','freq-words'")

	}
}
