package main

import "fmt"

//The Answer is "true" since (false&&false) is false
// and so its negation is true. Bue here is program
//that verifies that

func main() {

	ans := (true && false) || (false && true) || !(false && false)
	fmt.Println(ans)
}
