package main

import "fmt"

func grt(args ...int) int {
	grt := 0
	for _, s := range args {
		if s > grt {
			grt = s
		}
	}
	return grt
}

func main() {
	fmt.Println(grt(3, 6, 24, 13, 8, 50, 7, 10))
}
