package model

import "time"

type Passage struct { 
	Status   int32  		`json:"status:"`
	Exit	 bool   		`json:"is_exit:"`
	Time 	 time.Time 		`json:"time"`
	Terminal Terminal 		`json:"terminal"`
}