package main

import "fmt"

// ? Struct Method
/*
!Struct tipe data seperti tipe data lainm bisa digunakan sbg parameter u/ function
!Namun jika ingin menambahkan method dlm structs, sehingga seakan-akan sebuah struct memiliki function
!Method adalah function
*/
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My name is ," + customer.Name)
}

func main() {
	Acacil := Customer{Name: "Acacil"}
	Acacil.sayHello()
}
