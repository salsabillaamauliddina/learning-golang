package main

import "fmt"

// ? function tanpa nama

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blockeed", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// ! tidak wajib
// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	// * anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("aca", blacklist)

	// ! anonymous function inside params
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
