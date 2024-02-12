# ptable

Simple helper that print the slice of structs as a table.

## Installation

```shell
go get github.com/kazhuravlev/ptable
```

## Example

```go
package main

import "github.com/kazhuravlev/ptable"

func main() {
	type User struct {
		ID   int
		Name string
		Age  int
	}

	data := []User{
		{ID: 1, Name: "Bob", Age: 42},
		{ID: 2, Name: "Alice", Age: 32},
	}

	ptable.Println(data)
}
```
