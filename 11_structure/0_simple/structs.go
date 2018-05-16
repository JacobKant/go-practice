package main

import "fmt"

// определяем структуру (тип)(недо класс)
type example struct {
	Flag    bool
	counter int16
	pi      float32
}

func main() {
	// определяем переменную с типом этой структуры, все поля будут проинициализированы по умолчанию
	var e1 example

	// через %+v можно распечатать ее красиво
	fmt.Printf("%+v", e1)

	// что бы поля проставить нужные можно через скобочки вот так
	e2 := example{
		Flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// тожесамое что первый пример
	e3 := example{}
	fmt.Println(e3)

	// стучимся к полям через точку
	fmt.Println("Flag", e2.Flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
