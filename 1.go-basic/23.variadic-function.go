package main

import "fmt"

//! parameter yg berada diposisi terakhir, memiliki kemampuan dijadikan sebuah varargs =? datanya bisa menerima > 1 input, atau anggap saja semacam array
//? parameter tipe array => wajib membuat array terlebih dahulu sblm mengirimkan ke function
//? parameter varargs => langsung mengirim datanya, jika > 1, cukup menggunakan tanda koma

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)
}
