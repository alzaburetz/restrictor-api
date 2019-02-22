package models

type Restriction struct {
	ID int `json:"id"`
	App string `json:"app"`
	Rule string `json:"rule"`
	Time string `json:"time"`
	HF int `json:"hours_from"`
	HT int `json:"hours_to"`
	Exec string `json:"executable"`
	Group int `json:"user_gr"`
}

type Restrictions struct {
	Restrict []Restriction `json:"restrictions"`
}
