package main

import ("fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello World")
}


type HelloHandler struct{}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello")
}

type WorldHandler struct{} 
func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"World")
}


func main() {

	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}

	http := http.Server{
		Addr: "127.0.0.1:8080",

	}
	
	http.Handle("/hello",&helloHandler)
	http.Handle("world",&worldHandler)

	http.ListenAndServe()
}
