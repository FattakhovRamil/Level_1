package main

import (
	"fmt"
	"sync"
)

func calc(array []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, i := range array {
		sum += i * i
	}
	ch <- sum
}

func main() {
	array := []int{
		2, 4, 6, 8, 10,12,
	}

	var counerGorutin int = 2

	var wg sync.WaitGroup

	stepSize := (len(array) + counerGorutin - 1) / counerGorutin //

	ch := make(chan int, counerGorutin)

	for i := 0; i < len(array); i += stepSize {
		end := i + stepSize
		if end > len(array) {
			end = len(array)
		}
		wg.Add(1)
		go calc(array[i:end], ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalSum := 0
	for c := range ch {
		totalSum += c
	}
	fmt.Println(totalSum)
}
