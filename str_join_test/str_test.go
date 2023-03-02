package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// go test -bench=. -benchmem
// go test -bench=. -benchmem -benchtime=3s -cpuprofile=profile.out
func BenchmarkBuilder(b *testing.B) {
	for i := 0; i< b.N; i++ {
		builder()
	}
}

func builder() string {
	b := strings.Builder{}
	b.WriteString("这是基准测试，")
	b.WriteString("用来测试 + 拼接字符串的性能。\n")
	b.WriteString("根据项目来看，")
	b.WriteString("每个字符串拼接四次比较符合。\n")
	return b.String()
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i< b.N; i++ {
		buffer()
	}
}

func buffer() string {
	b := bytes.Buffer{}
	b.WriteString("这是基准测试，")
	b.WriteString("用来测试 + 拼接字符串的性能。\n")
	b.WriteString("根据项目来看，")
	b.WriteString("每个字符串拼接四次比较符合。\n")
	return b.String()
}


func BenchmarkSprint(b *testing.B) {
	for i := 0; i< b.N; i++ {
		sprint()
	}
}

func sprint() string {
	return fmt.Sprint("这是基准测试，",
		"用来测试 + 拼接字符串的性能。\n",
		"根据项目来看，",
		"每个字符串拼接四次比较符合。\n")
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join()
	}
}

func join() string {
	var s string
	s += "这是基准测试，"
	s += "用来测试 + 拼接字符串的性能。\n"
	s += "根据项目来看，"
	s += "每个字符串拼接四次比较符合。\n"
	return s
}

func BenchmarkStrJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strJoin()
	}
}

func strJoin() string {
	str := []string{"这是基准测试，",
		"用来测试 + 拼接字符串的性能。\n",
		"根据项目来看，",
		"每个字符串拼接四次比较符合。\n"}
	return strings.Join(str, "")
}