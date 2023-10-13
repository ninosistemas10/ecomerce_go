package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
)

var ErrFieldAreEmpty = "the fields are empty"

func BuildSQLInsert(table string, fields []string) string {
	if len(fields) == 0 {
		return ErrFieldAreEmpty
	}

	args := bytes.Buffer{}
	values := bytes.Buffer{}
	k := 0

	for _, valor := range fields {
		k++
		args.WriteString(valor)
		args.WriteString(", ")
		values.WriteString(fmt.Sprintf("$%d, ", k))
	}

	args.Truncate(args.Len() - 2)
	values.Truncate(values.Len() - 2)
	return fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", table, args.String(), values.String())
}

func BuildSQLUpdateByID(table string, fields []string) string {
	if len(fields) == 0 {
		return ErrFieldAreEmpty
	}

	// Move ID field to latest.
	fields = append(fields[1:], fields[0])
	args := bytes.Buffer{}
	k := 0
	for _, v := range fields {
		if v == "created_at" {
			continue
		}
		args.WriteString(fmt.Sprintf("%s = $%d, ", v, k+1))
		k++
	}
	args.Truncate(args.Len() - 2)

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", table, args.String(), k)
}


func BuildSQLDelete(table string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)
}


func BuildSQLSelect(table string, fields []string) string {
	if len(fields) == 0 {
		return ErrFieldAreEmpty
	}

	args := bytes.Buffer{}

	for _, valor := range fields {
		args.WriteString(fmt.Sprintf("%s, ", valor))
	}

	args.Truncate(args.Len() - 2)
	return fmt.Sprintf("SELECT %s FROM %s", args.String(), table)
}

func Int64ToNull(d int64) sql.NullInt64 {
	r := sql.NullInt64{Int64: d}
	if d > 0 {
		r.Valid = true
	}

	return r
}

