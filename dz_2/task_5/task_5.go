package main

import (
	"fmt"
)

func main() {
	var d, x, y, res, m int
	fmt.Scan(&d, &x, &y)
	var dist []int
	dist = append(dist, (x*x + y*y), ((x-d)*(x-d) + y*y), (x*x + (y-d)*(y-d)))
	if (0 <= x) && (y >= 0) && (x+y <= d) {
		res = 0
	} else {
		for i, v := range dist {
			if i == 0 || v < m {
				m = v
				res = i + 1
			}
		}
	}
	fmt.Println(res)
}
