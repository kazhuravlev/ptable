package ptable_test

import (
	"bytes"
	"github.com/kazhuravlev/ptable"
	"strings"
	"testing"
)

func ExamplePrintln() {
	type User struct {
		ID   int
		Name string
		Age  int
	}

	data := []User{
		{
			ID:   1,
			Name: "Bob",
			Age:  42,
		},
		{
			ID:   2,
			Name: "Alice",
			Age:  32,
		},
	}

	ptable.Println(data, ptable.WithIncludeFields("ID", "Name"))
	// Output:
	// +----+-------+
	// | ID | NAME  |
	// +----+-------+
	// |  1 | Bob   |
	// |  2 | Alice |
	// +----+-------+
}

func TestFPrintln_UnknownIncludedFieldReturnsError(t *testing.T) {
	type User struct {
		ID int
	}

	var out bytes.Buffer

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("FPrintln panicked for an unknown include field: %v", r)
		}
	}()

	_, err := ptable.FPrintln(&out, []User{{ID: 1}}, ptable.WithIncludeFields("Missing"))
	if err == nil {
		t.Fatal("expected an error for an unknown include field")
	}
}

func TestFPrintln_UnexportedFieldReturnsError(t *testing.T) {
	type User struct {
		ID     int
		secret string
	}

	var out bytes.Buffer

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("FPrintln panicked for a struct with an unexported field: %v", r)
		}
	}()

	_, err := ptable.FPrintln(&out, []User{{ID: 1, secret: "hidden"}})
	if err == nil {
		t.Fatal("expected an error for a struct with an unexported field")
	}

	if strings.Contains(out.String(), "hidden") {
		t.Fatal("unexported field value leaked into output")
	}
}
