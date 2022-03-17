package main

import (
	"testing"
)

const numOfSimulations = 100

func BenchmarkSimulate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulate(numOfSimulations)
	}
}
