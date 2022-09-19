package main

import (
	"fmt"
)

func main() {
	var n, temp int
	fmt.Scan(&n)
	var slc []int

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		slc = append(slc, temp)
	}
	pivot := n / 2

	fmt.Println(slc[pivot])
}
