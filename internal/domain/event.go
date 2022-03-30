package domain

import "time"

type EventType struct {
	Name string `json:"name"`
}

type DeclarationEvent struct {
	Creator Account   `json:"account"`
	Time    time.Time `json:"time"`
	Comment string    `json:"comment"`
}
