package main

import "fmt"

// ? kemampuan sebuah function berinteraksi dengan data-data disekitrnya dalam scop yang sama
// ? digunakan dengan bijak saat membuat aplikasi

func main() {
	name := "acacil"
	counter := 0
	increment := func() {
		//? bisa berganti menjadi cacil
		name = "cacil"
		fmt.Println("Increment")
		//? bisa mengakses counter
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
