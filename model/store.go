package model

import (
	//"context"
	"context"
	"database/sql"
	"encoding/csv"
)

type Database struct {
	DB *sql.DB
}

type Store struct {
	Traker interface {
		MoviesCSV(ctx context.Context , w *csv.Writer) (error)
	}
}

func NewStore(db *sql.DB) Store {
	if db == nil {
		panic("nil pointer passed to NewStore")
	}
	Database := &Database{DB: db}

	return Store{
		Traker:        Database,

	}
}
