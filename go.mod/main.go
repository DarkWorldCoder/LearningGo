package main

import (
	"fmt"
	"net/http"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
}

func greeter() {
	fmt.Println("Hey ther")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series </h1>"))
}
