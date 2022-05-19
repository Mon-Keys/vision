package domain

import "time"

type AddPassage struct {
	Status int32 `json:"status:"`
	Exit   bool  `json:"is_exit:"`
	PassID int32 `json:"pass_id"`
}

type Passage struct {
	Exit     bool      `json:"is_exit"`
	Fullname string    `json:"fullname"`
	Time     time.Time `json:"time"`
}
