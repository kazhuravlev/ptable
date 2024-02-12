package ptable

import (
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"os"
	"reflect"
)

// Println print slice of structs to the standard output.
func Println[T any](in []T) {
	_, _ = FPrintln(os.Stdout, in)
}

// FPrintln print the table to the writer.
func FPrintln[T any](outWriter io.Writer, in []T) (int, error) {
	sliceType := reflect.TypeOf(in)

	switch sliceType.Kind() {
	default:
		// NOTE: impossible case because we accept the slice. This branch is not reachable.
		return 0, errors.New("not a slice")
	case reflect.Slice, reflect.Array:
	}

	structType := sliceType.Elem()

	switch structType.Kind() {
	default:
		return 0, errors.New("input type should be slice/array of structs")
	case reflect.Struct:
	}

	fieldsCount := structType.NumField()
	fieldsNames := make([]string, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		fieldsNames[i] = structType.Field(i).Name
	}

	header := make(table.Row, fieldsCount)
	for i := range fieldsNames {
		header[i] = fieldsNames[i]
	}

	rows := make([]table.Row, len(in))
	for i := range in {
		sliceElemType := reflect.ValueOf(in[i])
		rows[i] = make(table.Row, fieldsCount)
		for ii := range fieldsNames {
			rows[i][ii] = sliceElemType.Field(ii).Interface()
		}
	}

	t := table.NewWriter()
	t.AppendHeader(header)
	t.AppendRows(rows)

	res := t.Render()
	n, err := outWriter.Write([]byte(res))
	if err != nil {
		return 0, fmt.Errorf("write to output")
	}

	return n, nil
}
