package method

import (
	"fmt"
)

type User struct{
	Name string
}

func (u User) Sapa() string {
	return "Hai, " +u.Name
}

func Meth() {
	u := User{
		Name: "Papoy",
	}

	result := u.Sapa()
	fmt.Println(result)
}