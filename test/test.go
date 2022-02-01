package main

import (
	"fmt"
	"net/http"
)

func SimpleIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "index.gohtml")
}

func main() {

	fmt.Println("Server Starting")
	http.HandleFunc("/", SimpleIndexHandler)
	http.HandleFunc("/index", HttpFileHandler)

	http.ListenAndServe(":8080", nil)
}
