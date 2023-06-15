package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Digite algo: ")
	fmt.Scanln(&x)

	y := strings.Fields(x)
	qnt := len(y)
	fmt.Println("A string contém", qnt, "palavras.")
}
