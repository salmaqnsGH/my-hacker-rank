package main

import "fmt"

func main() {
	s := "Hello_World!"
	var k int32 = 4
	fmt.Println("input:", s)
	// fmt.Println("expected: okffng-Qwvb")

	a := "abcdefghijklmnopqrstuvwxyz"
	b := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	alphabetA := map[int32]rune{}
	for i, v := range a {
		alphabetA[int32(i)] = v
	}

	alphabetB := map[int32]rune{}
	for i, v := range b {
		alphabetB[int32(i)] = v
	}

	var l int32 = int32(len(alphabetA)) // 26
	var result string
	var indexA int32
	var indexB int32

	for _, v := range s {
		indexA = getIndex(v, alphabetA)
		indexB = getIndex(v, alphabetB)
		if contains(v, alphabetA) {
			if indexA+k > l-1 {
				result += string(alphabetA[(indexA+k)%l])
				// fmt.Println(indexA, indexA-l+k-1, string(alphabetA[indexA+k-l-1]))
			} else {
				result += string(alphabetA[indexA+k])
				// fmt.Println(string(alphabetA[indexA+k-1]))
			}
		} else if contains(v, alphabetB) {
			if indexB+k > l-1 {
				result += string(alphabetB[(indexB+k)%l])
			} else {
				result += string(alphabetB[indexB+k])
			}
		} else {
			result += string(v)
		}
	}

	fmt.Println("result:", result)
}

func getIndex(input rune, arr map[int32]rune) int32 {
	var index int32
	for i, v := range arr {
		if v == input {
			index = i
		}
	}
	return index
}

func contains(input rune, arr map[int32]rune) bool {
	for _, v := range arr {
		if v == input {
			return true
		}
	}
	return false
}
