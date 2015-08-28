package helper

import (
	_ "AddressService/server/mysql"
	"database/sql"
	"strconv"
	// "fmt"
)

func ProcessResult(rows *sql.Rows) (results map[string]map[string]string, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	results = make(map[string]map[string]string)
	n := 0
	for rows.Next() {
		row := make(map[string]string)
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			row[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}
		results[strconv.Itoa(n)] = row
		// fmt.Println("-----------------------------------")
		n++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
