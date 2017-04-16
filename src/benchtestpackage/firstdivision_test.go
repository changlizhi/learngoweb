package benchtestpackage

import (
	"testing"
	"usetest"
)

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usetest.Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		usetest.Division(4, 5)
	}
}
