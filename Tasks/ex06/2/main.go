package main

import (
	"context"
	"log"
	"time"
)

// Завершение работы горутины с помощью context.WithDeadline
func main() {
	//отложенное завершение через 10 сек
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()

	someGorutina(ctx)

}

func someGorutina(ctx context.Context) {
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Println("начала работы..")
	for {
		select {
		case <-newCtx.Done():
			log.Println("ctx done: ", ctx.Err())
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

//
