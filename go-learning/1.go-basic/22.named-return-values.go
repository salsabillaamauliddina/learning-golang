package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Ramadhian"
	middleName = "Putra"
	lastName = "nanda"

	return
	// krn sudah di jabarkan di line 5
	// return firstName, middleName, lastName
}

func main() {
	// firstName, middleName, lastName := getCompleteName()
	// fmt.Print(firstName, middleName, lastName)
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

}
