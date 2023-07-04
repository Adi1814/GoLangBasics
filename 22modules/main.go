package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("hey there mod users")
}
func serveHome(w http.ResponseWriter, r *http.Request) /*r is recieving request like url and stuff and sending response is done through w*/ {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))

}
