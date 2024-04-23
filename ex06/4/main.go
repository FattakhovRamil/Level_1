package main

import (
	"log"
	"time"
)

// Завершение работы горутины с помощью канала
func main() {
	ch := make(chan struct{})
	defer close(ch)
	go someGorutina(ch)

	time.Sleep(7 * time.Second)
	ch <- struct{}{}

}

func someGorutina(ch chan struct{}) {

	for {
		select {
		case <-ch:
			log.Println("stoped")
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

//
