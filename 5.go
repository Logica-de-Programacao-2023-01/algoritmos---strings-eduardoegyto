package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	fmt.Print("Digite algo: ")
	fmt.Scanln(&x)

	_, err := strconv.ParseFloat(x, 64)
	if err == nil {
		fmt.Println("A string é um número válido em ponto flutuante.")
	} else {
		fmt.Println("A string não é um número válido em ponto flutuante.")
	}
}
