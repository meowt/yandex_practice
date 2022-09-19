package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var slc []int
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		slc = append(slc, temp)
	}

	min := slc[0]
	max := slc[len(slc)-1]
	maxDif := max - min

	for i := 0; i < len(slc)-1; i++ {
		if slc[i] > slc[i+1] {
			maxDif = -1
			break
		}
	}

	fmt.Println(maxDif)
}
