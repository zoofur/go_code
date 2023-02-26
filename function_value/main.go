package main

func main() {
	println(helper(nil, 1, 2))
}

//go:noinline
func helper(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}
