package main

import "testing"

func add(x, y int) int {
	return x + y
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}
