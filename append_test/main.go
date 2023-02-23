package main

import "fmt"

func main() {
	s1 := make([]int, 0, 4)
	fmt.Printf("%p\n", s1)
	var appendFunc = func(s *[]int) {
		*s = append(*s, 1, 2, 3)
		fmt.Printf("%p\n", s)
		*s = append(*s, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
		fmt.Println(s)
		fmt.Printf("%p\n", s)
	}
	fmt.Println(s1)
	appendFunc(&s1)
	fmt.Println(s1)
	fmt.Println(s1[:4])
}
