package domain

import "time"

type Position struct {
	Latitude  float64 `json:"latitude:"`
	Longitude float64 `json:"longitude:"`
}

type Terminal struct {
	ExpirationDate time.Time `json:"expiration_date:"`
	Name           string    `json:"name:"`
	Position       Position  `json:"position:"`
}
