package main

import "fmt"

func main() {
	var r, i, c, res int
	fmt.Scan(&r, &i, &c)

	switch i {
	case 0:
		if r != 0 {
			res = 3
		} else {
			res = c
		}
	case 1:
		res = c
	case 4:
		if r != 0 {
			res = 3
		} else {
			res = 4
		}
	case 6:
		res = 0
	case 7:
		res = 1
	default:
		res = i
	}
	fmt.Println(res)
}
