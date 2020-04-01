package test_demo

import (
	"testing"
)

func BenchmarkByAdd(b *testing.B) {
	s1 := "hello"
	s2 := "world"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byAdd(s1, s2)
	}
}

func BenchmarkBySprintf(b *testing.B) {
	s1 := "hello"
	s2 := "world"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bySprintf(s1, s2)
	}
}