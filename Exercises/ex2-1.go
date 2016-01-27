package main

import "fmt"

func half(num int) (int, bool) {
	return num / 2, num%2 == 0
}

func main() {
	var numb int
	fmt.Println("Enter a number.")
	fmt.Scan(&numb)

	fmt.Println(half(numb))
}
