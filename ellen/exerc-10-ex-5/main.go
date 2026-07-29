package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

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

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(v int) {
			canal2 <- trabalho(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(v int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return v
}
