package model

import (
	"context"
	"database/sql"
	"encoding/csv"
)

type Database struct {
	DB *sql.DB
}

type Store struct {
	Traker interface {
		CSVTabls(ctx context.Context, w *csv.Writer, table string) error
	}
}

func NewStore(db *sql.DB) Store {
	if db == nil {
		panic("nil pointer passed to NewStore")
	}
	Database := &Database{DB: db}

	return Store{
		Traker: Database,
	}
}
