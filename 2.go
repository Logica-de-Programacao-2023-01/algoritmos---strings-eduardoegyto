package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&x)

	y := strings.ReplaceAll(x, " ", "")
	fmt.Println("A string sem espaços em branco é:", y)
}
