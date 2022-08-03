package main

import "fmt"

// !break => menghentikan seluruh perulangan
// !continue => skip perulangan

// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i == 5 {
// 			break
// 		}
// 		fmt.Println("perulangan ke ", i)
// 	}
// }

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan continue ke ", i)
	}
}
