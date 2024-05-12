package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ch chan string, wg *sync.WaitGroup, signalShurtDown chan struct{}) {
	defer wg.Done()
	for {
		select {
		case str := <-ch:
			fmt.Printf("%s ", str)
		case <-signalShurtDown:
			return
		}

	}
}

func main() {

	counterGorutin  := flag.Int("workers", 4, "Number of workers") // Кол-во горутин
	//var counterJob int = 10    // works
	flag.Parse()
	if *counterGorutin <= 0 {
		fmt.Println("Error: number of workers must be greater than 0")
        flag.Usage()
        os.Exit(1)
	}

	
	chBuffer := make(chan string, *counterGorutin)
	signalShurtDown := make(chan struct{})
	var wg sync.WaitGroup

	// create workers

	for i := 0; i < *counterGorutin; i++ {
		wg.Add(1)
		go worker(chBuffer, &wg, signalShurtDown)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(signalShurtDown)
	}()

	for {
		select {
		case <-signalShurtDown:
			close(chBuffer)
			wg.Wait()
			fmt.Println("Завершение программы")
			return
		default:
			str := "Hello Go!\n"
			chBuffer <- str
		}
	}

}

//
