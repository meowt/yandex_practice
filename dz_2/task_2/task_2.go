package main

import (
	"fmt"
	"math"
)

func main() {
	var nums [10]int
	var maxDist int
	minDist := 11
	for i := 0; i < 10; i++ {
		fmt.Scan(&nums[i])
	}
	for i1, v1 := range nums {
		tempDist := 11
		for i2, v2 := range nums {
			if v1 == 1 && v2 == 2 {
				temp := int(math.Abs(float64(i1 - i2)))
				if temp < tempDist {
					tempDist = temp
					fmt.Println(i1, i2, temp, tempDist, minDist, maxDist)
				}
				if tempDist < minDist {
					minDist = tempDist

				}
				if minDist > maxDist {
					maxDist = minDist
				}
			}
		}
	}
	fmt.Println(maxDist)
}
