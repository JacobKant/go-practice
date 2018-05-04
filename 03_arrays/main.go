package main

import (
	"fmt"
)

func main() {
	// init
	var a1 [3]int
	fmt.Println("массив ", a1, " ", len(a1))

	// многомерные массивы
	var matrix [3][3][3]int
	matrix[1][1][1] = 999
	fmt.Println(matrix)

	return
}
