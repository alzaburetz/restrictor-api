package models

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Group int `json:"group_id"`
}

type Users struct {
	Users []User `json:"users"`
}