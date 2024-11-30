package filehandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const storageDir = "./file_store"

func HandleListFile(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(storageDir)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	fileList := []string{}
	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}
	fmt.Fprintf(w, "stored files:%v", fileList)
}
