package postgres

import (
	"bytes"
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
