package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "    another  "
	str = strings.Trim(str, " \t\n\r")

	var reversedString string
	for i := len(str) - 1; i >= 0; i-- {
		reversedString += string(str[i])
	}

	num := "0123456789"

	number := map[int]rune{}
	for i, v := range num {
		number[i] = v
	}

	var result string
	for _, v := range reversedString {
		if !contains(v, number) {
			result += string(v)
		}
	}

	fmt.Println(str)
	fmt.Println(reversedString)
	fmt.Println(result)

}

func contains(input rune, arr map[int]rune) bool {
	for _, v := range arr {
		if v == input {
			return true
		}
	}
	return false
}
