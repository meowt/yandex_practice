package main

import "fmt"

func main() {
	var x, y, z, res int
	fmt.Scan(&x, &y, &z)

	if x > 12 || y > 12 || x == y {
		res = 1
	} else {
		res = 0
	}

	fmt.Println(res)
}
