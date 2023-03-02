package main

import (
	"strings"
	"testing"
)

func builder3(str []string) string {
	b := strings.Builder{}
	l := len(str)

	n := 0
	for i := 0; i < l; i++ {
		n += len(str[i])
	}
	b.Grow(n)

	for i := 0; i < l; i++ {
		b.WriteString(str[i])
	}
	return b.String()
}

func BenchmarkPreGrowBuilder10(b *testing.B) {
	str := initStringSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder3(str)
	}
}

func BenchmarkPreGrowBuilder100(b *testing.B) {
	str := initStringSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder3(str)
	}
}

func BenchmarkPreGrowBuilder1000(b *testing.B) {
	str := initStringSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder3(str)
	}
}