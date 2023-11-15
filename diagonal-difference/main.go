package main

import "fmt"

func main() {
	arr := [][]int32{
		{6, 8},
		{-6, 9},
	}

	fmt.Println(arr)

	var leftArr []int32
	for i, a := range arr {
		for j := len(a) - 1; j >= 0; j-- {
			if i == j {
				leftArr = append(leftArr, arr[i][j])

			}
		}
	}

	var rightArr []int32
	for i, a := range arr {
		for j := len(a) - 1 - i; j >= 0; j-- {
			rightArr = append(rightArr, arr[i][j])
			break
		}
	}

	totalLeft := sumArray(leftArr)
	totalRight := sumArray(rightArr)

	var total int32

	if totalLeft < 0 && totalRight < 0 {
		if totalLeft > totalRight {
			total = totalLeft - totalRight
		} else {
			total = totalRight - totalLeft
		}
	} else if totalLeft > 0 && totalRight > 0 {
		if totalLeft > totalRight {
			total = totalLeft - totalRight
		} else {
			total = totalRight - totalLeft
		}
	} else if totalRight < 0 {
		total = totalLeft - totalRight
	} else if totalLeft < 0 {
		total = -(totalLeft - totalRight)
	}

	fmt.Println(total)
}

func sumArray(arr []int32) int32 {
	var sum int32
	for _, val := range arr {
		sum += val
	}
	return sum
}
