# File Store Application

A simple file store application consisting of a backend API and a CLI client. The backend API uses Go's `net/http` package, and the CLI client is built using `cobra`. The application allows you to add, retrieve, update, and delete text files, as well as perform word count and frequency analysis across all stored files.

---

## Table of Contents
- [File Store Application](#file-store-application)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Starting the Server](#starting-the-server)
    - [Building the CLI](#building-the-cli)
  - [Commands](#commands)
  - [Examples](#examples)
  - [NOTE:](#note)
    - [Step 1:](#step-1)
    - [Step 2:](#step-2)
- [License](#license)
    - [This project is licensed under the MIT License.](#this-project-is-licensed-under-the-mit-license)

---

## Features
- **Add Files**: Upload one or more text files to the store.
- **List Files**: Display all filenames currently stored.
- **Retrieve File**: Get the content of a specific file.
- **Update File**: Update the content of an existing file or add a new file.
- **Delete File**: Remove a file from the store.
- **Word Count**: Get the total word count across all stored files.
- **Frequent Words**: Display the most or least frequent words across all files.

---

## Requirements
- **Go**: Version 1.16 or higher.
- **Cobra**: For the CLI application.

---

## Installation
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/filestore.git
   cd filestore
   ```
2. **Install Dependencies: Install cobra for the CLI application:**
   ```bash
   go get -u github.com/spf13/cobra
   ```

---

## Usage
### Starting the Server
1. **Navigate to the server directory and run the server:**
   ```bash
   cd cmd/file-server
   go run main.go
   ```
2. **The server will start on port 8080.**
### Building the CLI
1. **Navigate to the cli directory and build the CLI application:**
   ```bash
   cd cmd/store-client
   go build -o store main.go
   ```
2. **This will create an executable named store.**

## Commands
1. **Add Files**   
   Upload one or more files to the store:

   ```bash
   ./store add file1.txt file2.txt
   ```

2. **List Files**   
   List all files stored on the server:

   ```bash
   ./store ls
   ```

3. **Get File Content**   
   Retrieve the content of a specific file:

   ```bash
   ./store get file1.txt
   ```

4. **Update File**   
   Update the content of an existing file or add a new file:

   ```bash
   ./store update file1.txt
   ```

5. **Delete File**   
   Remove a file from the store:

   ```bash
   ./store rm file1.txt
   ```

6. **Word Count**   
   Get the total word count across all stored files:

   ```bash
   ./store wc
   ```

7. **Frequent Words**   
   Display the most or least frequent words across all files:

   ```bash
   ./store freq-words [--limit|-n NUMBER] [--order=asc|dsc]
   ```


## Examples
1. **Adding Files**
     ```bash
     echo "Hello World" > hello.txt
     echo "Go is awesome" > go.txt
     ./store add hello.txt go.txt
     ```
2. **Listing Files**
    ```bash
    ./store ls
    ```
3. **Retrieving File Content**
    ```bash
    ./store get hello.txt
    ```
4. **Updating a File**
    ```bash
    echo "Hello Go" > hello.txt
    ./store update hello.txt
    ./store get hello.txt
    ```
5. **Deleting a File**
    ```bash
    ./store rm go.txt
    ./store ls
    ```
6. **Word Count**
    ```bash
    ./store wc
    ```
7. **Frequent Words**
    ```bash
    ./store freq-words -n 5 --order=dsc
    ```

## NOTE:
You can also try using the Docker build
### Step 1:
   ```bash
   docker build . -t fileserver
   ```

### Step 2:
```bash
docker run -p 8081:8080 fileserver:latest
```

One would see the following logs and can access the server at `localhost:8081/files`
`2024/12/03 06:31:10 Server is running on port 8080...`


# License
### This project is licensed under the MIT License.
