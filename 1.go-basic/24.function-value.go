package main

import "fmt"

// function => first class citizen
// function merupakan tipe data dan bisa disimpan dalam variable
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Acacil")

	fmt.Println(result)
	fmt.Println(getGoodBye(("Rama")))
}
