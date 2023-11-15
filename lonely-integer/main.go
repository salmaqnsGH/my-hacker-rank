package main

import "fmt"

func main() {
	a := []int32{1, 2, 3, 4, 4, 2, 1}
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	var twin []int32
	for i := 0; i < n-1; i++ {
		if a[i] == a[i+1] {
			twin = append(twin, a[i])
		}
	}

	var uniq int32

	for _, v := range a {
		if contains(v, twin) {
			continue
		} else {
			uniq = v
		}
	}

	fmt.Println("twin", twin)
	fmt.Println("uniq", uniq)
	fmt.Println(a)
}

func contains(num int32, arr []int32) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
