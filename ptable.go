package ptable

import (
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"os"
	"reflect"
	"strings"
)

// Println print slice of structs to the standard output.
func Println[T any](in []T, opts ...optionFn) {
	_, _ = FPrintln(os.Stdout, in, opts...)
}

// FPrintln print the table to the writer.
func FPrintln[T any](outWriter io.Writer, in []T, opts ...optionFn) (int, error) {
	var opt options
	for i := range opts {
		opts[i](&opt)
	}

	fieldFilter := func(field string) bool { return true }
	if len(opt.IncludeFields) != 0 {
		fieldMap := make(map[string]struct{}, len(opt.IncludeFields))
		for i := range opt.IncludeFields {
			fieldMap[opt.IncludeFields[i]] = struct{}{}
		}

		fieldFilter = func(field string) bool {
			_, ok := fieldMap[field]
			return ok
		}
	}

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
	fieldsNames := make([]string, 0, fieldsCount)
	foundFields := make(map[string]struct{}, fieldsCount)

	for i := 0; i < fieldsCount; i++ {
		structField := structType.Field(i)
		fieldName := structField.Name
		if !fieldFilter(fieldName) {
			continue
		}

		if !structField.IsExported() {
			return 0, fmt.Errorf("field %q is unexported", fieldName)
		}

		fieldsNames = append(fieldsNames, fieldName)
		foundFields[fieldName] = struct{}{}
	}

	if len(opt.IncludeFields) != 0 {
		missingFields := make([]string, 0)
		for i := range opt.IncludeFields {
			if _, ok := foundFields[opt.IncludeFields[i]]; ok {
				continue
			}

			missingFields = append(missingFields, opt.IncludeFields[i])
		}

		if len(missingFields) != 0 {
			return 0, fmt.Errorf("unknown fields: %s", strings.Join(missingFields, ", "))
		}
	}

	header := make(table.Row, len(fieldsNames))
	for i := range fieldsNames {
		header[i] = fieldsNames[i]
	}

	rows := make([]table.Row, len(in))
	for i := range in {
		sliceElemType := reflect.ValueOf(in[i])
		rows[i] = make(table.Row, len(fieldsNames))
		for ii := range fieldsNames {
			rows[i][ii] = sliceElemType.FieldByName(fieldsNames[ii]).Interface()
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
