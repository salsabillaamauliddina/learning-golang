package main

import "fmt"

func main() {
	// ! bisa ubah kumpulan data -> key dan value

	person := map[string]string{
		"name":    "John",
		"address": "Orchard",
	}

	person["job"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])

	// len(map) => jumlah data di map

	fmt.Println(len(person))

	delete(person, "job")

	fmt.Println(person)
	// map[key] => mengambil data di map dg key
	// map[key] = value => mengubag data di map dg key
	// make(map[typeKey]typevalue) => membuat map baru
	// delete(map,key) => menghapus data di map dengan key
}
