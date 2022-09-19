package main

import "fmt"

func main() {
	var n, i, j, res1, res2, res int
	fmt.Scan(&n, &i, &j)

	if i > j {
		res1 = i - j - 1
		res2 = n - i + j - 1
		if res1 > res2 {
			res = res2
		} else {
			res = res1
		}
	} else {
		res1 = j - i - 1
		res2 = n + i - j - 1
		if res1 > res2 {
			res = res2
		} else {
			res = res1
		}
	}

	fmt.Println(res)
}
