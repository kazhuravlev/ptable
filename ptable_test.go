package ptable_test

import (
	"github.com/kazhuravlev/ptable"
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
