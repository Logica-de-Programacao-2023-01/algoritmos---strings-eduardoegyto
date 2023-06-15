package main

import (
	"fmt"

)

func main() {
	var x string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&x)

	y := ""
	for _, char := range x {
		y = string(char) + y
	}

	fmt.Println("A palavra invertida é:", y)
}
