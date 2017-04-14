package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	s   rand.Source
	rnd rand.Rand
	m   *sync.Mutex
)

func handle(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	delay := rnd.Int() % 5000
	errProb := rand.Int() % 10
	m.Unlock()
	time.Sleep(delay)
	if errProb == 0 {
		w.WriteHeader(500)
	}

}

func health(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	errProb := rand.Int() % 100
	m.Unlock()

	if errProb == 0 {
		w.WriteHeader(500)
	}
}

func main() {

	s = rand.NewSource(time.Now().UnixNano())
	r = s.New()
	m = &sync.Mutex{}

	server := &http.Server{Addr: "0:8080"}

	http.HandleFunc("/", handle)
	http.HandleFunc("/health", health)

	server.ListenAndServer()

}
