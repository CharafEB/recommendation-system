package model

import (
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

func (db *Database) CSVTabls(ctx context.Context, w *csv.Writer, table string) error {
	// list of the tabls that allowed 
	allowedTables := map[string]bool{
		"ratings": true,
		"users":  true,
	}

	if !allowedTables[table] {
		return fmt.Errorf("table not allowed: %s", table)
	}

	query := fmt.Sprintf(`SELECT * FROM %q;`, table)

	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	if err := w.Write(columns); err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return err
		}

		record := make([]string, len(columns))
		for i, val := range values {
			if val != nil {
				switch v := val.(type) {
				case []uint8:
					str := string(v)

					if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {

						clean := strings.Trim(str, "{}")
						parts := strings.Split(clean, ",")
						var floats []string
						for _, p := range parts {
							p = strings.TrimSpace(p)
							if _, err := strconv.ParseFloat(p, 32); err == nil {
								floats = append(floats, p)
							}
						}

						record[i] = "[" + strings.Join(floats, ",") + "]"
					} else {
						record[i] = str
					}
				default:
					record[i] = fmt.Sprintf("%v", val)
				}
			} else {
				record[i] = ""
			}
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}
