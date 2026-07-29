package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := soma(1, 2, 3)
	esperado := 6

	if resultado != esperado {
		t.Errorf("soma(1, 2, 3); expected %d, got %d", esperado, resultado)
	}
}

func TestMultiplica(t *testing.T) {
	resultado := multiplica(10, 10)
	esperado := 100

	if resultado != esperado {
		t.Errorf("multiplica(10, 10); expected %d, got %d", esperado, resultado)
	}
}

// Teste falho para demonstrar o funcionamento do teste
// func TestFailure(t *testing.T) {
// 	resultado := soma(1, 2, 3)
// 	esperado := 7

// 	if resultado != esperado {
// 		t.Errorf("soma(1, 2, 3); expected %d, got %d", esperado, resultado)
// 	}
// }
