package main

import (
	"fmt"
)

func main() {
	var y int
	var x int
	var n int
	var sum int
	fmt.Scanln(&y)
	for i := 0; i < y; i++ {

		sum = 0
		fmt.Scanln(&x)
		for a := 0; a < x; a++ {

			fmt.Scanln(&n)
			if n < 0 {

				continue
			}
			sum += n * n
		}
		fmt.Println(sum)
	}
}
