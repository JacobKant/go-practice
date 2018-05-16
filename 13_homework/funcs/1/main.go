package main

import "fmt"

type memoizeFunction func(int ...int) interface{}

// TODO реализовать
// какова хера функции принимают int, ...int нахера эти ...int я так и не понял, беру первое значение
var fibonacci memoizeFunction = func(int ...int) interface{} {
	i := 2
	temp := 1
	n1 := 1
	n2 := 1
	for {
		if i == int[0] || int[0] == 1 || int[0] == 2 {
			return interface{}(n2)
		}
		temp = n2
		n2 = n1 + n2
		n1 = temp
		i++
	}
}

// ???????????? что должна делать эта функция ????????????
var romanForDecimal memoizeFunction = func(int ...int) interface{} {
	return interface{}("????")
}

//TODO Write memoization function
func memoize(function memoizeFunction) memoizeFunction {
	var cache = make(map[int]interface{})
	return func(int ...int) interface{} {
		if value, found := cache[int[0]]; found {
			return value
		}
		value := function(int...)
		cache[int[0]] = value
		return value
	}
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
	fibonacci = memoize(fibonacci)
}

func main() {
	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}
