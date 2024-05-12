package main

/* Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/
import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <N>")
		os.Exit(1)
	}

	N, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Invalid argument for N")
		os.Exit(0)
	}

	ch := make(chan int)
	done := make(chan struct{})

	var wg sync.WaitGroup
	timer := time.NewTimer(time.Duration(N) * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case ch <- i:
				time.Sleep(500 * time.Millisecond)
			case <-done:
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case num := <-ch:
				fmt.Printf("%d ", num)
			case <-timer.C:
				fmt.Println("\nВремя истекло, программа завершается")
				close(ch)
				close(done)
				return
			}
		}
	}()
	wg.Wait()
}

//
