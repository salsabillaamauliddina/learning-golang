package main

import "fmt"

// ! if short statement

func main() {
	name := "Dedek"

	if length := len(name); length > 5 {
		fmt.Println("Too many")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
