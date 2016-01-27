package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 334; i++ {
		sum = sum + 3*i
	}
	for j := 1; j < 200; j++ {
		sum = sum + 5*j
	}
	for k := 1; k < 67; k++ {
		sum = sum - 15*k
	}

	fmt.Println(sum)
}
