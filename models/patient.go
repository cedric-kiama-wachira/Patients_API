package models

type Patient struct {
	ID        string `json:"id"`
	Birthdate string `json:"birthdate"`
	First     string `json:"first"`
	Last      string `json:"last"`
	Gender    string `json:"gender"`
}

