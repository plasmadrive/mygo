package main


import (
	"io"
	"net/http"
)

func echo (w http.ResponseWriter, r *http.Request) {

	_,err := io.Copy(w,r.Body)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
}


func main() {
	server := http.Server{Addr: ":8080",}
	http.HandleFunc("/",echo)
	server.ListenAndServe()
}
