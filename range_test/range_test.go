package main

import "testing"

func initSlice(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = "基准测试for range"
	}
	return slice
}

func BenchmarkRange(b *testing.B) {
	slice := initSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k, v := range slice {
			_, _ = k, v
		}
	}
}

func BenchmarkRangeNoV(b *testing.B) {
	slice := initSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(slice); i++ {
			_, _ = i, slice[i]
		}
	}
}
