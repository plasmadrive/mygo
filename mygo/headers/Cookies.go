package main

import ("fmt"
	"net/http"
	"strings"
	"time"
)


func cookieHandler(w http.ResponseWriter, r *http.Request) {
	// check if there is a cookie
	_,err := r.Cookie("G")
	// cookies and other headers need to be written before anything else
	if err != nil {
		cookie := http.Cookie{
			Name: "G",
			Value: "Bad",
			HttpOnly: false,
			MaxAge: -1,
			Expires: time.Unix(1,0),
			Domain: "",

		}
		w.Header().Set("Ny Header",cookie.String())
	
	}
	fmt.Fprintln(w,r.Cookies())
}


func headerHandler(w http.ResponseWriter, r *http.Request) {
	for header, values  := range r.Header {
		headerString := strings.Join(values,",")
		fmt.Fprintf(w,"Name : %s Value: %s\n",header,headerString)
	}
}

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/cookies",cookieHandler)

	http.HandleFunc("/headers",headerHandler)

	server.ListenAndServe()
}
