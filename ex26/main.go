package main

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные
	(true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false

*/

import (
	"fmt"

	"strings"
)

func IsUniqueChar(str string) bool {
	strLower := strings.ToLower(str)

	tmp := make(map[rune]struct{}, len(strLower))

	for _, val := range strLower {
		if _, ok := tmp[val]; ok {
			return false
		}
		tmp[val] = struct{}{}
	}
	return true

}

func main() {
	fmt.Println(IsUniqueChar("abcd"))      // true
	fmt.Println(IsUniqueChar("abCdefAaf")) // false
	fmt.Println(IsUniqueChar("aabcd"))     // false

}
