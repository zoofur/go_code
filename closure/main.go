package main

func main() {
	println(mc(1))
}

//go:noinline
func mc(n int) func() int {
	return func() int {
		return n
	}
}
