package main

import "fmt"

/*
!Interface -> tipe data Abstract, tidak memiliki implementasi langsung
!Sebuah interface berisikan definisi-definisi method
!Biasanya interface digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string
}

func SayHi(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

/*
*Setiap tioe data yang ssuai dengan kontrak interfase, secara otomatis dianggap sebagai interface tersebut
*Sehingga tidak perlu mengimplementasikan interface secara manual
*Hal ini berbeda dengan bahsa pemrograman lainyang ketika membuat interface, kita haru menyebutkan secara eksplisit akan menggunakan interface mana
 */

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var Eko Person
	Eko.Name = "Eko"

	SayHi(Eko)
	// ! tidak mengeluarkan error karena sudah ada kontrak dengan HasName
	cat := Animal{
		Name: "Purr",
	}

	SayHi(cat)
}
