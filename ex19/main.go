package main

import "fmt"

/*
	Разработать программу, которая переворачивает
	подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/

func reversString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)

}

func main() {
	s := "Привет, мир! 😊"
	resS := reversString(s)
	fmt.Println(resS) // 😊 !рим ,тевирП
}
