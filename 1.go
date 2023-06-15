package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&x)

	m := strings.ToUpper(x)
	fmt.Println("A palavra em letras maiúsculas é:", m)
}
