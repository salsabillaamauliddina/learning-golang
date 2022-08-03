package main

import "fmt"

// ! struct => template data / prototype data
// ? struct tidak bisa langsung digunakan
// * bisa membuat data/obj dari struct yg telah kita buat
type Customer struct {
	// Name    string
	// Address string
	Name, Address string
	Age           int
}

func sayHi(customer Customer, name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	sayHi(eko, "Joko")

	// fmt.Println(eko)
	// fmt.Println(eko.Name)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Indonesia",
	// 	Age:     30,
	// }

	// budi := Customer{"Budi", "Jakarta", 40}
	// fmt.Println(joko)
	// fmt.Println(budi)

}

//  struct literals
