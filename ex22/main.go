package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает,
	вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	a := big.NewInt(8388608)      // 2^22
	b := big.NewInt(16777216)     // 2^23
	fmt.Println("Наши числа:", a, b)
	mul := new(big.Int).Mul(a, b) // умножение
	fmt.Println("умножение: ", mul)

	div := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b)) // деление
	fmt.Println("деление: ", div)

	sum := new(big.Int).Add(a, b) // сложение
	fmt.Println("сложение: ", sum)

	diff := new(big.Int).Sub(a, b) // вычитание
	fmt.Println("вычитание: ", diff)
}
