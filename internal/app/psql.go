package vision

import (
	"fmt"

	"github.com/jackc/pgx"
)

func NewPostgreSQLDataBase(connectionString string) (*pgx.ConnPool, error) {
	fmt.Println(connectionString)
	pgxConn, err := pgx.ParseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	pgxConn.PreferSimpleProtocol = true
	config := pgx.ConnPoolConfig{
		ConnConfig:     pgxConn,
		MaxConnections: 100,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}

	connPool, err := pgx.NewConnPool(config)
	if err != nil {
		return nil, err
	}
	return connPool, nil

}
