package domain

import "errors"

var (
	ErrorJSONSerialization     = errors.New("Json serialization error")
	ErrorPSQLEmpty             = errors.New("Empty PSQL response")
	ErrorUserConflict          = errors.New("PSQL user conflict")
	ErrorCantFindUserWithEmail = errors.New("Can`t find user with given email")
	ErrorWrongPassword         = errors.New("Wrong password for given email")
)
