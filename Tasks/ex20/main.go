package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func reversStringSpace(str string) string {
	// runes := []rune(str)
	// length := len(runes)

	// for i := 0; i < length/2; i++ {
	// 	runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	// }
	// return string(runes)

	words := strings.Split(str, " ")
	length := len(words)
	for i := 0; i < length/2; i++ {
		words[i], words[length-1-i] = words[length-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	resS := reversStringSpace(s)
	fmt.Println(resS) // sun dog snow
}
