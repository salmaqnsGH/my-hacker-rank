package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

// func main() {
// 	fmt.Println(factorialRecursive(5))
// }

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(i), ", ")
	}
	fmt.Println()

	for i := 1; i < 10; i++ {
		fmt.Print(recursive1(4, i), ", ")
	}
	fmt.Println()

	for i := 1; i < 10; i++ {
		fmt.Print(recursive1(5, i), ", ")
	}
	fmt.Println()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 4, 8, 12, 16, 20, 24, 28, …
func recursive1(num int, max int) int {
	if max == 1 {
		return num * max
	}

	return num + recursive1(num, max-1)
}

// 5, 10, 15, 20, 25, …
func recursive2(num int, max int) int {
	if max == 1 {
		return num * max
	}

	return num + recursive1(num, max-1)
}
