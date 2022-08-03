package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke - ", counter)
	// 	counter++
	// }

	// for statement
	// ?1. init statement => sblm for di eksekusi
	// ?2. post statement => statement yg akan diexecute di akhir tiap perulangan

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke - ", counter)
	}

	slice := []string{"Eko", "Kurniawan", "Khannedy"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// ! for range
	// for => iterasi terhadap semua data collection
	//  data collection yaitu array,slice, map

	names := []string{"Eko", "Aca", "Rama"}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
