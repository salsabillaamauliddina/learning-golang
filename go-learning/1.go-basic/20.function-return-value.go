package main

import "fmt"

// jika function deklarasi dg tipe data pengembalian, maka wajib fuction hrs mengembalikan data
// menggunakan kata kunci return

func getHello(name string) string {
	return "Hello" + name
}
func main() {
	result := getHello("Acacil")
	fmt.Println(result)
}
