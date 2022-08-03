package main

import "fmt"

func main() {
	const age int = 20
	// age = 21
	fmt.Println(age)

	const (
		name  string = "Iteng"
		price        = "Limited edition"
	)

	fmt.Println(name, price)
}
