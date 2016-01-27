package main

import "fmt"

func main() {
	var fname, lname string
	fmt.Println("Enter your first name and then your last name.")
	fmt.Scan(&fname)
	fmt.Scan(&lname)

	fmt.Println("Hello", fname, lname)
}
