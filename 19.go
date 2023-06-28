package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&str)

	isNumeric := regexp.MustCompile("^[0-9]+$").MatchString(str)

	if isNumeric {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}
