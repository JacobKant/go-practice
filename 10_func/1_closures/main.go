package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := getTimer()

	// объявляем анонимную функцию
	f := func() {
		time.Sleep(1000)
		myTimer()
	}

	// экзикутим анонимную функцию
	f()
}

// Оо да... функция возвращающая функцию
func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v", time.Since(start))
	}
}
