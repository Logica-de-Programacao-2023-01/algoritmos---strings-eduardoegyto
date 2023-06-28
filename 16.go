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
	str = strings.ReplaceAll(str, "a", "*")
	str = strings.ReplaceAll(str, "e", "*")
	str = strings.ReplaceAll(str, "i", "*")
	str = strings.ReplaceAll(str, "o", "*")
	str = strings.ReplaceAll(str, "u", "*")

	fmt.Println("String com as vogais substituídas por *:", str)
}
