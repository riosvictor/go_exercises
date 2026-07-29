package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Utilizando goroutines, crie um programa incrementador:
Tenha uma variável com o valor da contagem
Crie um monte de goroutines, onde cada uma deve:
Ler o valor do contador
Salvar este valor
Fazer yield da thread com runtime.Gosched()
Incrementar o valor salvo
Copiar o novo valor para a variável inicial
Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
Demonstre que há uma condição de corrida utilizando a flag -race

// go run exerc-9.3/main.go
// go run -race exerc-9.3/main.go
*/

var wg sync.WaitGroup
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
			num := contador
			runtime.Gosched()
			num++
			contador = num
			wg.Done()
		}()
	}
}
