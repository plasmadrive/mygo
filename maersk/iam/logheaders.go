package main

import (
	"fmt"
	"net/http"

)


func handle (w http.ResponseWriter, r * http.Request) {
	headers := r.Header
	path := r.URL.Path
	if len(path) != 0  {
		fmt.Printf("Path %s\n",path)
	}
	if headers != nil {
		if len(headers) == 0 {
			fmt.Println("No request headers sent")
		} else {
			for key,_ := range headers {
				fmt.Printf("%s  : %v\n",key,headers.Get(key))
			}
		}
	}
}





func main () {

	server := &http.Server {Addr: ":8080",}
	http.HandleFunc("/",handle)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Oops something went wrong %v",err)
	}
	

}
