package main

import "fmt"

// function bisa mengembalikan multiple values
//! menghiraukan return value
//? multiple return value wajib ditangkap semua values
//? jika ingin menghiraukan return value, bisa menggunakan tanda _

func getFullName() (string, string) {
	return "Salsabilla", "Maulididna"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	// firstName, _ := getFullName()
	// fmt.Println(firstName)
}
