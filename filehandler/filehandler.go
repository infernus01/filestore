package filehandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func HandleAddFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.Header.Get("File_Name")
	if fileName == "" {
		http.Error(w, "File-Name is required", http.StatusBadRequest)
		return
	}
	content := []byte("this is a file content")
	err := ioutil.WriteFile(storageDir+"/"+fileName, content, 0644)
	if err != nil {
		http.Error(w, "failed to store file", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(storageDir + "/" + fileName); err != nil {
		http.Error(w, "FIle already exits", http.StatusConflict)
	}
	fmt.Fprintf(w, "File %s added successfully", fileName)
}
