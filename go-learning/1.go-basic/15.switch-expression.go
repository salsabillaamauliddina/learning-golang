package main

import "fmt"

func main() {

	name := "Acil"

	switch name {
	case "Acil":
		fmt.Println("Hai acill")
	case "Rama":
		fmt.Println("Hai rama")
	default:
		fmt.Println("Halo kamu siapa?")
	}

	// ! short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Too many")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
