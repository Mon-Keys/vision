package model

import "time"

type Pass struct {
	IssueDate	   time.Time 		`json:"issue_date"`
	ExpirationDate time.Time 		`json:"expiration_date"`
	SecureData 	   string 	 		`json:"secure_data"`
	Name		   string	 		`json:"pass_name"`
	DynamicQR	   bool		 		`json:"is_dynamic"`
	Passages	   []Passage		`json:"passages"`
}