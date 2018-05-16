package main

import "fmt"

// Определяем интерфейс
type Flyer interface {
	Fly()
	Greet()
}

// Структура с соответствующими прилеплиными функциями типо реализация
type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying")
}

func (b Bird) Greet() {
	fmt.Println("Hey there")
}

func main() {
	duckPlane := Bird{"Duck plane"}
	GoFly(duckPlane)
}

func GoFly(f Flyer) {
	f.Fly()
	//b := f.(Bird) пытаемся кастануть интерфейс к типу если не угадали падаем с паникой

	if b, ok := f.(Bird); ok { // безопасное кастование type assertions
		fmt.Println(b.Name)
	}
}
