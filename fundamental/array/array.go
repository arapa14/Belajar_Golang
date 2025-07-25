package array

import (
	"fmt"
)

func Arr() {
	number := [3]int{1,2,3} //array (sudah ditentukan)
	name := []string{"Papoy", "Alg"} //slice (tidak ditentukan)

	fmt.Println(number[0])
	fmt.Println(name)
}