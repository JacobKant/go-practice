package main

import (
	"fmt"
)

func main() {
	var sl []int // как массив без типа
	fmt.Println("пустой слайс интов", sl)

	sl = append(sl, 100) // добавление элемента
	fmt.Println("не пустой слайс интов", sl)
	fmt.Println("длина слайса ", len(sl), "капасити слайса (длинна массива) ", cap(sl))

	// короткая запись
	sl2 := []int{
		10, 20, 30,
	}
	fmt.Println(sl2)

	// добавить слайс в слайс
	sl = append(sl, sl2...)
	fmt.Println(sl)

	// создать слайс с нужным капасити (длинной внутреннего массива)
	slice3 := make([]int, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// создать слайс с нужной длинной и капасти
	slice4 := make([]int, 10, 20)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// внутри слайса ссылка на массив, она копируется просто присвоить
	slice4 = slice3
	slice4[0] = 99999
	fmt.Println(slice4, slice3)

	// копирование слайса
	slice7 := make([]int, len(slice4), len(slice4))
	copy(slice7, slice4)
	fmt.Println(slice7, slice4)

	// часть слайса (невключая границ)
	fmt.Println(slice7[1:4], slice7[:4], slice7[2:])
}
