package main

import "fmt"

/*
?Biasanya dalam bahasa pemrograman lain,object yang belum diinisiasi maka secara otomatis nilainya adalah null / nil
*Berbeda dgn golang, di golang saatkita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan defaul value nya
!Namun di golang ada data nil , yaitu data kosong
?Nil seniri hanya bisa digunakan dibeberapa tioe data, seperti interface, function,map,slice,pointer dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// person := NewMap("Acacil")
	// fmt.Println(person)

	// ? tanpa nil
	// var person map[string]string

	// if person["name"] == "" {
	// 	fmt.Println("Data Kosong")
	// } else {
	// 	fmt.Println(person)
	// }

	var person map[string]string = NewMap("Acacil")

	if person == nil {
		fmt.Println("data Kosong")
	} else {
		fmt.Println(person)
	}
}
