package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor recebido", v)
	}
}

func envia(par, impar chan int) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func recebe(par, impar, converge chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range par {
			converge <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range impar {
			converge <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(converge)
}
