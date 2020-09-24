package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	pkgpath = "plugins"
)

var (
	bindserver = "127.0.0.1"
	bindport   = "8090"
)

/*
WebReturn generic struct for retruning JSON output
*/
type WebReturn struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

/*
chkErr Generic error handling
*/
func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// set bindserver and bindport from environment
	if os.Getenv("BIND_SERVER") != "" {
		bindserver = os.Getenv("BIND_SERVER")
	}
	if os.Getenv("BIND_PORT") != "" {
		bindport = os.Getenv("BIND_PORT")
	}

	_ = os.Mkdir(pkgpath, 0644)

	r := mux.NewRouter()

	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/upload/{plugin}", UploadHandler).Methods("POST")
	r.HandleFunc("/download/{plugin}/{package}", DownloadHandler)
	r.HandleFunc("/plugins", PluginListHandler)
	r.HandleFunc("/plugins/{plugin}", PluginPackageListHandler)
	r.HandleFunc("/plugins/{plugin}", DeletePluginHandler).Methods("POST")
	r.HandleFunc("/plugins/{plugin}/{package}", DeletePluginHandler).Methods("POST")
	r.HandleFunc("/", StatusHandler)
	fmt.Printf("Running on %s:%s\n", bindserver, bindport)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", bindserver, bindport), r)
	chkErr(err)
}

/*
StatusHandler Gives OK if the API is running and also lists the number of packages

GET
https://krew-museum/status

Returns JSON:
{
    "Status":"OK",
	"Message":"0 Packages"
}
*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	WebResult(w, http.StatusOK, "API is running")
}

/*
WebResult Generic way to return JSON data
*/
func WebResult(w http.ResponseWriter, status int, msg string) {
	var ret WebReturn
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	ret.Status = status
	ret.Message = msg
	jsonbytes, err := json.Marshal(&ret)
	chkErr(err)
	w.Write(jsonbytes)
}
