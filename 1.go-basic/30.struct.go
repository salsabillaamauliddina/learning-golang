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
	// Married       bool
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

	//!  Struct Literals

	Sule := Customer{
		Name:    "Sule",
		Address: "Jakarta",
		Age:     30,
	}

	fmt.Println(Sule)

	// * posisi harus sama / akan error
	Budi := Customer{"Budi", "Jakarta", 40}
	fmt.Println(Budi)
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
