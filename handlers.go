package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/alzaburetz/myrestAPI/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strings"
)

var db *sql.DB
var err error

func restrictionsIndex(w http.ResponseWriter, r *http.Request) {
	var restriction models.Restriction
	var restrictions []models.Restriction
	db, err = sql.Open("sqlite3", "./database.db3")
	if err != nil {
		log.Printf("Error connection to database: %v", err)
	}
	index := strings.LastIndex(r.URL.Path, "/")
	user := r.URL.Path[index+1:]
	log.Printf("Got request from %s", user)
	ugroup, err := db.Query(`SELECT group_id FROM user WHERE username=?`, user)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err = ugroup.Close()
		if err != nil {
			log.Println("Error closing connection")
		}
	}()
	var group int
	for ugroup.Next() {
		err = ugroup.Scan(&group)
		if err != nil {
			log.Println(err)
		}
	}
	res, err := db.Query(`SELECT * FROM restriction WHERE user_gr=?`, group)
	if err != nil {
		log.Printf("Error reading from database: %v", err)
	}
	defer db.Close()
	for res.Next() {
		err = res.Scan(&restriction.ID, &restriction.App, &restriction.Rule, &restriction.Time,
			&restriction.HF, &restriction.HT, &restriction.Exec, &restriction.Group)
		restrictions = append(restrictions, restriction)
		if err != nil {
			log.Printf("Error displayong data: %v", err)
		}

	}
	json.NewEncoder(w).Encode(restrictions)

}

func handleUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var users []models.User
	db, err = sql.Open("sqlite3", "./database.db3")
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
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func handleGroups(w http.ResponseWriter, r *http.Request) {
	db, err = sql.Open("sqlite3", "./database.db3")
	if err != nil {
		fmt.Printf("%v", err)
	}

	var usergroup models.Usergroup
	var usergroups []models.Usergroup
	res, err := db.Query(`SELECT * FROM usergroup`)

	for res.Next() {
		err = res.Scan(&usergroup.ID, &usergroup.Groupname)
		if err != nil {
			log.Print(err)
		}

		usergroups = append(usergroups, usergroup)
	}
	json.NewEncoder(w).Encode(usergroups)
}

func restrictionsAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "addrestriction.html")
	}
	if r.Method == "POST" {
		db, err := sql.Open("sqlite3", "./database.db3")
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		_, err = db.Exec(`INSERT INTO restriction(app,rule,time,hours_from,hours_to,executable,user_gr) VALUES(?,?,?,?,?,?,?)`,
			r.FormValue("app"),
			r.FormValue("rule"),
			r.FormValue("time"),
			r.FormValue("hours_from"),
			r.FormValue("hours_to"),
			r.FormValue("executable"),
			r.FormValue("user"))
		if err != nil {
			log.Println(err)
		}

	}
}
