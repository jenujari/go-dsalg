package main

import "fmt"

func main() {
	stringOne := "This is not palindrome string"
	revStrOne := reverse_str(stringOne)

	stringTwo := "madam"
	revStrTwo := reverse_str(stringTwo)

	fmt.Println(stringOne)
	fmt.Println(revStrOne)

	fmt.Println(stringTwo)
	fmt.Println(revStrTwo)
}

func reverse_str(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
