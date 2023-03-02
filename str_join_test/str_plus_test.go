package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var s = "这是基准测试，用来测试 + 拼接字符串的性能。\n根据项目来看，每个字符串拼接四次比较符合。\n"

func initStringSlice(n int) []string {
	str := make([]string, n)
	for i := 0; i < n; i++ {
		str = append(str, s)
	}
	return str
}

func BenchmarkBuilder2(b *testing.B) {
	str := initStringSlice(1000)
	b.ResetTimer()
	for i := 0; i< b.N; i++ {
		builder2(str)
	}
}

func builder2(str []string) string {
	b := strings.Builder{}
	l := len(str)
	for i := 0; i < l; i++ {
		b.WriteString(str[i])
	}
	return b.String()
}

func BenchmarkBuffer2(b *testing.B) {
	str := initStringSlice(1000)
	b.ResetTimer()
	for i := 0; i< b.N; i++ {
		buffer2(str)
	}
}

func buffer2(str []string) string {
	b := bytes.Buffer{}
	l := len(str)
	for i := 0; i < l; i++ {
		b.WriteString(str[i])
	}
	return b.String()
}

func BenchmarkSprint2(b *testing.B) {
	str := initStringSlice(1000)
	b.ResetTimer()
	for i := 0; i< b.N; i++ {
		sprint2(str)
	}
}

func sprint2(str []string) string {
	return fmt.Sprint(str)
}

func BenchmarkJoin2(b *testing.B) {
	str := initStringSlice(1000)
	b.ResetTimer()
	for i := 0; i< b.N; i++ {
		join2(str)
	}
}

func join2(str []string) string {
	l := len(str)
	var s string
	for i := 0; i < l; i++ {
		s += str[i]
	}
	return s
}

func BenchmarkStrJoin2(b *testing.B) {
	str := initStringSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strJoin2(str)
	}
}

func strJoin2(str []string) string {
	return strings.Join(str, "")
}