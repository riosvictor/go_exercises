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
	wg.Add(2)
	go routine1()
	go routine2()
	wg.Wait()
}

func routine1() {
	fmt.Println("Hello from routine 1")
	wg.Done()
}
func routine2() {
	fmt.Println("Hello from routine 2")
	wg.Done()
}
