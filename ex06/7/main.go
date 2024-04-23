package main

import (
	"log"
	"time"
)

// Завершение работы горутины с помощью канала time.After()
func main() {
	go someGorutina()
	time.Sleep(6 * time.Second)
}

func someGorutina() {
	timer := time.After(5 * time.Second)
	for {
		select {
		case <-timer:
			log.Println("stoped")
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

//
