package main

import (
	"fmt"
	"sync"
)

/*
Além da goroutine principal, crie duas outras goroutines.
Cada goroutine adicional devem fazer um print separado.
Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

var wg sync.WaitGroup

func main() {
	createGoRoutine(100)
	wg.Wait()
}

func createGoRoutine(num int) {
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(n int) {
			fmt.Printf("Hello from routine %d\n", n)
			wg.Done()
		}(i + 1)
	}
}
