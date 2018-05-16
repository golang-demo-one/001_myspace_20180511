package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	PORT = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "Time is now " + time.Now().String()
	fmt.Fprintf(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

/*
func serveError() {
	fmt.Fprintln("There's no way I'll work!")
}
*/

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	//http.HandleFunc("/error", serveError)
	http.ListenAndServe(PORT, nil)
}
