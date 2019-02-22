package main

import (
	_ "encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)



func main() {
	port := ":3000"
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")
	r.HandleFunc("/users", handleUser).Methods("GET")
	log.Fatal(http.ListenAndServe(port, r))
}
