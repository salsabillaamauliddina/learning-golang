package main

import "fmt"

// ! function yang memanggil function sendiri

func factorialLoop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	factorial := factorialLoop(5)
	fmt.Println(factorial)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(factorialRecursive(6))
}
