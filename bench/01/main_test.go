package main

import (
	"fmt"
	"testing"
)

func ExampleSoma() {
	fmt.Println(Soma(1, 2, 3))
	// Output: 6
}

func ExampleSoma_multiple() {
	fmt.Println(Soma(1, 2, 3))
	fmt.Println(Soma(1, 2, 4))
	fmt.Println(Soma(1, 2, 5))
	// Output: 6
	// 7
	// 8
}

func TestSoma(t *testing.T) {
	resultado := Soma(1, 2, 3)
	esperado := 6

	if resultado != esperado {
		t.Errorf("soma(1, 2, 3); expected %d, got %d", esperado, resultado)
	}
}

func TestMultiplica(t *testing.T) {
	resultado := Multiplica(10, 10)
	esperado := 100

	if resultado != esperado {
		t.Errorf("Multiplica(10, 10); expected %d, got %d", esperado, resultado)
	}
}

type TestCase struct {
	input    []int
	expected int
}

func TestSomaTableDriven(t *testing.T) {
	testCases := []TestCase{
		{input: []int{1, 2, 3}, expected: 6},
		{[]int{4, 5, 6}, 15},
		{[]int{-1, -2, -3}, -6},
	}

	for _, tc := range testCases {
		resultado := Soma(tc.input...)
		if resultado != tc.expected {
			t.Errorf("soma(%v); expected %d, got %d", tc.input, tc.expected, resultado)
		}
	}
}

// Benchmarks
func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 2, 3)
	}
}

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(10, 10)
	}
}

// to run:
// go test -bench=. -benchmem

// to run coverage:
// go test -cover
//
// to run coverage with html report:
// go test -coverprofile=coverage.out // gera o arquivo de cobertura
// go tool cover -html=coverage.out // abre o relatório em html
