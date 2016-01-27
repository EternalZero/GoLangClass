package main

import "fmt"

func main() {
	var numb int
	half := func(num int) (int, bool) {
		return num / 2, num%2 == 0
	}
	fmt.Println("Enter a number.")
	fmt.Scan(&numb)

	fmt.Println(half(numb))
}
