package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	isIncreasing := true
	for i := 1; i < len(str); i++ {
		prev, _ := strconv.Atoi(string(str[i-1]))
		curr, _ := strconv.Atoi(string(str[i]))
		if curr != prev+1 {
			isIncreasing = false
			break
		}
	}

	if isIncreasing {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
