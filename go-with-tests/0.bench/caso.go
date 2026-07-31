package main

import (
	"fmt"
	"time"
)

func main() {

}

func gerarIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i + 1
	}
	return ids
}

func ProcessarLote(ids []int) {
	for _, id := range ids {
		ProcessarID(id)
	}
}

func ProcessarID(id int) {
	time.Sleep(1 * time.Millisecond) // Simula um processamento demorado
	fmt.Printf("Processando ID: %d\n", id)
}
