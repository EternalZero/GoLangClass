package main

import "fmt"

func main() {
	var smaller int
	var larger int

	fmt.Println("Enter 2 different positive integers, the smaller number first.")
	fmt.Scan(&smaller)
	fmt.Scan(&larger)

	if smaller < larger {
		fmt.Println("The remainder of the larger divided by the smaller is:", larger%smaller)
	} else {
		fmt.Println("Your first number was not the smaller number.")

	}
}
