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

	ptable.Println(data)
	// Output:
	// +----+-------+-----+
	// | ID | NAME  | AGE |
	// +----+-------+-----+
	// |  1 | Bob   |  42 |
	// |  2 | Alice |  32 |
	// +----+-------+-----+
}
