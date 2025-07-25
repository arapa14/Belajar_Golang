package training

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

var users []User

func AddUser(name string, age int) {
	newUser := User {
		Name: name,
		Age: age,
	}
	users = append (users, newUser)
}

func ShowUsers() {
	if len(users) == 0 {
		fmt.Println("Belum ada data user.")
		return
	}

	fmt.Println("Daftar user:")
	for i, user:= range users {
		fmt.Printf("%d. %s (Umur: %d)\n", i+1, user.Name, user.Age)
	}
}

func Train() {
	var pilihan string

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah user")
		fmt.Println("2. Lihat semua user")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu (1/2/3): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			var name string
			var age int
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&name)
			fmt.Print("Masukkan umur: ")
			fmt.Scanln(&age)

			AddUser(name, age)
			fmt.Println("User berhasil ditambahkan.")
		case "2":
			ShowUsers()
		case "3":
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}