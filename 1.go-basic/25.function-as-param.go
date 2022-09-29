package main

import "fmt"

// ! function tidak hanya bisa disimpan dalam variable sebagai value
// ? namun juga bisa sebagai parameter untuk function lain

func sayHelloWithFilter(name string, filter func(string) string) {
	// parameter function berupa string
	// memanggil filter sebagai function
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Rama", spamFilter)

	// filter := spamFilter
	sayHelloWithFilter("Anjing", spamFilter)
}
