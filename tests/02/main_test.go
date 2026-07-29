package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := Soma(1, 2, 3)
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
