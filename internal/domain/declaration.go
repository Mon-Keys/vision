package domain

import "time"

type Declaration struct {
	Creator      Account            `json:"creator"`
	Comment      string             `json:"comment"`
	Name         string             `json:"name"`
	CreationDate time.Time          `json:"creation_date"`
	Events       []DeclarationEvent `json:"events"`
	Requests     []PassRequest      `json:"pass_requests"`
}
