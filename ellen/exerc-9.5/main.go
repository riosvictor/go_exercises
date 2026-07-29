package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
resolva o race com mutex

// go run exerc-9.5/main.go
// go run -race exerc-9.5/main.go
*/

var wg sync.WaitGroup
var contador int64 = 0

func main() {
	createGoRoutine(500)
	wg.Wait()
	fmt.Println("Valor final do contador:", contador)
}

func createGoRoutine(num int) {
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
}
