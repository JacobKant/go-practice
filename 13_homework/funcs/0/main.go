package main

import (
	"fmt"
	"math"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	if x <= 0 {
		return 0.0
	}
	var a = 1.0
	for {
		var na = 0.5 * (a + x/a)
		if math.Abs(a-na) < 1e-9 {
			return na
		}
		a = na
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
