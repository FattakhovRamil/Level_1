/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
)

func main() {
	chFirst := make(chan int)  // 1 ch
	chSecond := make(chan int) // 2 ch
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go func(array []int, chFirst chan<- int) {
		defer close(chFirst)
		for _, val := range array {
			chFirst <- val
		}
	}(array, chFirst)

	go func(chFirst <-chan int, chSecond chan<- int) {
		defer close(chSecond)
		for val := range chFirst {
			chSecond <- val * 2
		}
	}(chFirst, chSecond)

	for val := range chSecond {
		fmt.Println(val)
	}
}
