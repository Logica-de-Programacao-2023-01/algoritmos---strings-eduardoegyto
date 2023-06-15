package main

import (
	"fmt"
	"sort"
	"strings"
)
func anagrama(s string) string {
	s = strings.ToLower(s)
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func main() {
	var x, y string
	fmt.Print("Digite a primeira palavra: ")
	fmt.Scanln(&x)
	fmt.Print("Digite a segunda palavra: ")
	fmt.Scanln(&y)

	x = anagrama(x)
	y = anagrama(y)

	if x == y {
		fmt.Println("As palavras são anagramas.")
	} else {
		fmt.Println("As palavras não são anagramas.")
	}
}

