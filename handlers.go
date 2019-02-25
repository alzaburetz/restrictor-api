package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alzaburetz/myrestAPI/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func restrictionsIndex(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(rest)

}


func handleUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var users []models.User
	db, err := sql.Open("sqlite3", "./database.db3")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer db.Close()

	res, err := db.Query(`SELECT * FROM user`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer res.Close()

	for res.Next() {
		err = res.Scan(&user.ID, &user.Username, &user.Group)
		users = append(users,user)
	}
	json.NewEncoder(w).Encode(users)
}

func handleGroups (w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database.db3")
	if err != nil {
		fmt.Printf("%v", err)
	}

	var ug models.Usergroup
	var ugs []models.Usergroup
	//var usergroups models.Usergroups
	res, err := db.Query(`SELECT * FROM usergroup`)


	for res.Next() {
		err = res.Scan(&ug.ID, &ug.Groupname)
		if err != nil {
			log.Fatal(err)
		}

		ugs = append(ugs, ug)
	}
	json.NewEncoder(w).Encode(ugs)
}
