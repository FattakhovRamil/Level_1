package main

import (
	"log"
	"sync"
	"time"
)

// Завершение работы горутины с помощью sync.WaitGroup
func main() {
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go someGorutina(&wg)
		}
	}()

	
	wg.Wait()
	

}

func someGorutina(wg *sync.WaitGroup) {
   wg.Done()
	time.Sleep(5 * time.Second)
	log.Println("working...")
}

//
