package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const lfn = "/var/log/messages"

func init() {
	fl, err := os.Create(lfn)
	if err != nil {
		fmt.Printf("Error creating %s, %v", lfn, err)
	}
	defer fl.Close()

}

//
type requestDataHeader struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type requestDataHeaders struct {
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fl, err := os.OpenFile(lfn, os.O_WRONLY|os.O_APPEND, 0666)
	defer fl.Close()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}
	path := r.URL.Path
	var rdHdr requestDataHeader
	logEncoder := json.NewEncoder(fl)
	rdHdr = requestDataHeader{Path: path, Name: "UserAgent", Value: "one"}
	err = logEncoder.Encode(rdHdr)
	/*
		for headerName, _ := range r.Header {
			headerValue := r.Header.Get(headerName)

			rdHdr = requestDataHeader{Path: path, HeaderName: headerName, HeaderValue: headerValue}
			err = logEncoder.Encode(rdHdr)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
		}
	*/

	headers := make(map[string]string)
	for header, _ := range r.Header {
		headers[header] = r.Header.Get(header)
	}
	rdHdrs := requestDataHeaders{Path: path, Headers: headers}

	httpEncoder := json.NewEncoder(w)
	err = httpEncoder.Encode(rdHdrs)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)

}

func main() {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", handler)
	server.ListenAndServe()

}
