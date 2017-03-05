package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"
)

var logFile *os.File

const (
	logFileName = "/var/log/app"
)

func init() {
	logFile, err := os.Create(logFileName)
	if err != nil {
		os.Exit(1)
	}
	logFile.Close()
}

type requestInfo struct {
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	headers := make(map[string]string)
	for header, _ := range r.Header {
		headers[header] = r.Header.Get(header)
	}

	//	fmt.Printf("Path : %s, Headers : %v", path, headers)
	reqInfo := requestInfo{Path: path, Headers: headers}
	log(&reqInfo)
	w.WriteHeader(200)
	w.Write([]byte(path))

}

func log(rInf *requestInfo) {

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	bs, err := json.Marshal(rInf)
	if err != nil {
		fmt.Fprintf(logFile, "Error marshalling %v\n", err)
	} else {
		fmt.Fprintln(logFile, string(bs))

	}
	err = logFile.Sync()
	err = logFile.Close()
	fmt.Println(string(bs))

}

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", handler)
	server.ListenAndServe()
}
