package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	isCamelCase := false
	words := 1

	if len(str) > 0 && unicode.IsLower(rune(str[0])) {
		isCamelCase = true
		for _, char := range str {
			if unicode.IsUpper(char) {
				words++
			}
		}
	}

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavra(s).\n", words)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
