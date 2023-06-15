package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y string
	fmt.Print("Digite a primeira palavra: ")
	fmt.Scanln(&x)
	fmt.Print("Digite a segunda palavra: ")
	fmt.Scanln(&y)

	if strings.EqualFold(x, y) {
		fmt.Println("As palavras são iguais.")
	} else {
		fmt.Println("As palavras são diferentes.")
	}
}
