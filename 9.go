package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y, z string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&x)
	fmt.Print("Digite as letras a serem substituído: ")
	fmt.Scanln(&y)
	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&z)

	nova := strings.ReplaceAll(x, y, z)
	fmt.Println("A nova palavra é:", nova)
}
