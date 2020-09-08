package main

import (
	"errors"
	"net/http"
	"fmt"
	"io"
	"os"

	"github.com/gorilla/mux"
)

/*
UploadFileHelper Saves an uploaded file to the local filesystem

Returns:
  handler.Filename as string
  err as error
*/
func UploadFileHelper(r *http.Request) (string, error) {
	vars := mux.Vars(r)
    plugin := vars["plugin"]
	r.ParseMultipartForm(32 << 20)
	//ParseMultipartForm parses a request body as multipart/form-data
	file, handler, err := r.FormFile("file") //retrieve the file from form data
	//replace file with the key your sent your image with
	if err != nil {
		return "", err
	}
	defer file.Close() //close the file when we finish
	//this is path which  we want to store the file
	filepath := fmt.Sprintf("%s/%s", pkgpath, plugin)
	// make sure the directory for the plugin exists
	_ = os.Mkdir(filepath, 0644)
	filename := fmt.Sprintf("%s/%s", filepath, handler.Filename)
	//check if the file already exists
	if _, err := os.Stat(filename); err == nil {
		return "", errors.New("Package already exists")
    }

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	io.Copy(f, file)
	//here we save our file to our path
	return handler.Filename, nil
}
