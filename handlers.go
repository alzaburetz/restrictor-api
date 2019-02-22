package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alzaburetz/myrestAPI/models"
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

	var userss models.Users
	userss.Users = users

	byteval, _  := json.Marshal(userss)
	fmt.Fprintf(w, "%v", string(byteval))
}
