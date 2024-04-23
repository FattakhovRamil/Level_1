package main

import (
	"log"
	"runtime"
	"time"
)

// Завершение работы горутины с помощью runtime.Goexit()
func main() {
	go someGorutina()
	time.Sleep(6 * time.Second)
}

func someGorutina() {
	log.Println("working...")
	time.Sleep(4 * time.Second)
	log.Println("stoped")
	runtime.Goexit()

}

//
