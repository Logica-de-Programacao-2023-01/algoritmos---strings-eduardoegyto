package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	str = strings.ToLower(str)

	uniqueLetters := ""
	for _, char := range str {
		count := strings.Count(str, string(char))
		if count == 1 {
			uniqueLetters += string(char)
		}
	}

	fmt.Println("Letras únicas da string:", uniqueLetters)
}
