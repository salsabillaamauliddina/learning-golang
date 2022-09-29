package main

import "fmt"

/*
!Interface Kosong
?Golang bukan bahasa pemrograman yang berorientasi objek
?Biasanya dlm bahsa pemrograman berorientasi objek, ada 1 data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada dibahasa pemrograman tersebut
?Contoh di java -> java.lang.Object
?Untuk mengangani kasus seperti ini, Golang bisa menggunakan interface kosong
?Interface Kosong -> interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
*/

/*
contoh Inteface Kosong
*fmt.Println(a ... interface{})
*panic(v interface{})
*recover()interface{}
dan lain lain
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	// bukan string,boolean tapi berupa interface kosong -> maka akan mengikuti tipe data nya
	// var data int = Ups(1)
	// ! tidak bisa menggunakan spesifik int hanya bisa memakai interface kosong
	var data interface{} = Ups(3)
	fmt.Println(data)
}
