package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := readInput(scanner)
	fmt.Println("input", input)

	var arr []int32
	for _, v := range input {
		arr = append(arr, v)
	}

	fmt.Println("arr", arr)

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	mid := n / 2
	arr[mid], arr[n-1] = arr[n-1], arr[mid]

	right := arr[mid+1:]
	r := len(right)
	for i := 0; i < r-1; i++ {
		for j := 0; j < r-i-1; j++ {
			if right[j] < right[j+1] {
				right[j], right[j+1] = right[j+1], right[j]
			}
		}
	}

	copy(arr[mid+1:], right)

	fmt.Println(arr)

}

func readInput(scanner *bufio.Scanner) []int32 {
	var result []int32
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		incomingInt, _ := strconv.ParseInt(input, 10, 64)

		result = append(result, int32(incomingInt))
	}

	return result
}
