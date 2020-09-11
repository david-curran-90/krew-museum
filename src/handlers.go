package main

import (
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"strconv"
	"os"

	"github.com/gorilla/mux"
)

type pluginList struct {
	Plugin string `json:"plugin"`
}

/*
UploadHandler Accepts a file and saves it to the filesystem

POST
https://krew-museum/upload/{plugin}
*/
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, err := UploadFileHelper(r)
	var msg string
	var status int
	if err != nil {
		msg = fmt.Sprintf("Failed to create file")
		status = http.StatusBadRequest
		fmt.Println(err)
	} else {
		msg = fmt.Sprintf("File created %s", file)
		status = http.StatusOK
		fmt.Println(msg)
	}
	WebResult(w, status, msg)
}

/*
DownloadHandler Takes a package name and offers it for download

GET
https://krew-museum/download/{plugin}/{package}

Downloads package from $pkgpath/$plugin/$package
e.g.
/plugins/kubectl-plugin.1.0.0.tar.gz
*/
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plugin := vars["plugin"]
	pkg := vars["package"]

	filename := fmt.Sprintf("%s/%s/%s", pkgpath, plugin, pkg)
	openfile, err := os.Open(filename)
	defer openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		WebResult(w, http.StatusNotFound, "File not found")
	}


	
	//Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openfile.Read(fileHeader)
	//Get content type of file
	fileContentType := http.DetectContentType(fileHeader)

	//Get the file size
	fileStat, _ := openfile.Stat()                     //Get info from file
	fileSize := strconv.FormatInt(fileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Length", fileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	openfile.Seek(0, 0)
	io.Copy(w, openfile) //'Copy' the file to the client
}

/*
PluginListHandler Lists available pacakges

GET
https://krew-museum/plugins
*/
func PluginListHandler(w http.ResponseWriter, r *http.Request) {
	var pl []pluginList
	plugins, err := ioutil.ReadDir(pkgpath)
	if err != nil {
		panic(err)
	}

	for _, p := range plugins {
		pl = append(pl, pluginList{p.Name()})
	}
	jsonbytes, err :=json.Marshal(&pl)
	chkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonbytes)
}
