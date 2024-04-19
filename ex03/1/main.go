package main

import (
	"fmt"
)

func calc(ch chan int, array []int) {
	defer close(ch)
	for _, i := range array {
		ch <- i * i
	}
}

func main() {
	array := []int{
		2, 4, 6, 8, 10,
	}

	ch := make(chan int)
	go calc(ch, array)

	var counter int
	for i := range ch {
		counter += i
	}
	fmt.Println(counter)
}
