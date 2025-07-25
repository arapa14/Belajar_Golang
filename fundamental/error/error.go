package error

import (
	"fmt"
	"strconv"
)

func ErrHand() {
	number, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Terjadi Error", err)
	} else {
		fmt.Println("Angka:", number)
	}
}