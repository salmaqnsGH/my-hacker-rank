package main

import "fmt"

func main() {
	numbers := []int{6, 1, 4, 2, 9}

	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}

	fmt.Println(numbers)
}
