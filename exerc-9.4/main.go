package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
resolva o race com mutex

// go run exerc-9.4/main.go
// go run -race exerc-9.4/main.go
*/

var wg sync.WaitGroup
var mutex sync.Mutex
var contador int = 0

func main() {
	createGoRoutine(500)
	wg.Wait()
	fmt.Println("Valor final do contador:", contador)
}

func createGoRoutine(num int) {
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			mutex.Lock()

			num := contador
			runtime.Gosched()
			num++
			contador = num
			mutex.Unlock()

			wg.Done()
		}()
	}
}
