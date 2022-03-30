package domain

type Role struct {
	Name        string `json:"name"`
	AccessLevel int32  `json:"accesslevel"`
}
