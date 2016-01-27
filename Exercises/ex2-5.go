package main

import "fmt"

func foo(args ...int) int {
	sum := 0
	for _, s := range args {
		sum += s
	}
	return sum
}

func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())
}
