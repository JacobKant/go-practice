package main

import (
	"fmt"
)

func main() {
	// объявление без инициализации
	var mm map[string]string
	fmt.Println(mm)

	// объявление с инициализацией
	var mm1 = map[string]string{}
	mm2 := map[string]string{}
	var mm3 = make(map[string]string)
	mm1["test"] = "ok"
	mm2["test"] = "ok"
	mm3["test"] = "ok"
	fmt.Println(mm1, mm2, mm3)

	// если обратиться к элементу которого нет вернется значение по умолчанию
	name := mm3["name"]
	fmt.Println(name)

	// можно инициализировать значение и признак наличия значения в мапе
	name, exist := mm3["name"]
	_, existOnly := mm3["name"]
	fmt.Println(name, exist, existOnly)
}
