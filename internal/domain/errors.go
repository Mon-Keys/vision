package domain

import "errors"

var (
	ErrorJSONSerialization = errors.New("Json serialization error")
	ErrorPSQLEmpty         = errors.New("Empty PSQL response")
)
