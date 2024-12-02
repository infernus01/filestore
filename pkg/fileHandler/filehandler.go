package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var fileStore = NewFileStore()

func (fs *FileStore) GetFile(filename string) (string, bool) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()
	content, exists := fs.files[filename]
	return content, exists
}

type FileStore struct {
	files map[string]string // filename -> content
	mutex sync.RWMutex
}

func NewFileStore() *FileStore {
	return &FileStore{
		files: make(map[string]string),
	}
}

func (fs *FileStore) UpdateFile(filename, content string) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	fs.files[filename] = content
}

func (fs *FileStore) DeleteFile(filename string) bool {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	if _, exists := fs.files[filename]; !exists {
		return false
	}
	delete(fs.files, filename)
	return true
}

func (fs *FileStore) AddFile(filename, content string) bool {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	if _, exists := fs.files[filename]; exists {
		return false // File already exists
	}
	fs.files[filename] = content
	return true
}

func (fs *FileStore) ListFiles() []string {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()
	filenames := make([]string, 0, len(fs.files))
	for filename := range fs.files {
		filenames = append(filenames, filename)
	}
	return filenames
}

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// List files
		files := fileStore.ListFiles()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}
	filename := pathParts[1]
	switch r.Method {
	case http.MethodGet:
		// Get file content
		content, exists := fileStore.GetFile(filename)
		if !exists {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write([]byte(content))
	case http.MethodPost:
		// Add file
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read file content", http.StatusBadRequest)
			return
		}
		success := fileStore.AddFile(filename, string(content))
		if !success {
			http.Error(w, "File already exists", http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case http.MethodPut:
		// Update file
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read file content", http.StatusBadRequest)
			return
		}
		fileStore.UpdateFile(filename, string(content))
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		// Delete file
		success := fileStore.DeleteFile(filename)
		if !success {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// const storageDir = "./file_store"

// func HandleListFile(w http.ResponseWriter, r *http.Request) {
// 	files, err := os.ReadDir(storageDir)
// 	if err != nil {
// 		http.Error(w, "Error reading directory", http.StatusInternalServerError)
// 		return
// 	}
// 	fileList := []string{}
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			fileList = append(fileList, file.Name())
// 		}
// 	}
// 	fmt.Fprintf(w, "stored files:%v", fileList)
// }

// func HandleAddFile(w http.ResponseWriter, r *http.Request) {
// 	fileName := r.Header.Get("File_Name")
// 	if fileName == "" {
// 		http.Error(w, "File-Name is required", http.StatusBadRequest)
// 		return
// 	}
// 	content := []byte("this is a file content")
// 	err := os.WriteFile(storageDir+"/"+fileName, content, 0644)
// 	if err != nil {
// 		http.Error(w, "failed to store file", http.StatusInternalServerError)
// 		return
// 	}
// 	if _, err := os.Stat(storageDir + "/" + fileName); err != nil {
// 		http.Error(w, "FIle already exits", http.StatusConflict)
// 	}
// 	fmt.Fprintf(w, "File %s added successfully", fileName)
// }
