package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	isDecreasing := true
	for i := 1; i < len(str); i++ {
		prev, _ := strconv.Atoi(string(str[i-1]))
		curr, _ := strconv.Atoi(string(str[i]))
		if curr != prev-1 {
			isDecreasing = false
			break
		}
	}

	if isDecreasing {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}
