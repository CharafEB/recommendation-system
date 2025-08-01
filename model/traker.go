package model

import (
	"context"
	"encoding/csv"
	"fmt"
)

func (db *Database) MoviesCSV(ctx context.Context , w *csv.Writer) (error){
	rows, err := db.DB.Query("SELECT * FROM your_table")
	if err != nil {
		return  err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return  err
	}

	w.Write(columns)

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return  err
		}

		record := make([]string, len(columns))
		for i, val := range values {
			if val != nil {
				record[i] = fmt.Sprintf("%v", val)
			} else {
				record[i] = ""
			}
		}

		w.Write(record)
	}

	return  nil
}
