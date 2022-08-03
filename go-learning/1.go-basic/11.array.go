package main

import "fmt"

func main() {
	// ? array -> tipe data yang berisi kumpulan data dg tipe yg sama
	// ? harus menenetukan jumlah data yg bisa ditampung oleh array
	// ? daya tampung array tidak bisabertambah setelah array dibuat

	var names [5]string
	names[0] = "ucil"
	names[1] = "icil"
	names[2] = "acil"
	names[3] = "ucul"
	names[4] = "ucil"

	fmt.Println(names[0])
	fmt.Println(names[4])

	var ages = [3]int{
		19,
		20,
		21,
	}

	fmt.Println(ages)

	//! len(array) get the length of the array
	// ? array[index]get the postion of index
	// * array[index] = value change the data of the position of index

	fmt.Println(len(ages))
	fmt.Println(ages[0])

	var example [10]string

	fmt.Println(len(example))

}
