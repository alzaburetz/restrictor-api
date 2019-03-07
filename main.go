package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	r := mux.NewRouter()
	r.HandleFunc("/restrictions", restrictionsIndex).Methods("GET")
	r.HandleFunc("/restrictions/add", restrictionsAdd).Methods("POST","GET")
	r.HandleFunc("/users", handleUser).Methods("GET")
	r.HandleFunc("/usergroups", handleGroups).Methods("GET")
	log.Fatal(http.ListenAndServe(port, r))
}
