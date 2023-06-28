package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	str = removeVowels(str)

	fmt.Println("String sem vogais:", str)
}

func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	result := ""
	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			result += string(char)
		}
	}
	return result
}
