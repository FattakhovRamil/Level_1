package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Реализовать конкурентную запись данных в map.
func main() {

	myMap := make(map[int]int)

	var wg sync.WaitGroup
	var mu sync.Mutex
	//ch := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {

		wg.Add(1)
		go worker(myMap, &wg, &mu)

	}
	wg.Wait()
	fmt.Println(myMap)
}

func worker(myMap map[int]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := 10; i < 20; i++ {
		mu.Lock()
		myMap[i] = i * i
		mu.Unlock()
	}

}
