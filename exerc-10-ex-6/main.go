package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// cadencia a quantidade de go routines que vão consumir o canal
	// evitando pico de cpU e memória
	steps := 5
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(100, canal1)
	go outra(steps, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(total int, canal chan int) {
	for i := 0; i < total; i++ {
		canal <- i
	}
	close(canal)
}

func outra(steps int, canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < steps; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(canal2)
}

func trabalho(v int) int {
	time.Sleep(time.Millisecond * 1000)
	return v
}
