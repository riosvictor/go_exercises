package main

import (
	"fmt"
)

func main() {
	// result, error := sumAndDifference(5, 3)
	result, error := sumAndDifference(-1, 3)
	if error != nil {
		fmt.Println("Erro:", error)
	} else {
		fmt.Println("Soma:", result[0], "Diferença:", result[1])
	}
}

// crie uma função que recebe dois parametros e retorna a soma e a diferença
func sumAndDifference(a int, b int) ([]int, error) {
	if a < 0 || b < 0 {
		return nil, fmt.Errorf("os parametros devem ser positivos")
	}

	sum := a + b
	difference := a - b
	resultado := []int{sum, difference}

	return resultado, nil
}
