package database

import (
	"go.fiber.sqlx/app/queries"
)

type Queries struct {
	*queries.HwQuery
	*queries.FileQuery
}

func OpenDBConn() (*Queries, error) {
	db, err := PostgreSQLConn()
	if err != nil {
		return nil, err
	}
	return &Queries{
		HwQuery:   &queries.HwQuery{DB: db},
		FileQuery: &queries.FileQuery{DB: db},
	}, nil
}
