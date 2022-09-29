package main

// !golang memilili interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error

/*
Membuat Error
?Untuk membuat error, tidak perlu manual
?Golang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors
*/
type error interface {
	Error() string
}

func main() {

}
