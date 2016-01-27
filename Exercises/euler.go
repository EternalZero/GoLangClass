package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3
	sum := 2
	for c < 4000001 {
		if c%2 == 0 {
			sum += c
		}
		c = a + b
		a = b
		b = c
	}
	fmt.Println(sum)
}
