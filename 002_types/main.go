package main

import (
	"fmt"
)

func main() {
	var astring = "Baymax"
	var anum = 123

	fmt.Printf("%T\n", astring)
	fmt.Printf("%T\n", anum)

	types := []interface{}{"j", 7, 4.3, true}
	for _, v := range types {
		fmt.Printf("%T\n", v)
	}
}
