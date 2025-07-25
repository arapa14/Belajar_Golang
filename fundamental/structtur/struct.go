package structtur

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func Structur() {
	u := User{
		Name: "Arapa", Age: 20,
	}

	fmt.Println(u.Name)
}