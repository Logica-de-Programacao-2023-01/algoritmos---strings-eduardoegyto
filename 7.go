package main

import (
	"fmt"
	"unicode"
)

func main() {
	var x string
	fmt.Print("Digite algo: ")
	fmt.Scanln(&x)

	y := false
	for _, char := range x {
		if unicode.IsDigit(char) {
			y = true
			break
		}
	}

	if y {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém números.")
	}
}
