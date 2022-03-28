package model

import "time"

type User struct {
	ID      	int32     `json:"id,omitempty"`
	Email 		string 	  `json:"email"`
	Created 	time.Time `json:"created"`
	Login 		string 	  `json:"login"`
	Password 	string 	  `json:"password"`
}

type NewUser struct {
	Email 		string 	  `json:"email"`
	Created 	time.Time `json:"created"`
	Login 		string 	  `json:"login"`
	Password 	string 	  `json:"password"`
}