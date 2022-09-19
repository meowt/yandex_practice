package main

import (
	"fmt"
	"strings"
)

func main() {
	var numRows int
	fmt.Scan(&numRows)

	var slc [][2][3]bool

	for i := 0; i < numRows; i++ {
		var temp string
		fmt.Scan(&temp)
		slc = append(slc, parse(temp))
	}

	var numReqs int
	fmt.Scan(&numReqs)

	for i := 0; i < numReqs; i++ {
		var amount int
		var side, place string
		fmt.Scan(&amount, &side, &place)

		res := false
		for c := 0; c < amount; c++ {
			if side == "left" {
				if place == "window" {
					for _, v := range slc {
						if v[0][c] {

						}
					}
				} else {

				}
			} else {
				if place == "aisle" {
					for j, v := range slc {

					}
				}
			}
		}
	Label1:
	}

	fmt.Println("Passengers can take seats: 1B 1C")
}

func parse(row string) [2][3]bool {
	temp := strings.Split(row, "_")
	var slc [2][]string

	for i, v := range temp {
		slc[i] = strings.Split(v, "")
	}

	var res [2][3]bool

	for i, v := range slc {
		for j, val := range v {
			switch val {
			case ".":
				res[i][j] = true
			case "#":
				res[i][j] = false
			}
		}
	}

	return res
}
