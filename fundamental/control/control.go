package control 

import (
	"fmt"
	"strconv"
)

func Looping() {
	var age int

	fmt.Println("Masukkan Umur")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("Sudah Dewasa")
	} else {
		fmt.Println("Masih Remaja")
	}

	for i := 1;i <= age;i++ {
		fmt.Println("Angka ke" + strconv.Itoa(i))
	}
}