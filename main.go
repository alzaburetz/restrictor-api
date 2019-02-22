package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/alzaburetz/myrestAPI/models"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)


func handleHome(w http.ResponseWriter, r *http.Request) {
	var restriction models.Restriction
	var rest []models.Restriction
	db, err := sql.Open("sqlite3", "./database.db3")
	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
	}
	res, err := db.Query(`SELECT * FROM restriction`)
	if err != nil {
		log.Fatalf("Error reading from database: %v", err)
	}
	defer db.Close()
	for res.Next() {
		err = res.Scan(&restriction.ID,&restriction.App,&restriction.Rule,&restriction.Time,
			&restriction.HF, &restriction.HT, &restriction.Exec, &restriction.Group)
		rest = append(rest,restriction)
		if err != nil {
			log.Fatalf("Error displayong data: %v", err)
		}

	}
	var restrictions models.Restrictions
	restrictions.Restrict = rest
	byteval, err := json.Marshal(restrictions)
	if err != nil {
		log.Printf("%v",err)
	}
	fmt.Fprintf(w, "%v\n", string(byteval))

}


func main() {
	port := ":3000"
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")
	log.Fatal(http.ListenAndServe(port, r))
}
