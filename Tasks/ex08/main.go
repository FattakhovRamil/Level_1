package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBite(num int64, i int, value int) int64 {
	if value == 1 {
		return num | (1 << i) // Операция ИЛИ
	} else {
		return num &^ (1 << i) // Операция И-НЕ
	}
}

func main() {
	var num int64 = 100 
	i := 3              
	value := 1          // Устанавливаемый бит: 1 или 0

	num = setBite(num, i, value)
	fmt.Println(num)

}
