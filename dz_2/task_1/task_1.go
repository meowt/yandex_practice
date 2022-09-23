package main

import "fmt"

func main() {
	nums := map[int]int{}
	var max int
	temp := 1
	for temp != 0 {
		fmt.Scan(&temp)
		if temp == 0 {
			break
		}
		if temp > max {
			max = temp
		}
		nums[temp]++
	}

	fmt.Println(nums[max])
}
