package main

import("fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	
)

type User struct {
	FirstName : string,
	LastName : string,
}


func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,r.Header)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([] byte,len)
	r.Body.Read(body)
	r.Body.Close()
	fmt.Fprintln(w,string(body))
 
}


func form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w,r.Form)
}


func multipartform(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileheader := r.MultipartForm.File["uploaded"][0]
	file,err := fileheader.Open()
	if err == nil {
		data,err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w,string(data))
		}
	}

}

func writeExample (w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w,"Not implemented")
}

func redirect (w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Location","www.google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	user := &User{
		FirstName: "Gareth",
		LastName: "Badenhorst",
	}
	json,_ := json.Marshal(user)
	w.Write(json)
		
}


func main () {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body",body)
	http.HandleFunc("/form",form)
	http.HandleFunc("/multipartform",multipartform)
	http.HandleFunc("/writeExample",writeExample)
	http.HandleFunc("/writeHeaderExample",writeHeaderExample)
	http.HandleFunc("/redirect",redirect)
	http.HandleFunc("/json",jsonExample)
	server.ListenAndServe()
}

