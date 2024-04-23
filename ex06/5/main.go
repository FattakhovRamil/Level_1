package main

import (
	"log"
	"time"
)

// Завершение работы горутины с помощью переменной и условия
func main() {
	var num int = 16
	go someGorutina(num)
	time.Sleep(4 * time.Second)
}

func someGorutina(num int) {
	if num == 16 {
		log.Println("stoped (true)")
		return
	} else {
		time.Sleep(2 * time.Second)
		log.Println("stoped (false)")
		return
	}

}

//
