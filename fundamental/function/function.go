package function

import (
	"fmt"
)

func Sapa(nama string) string {
	return "Hallo " + nama
}

func Inp() {
	var name string
	var result string

	fmt.Println("Halo, siapa nama kamu?")
	fmt.Scanln(&name)

	result = Sapa(name)
	fmt.Println(result + " Semoga betah")
}