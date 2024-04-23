package main

import (
	"context"
	"log"
	"time"
)

// Завершение работы горутины с помощью context.WithCancel (отменой)
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go someGorutina(ctx)

	time.Sleep(7 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func someGorutina(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			log.Println("stoped")
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

//
