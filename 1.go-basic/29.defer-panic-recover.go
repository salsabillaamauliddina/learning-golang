package main

import "fmt"

//* defer function -> func yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
// * akan selalu dieksekusi walaupun terjadi error di funtion yang dieksekusi

// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication() {
// 	// ! mencetak run app lalu memanggil func logging
// 	defer logging()
// 	fmt.Println(("Run application"))
// }

// func main() {
// 	runApplication()
// }

// ? panic -> func yang digunakan untuk menghentikan program
// ? panic dipanggil ketika error pada saat program berjalan
// ? saat panic dipanggil, program akan terhenti, namun defer func tetap akan dieksekusi

// func endApp() {
// 	fmt.Println("Aplikasi selesai ")
// }

// func runApp(error bool) {
// 	// ! mencetak run app lalu memanggil func logging
// 	defer endApp()
// 	if error {
// 		panic("APLIKASI ERROR")
// 	}
// 	fmt.Println("Aplikasi berjalan")
// }

// func main() {
// 	runApp(false)
// }

// ! recover -> func yang bisa digunakan untuk menagkap data panic
// ! proses panic akan terhenti sehingga program akan tetap berjalan
func endApp() {
	fmt.Println("Aplikasi selesai ")
	message := recover()
	fmt.Println("error dengan jalan", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("EROR")
	}
}

func main() {
	runApp(true)
	fmt.Println("Aca")
}
