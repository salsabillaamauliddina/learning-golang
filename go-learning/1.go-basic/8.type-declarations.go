package main

import "fmt"

func main() {
	// ! membuat ulang tipe data baru dari tipe data yg sudah ada
	type NoKTP string
	// ? sebagai aliases
	var ktpAca NoKTP = "12345"

	fmt.Println(ktpAca)
}
