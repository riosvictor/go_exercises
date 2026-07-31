package main

import "testing"

func BenchmarkProcessarLote(b *testing.B) {
	ids := gerarIDs(1000) // Preparação do cenário

	b.ResetTimer() // ⏱️ Zera o cronômetro para não contar a preparação

	for i := 0; i < b.N; i++ {
		ProcessarLote(ids)
	}
}

// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out
// brew install graphviz
// go tool pprof -http=:8080 mem.out
