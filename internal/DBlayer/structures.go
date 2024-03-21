package dblayer

import "database/sql"

type Query struct {
	Authorization
}

func NewDBH(db *sql.DB) *Query {
	return &Query{
		Authorization: authNew(db),
	}
}
