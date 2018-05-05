package main

func main() {
	a := true
	if a {
		println("hello")
	}

	b := 1
	if b == 1 {
		println("b == 1")
	}

	mm := map[string]string{
		"firstName": "Vasya",
		"lastName":  "Petrov",
	}
	if firstName, ok := mm["firstName"]; ok {
		println("firstName key exist, = ", firstName)
	} else {
		println("no firstname")
	}

	// бесконечный фор
	for {
		println("бесконечный цикл")
		break // выход из цикла
	}

	// фор только с условием
	i := 0
	for i < 4 {
		if i < 2 {
			i++
			println("<2")
			continue
		}
		println("<4")
		i++
	}

	// стандартный фор
	slice1 := []int{1, 2, 3, 4, 5}
	for ix := 0; ix < len(slice1); ix++ {
		println("c style loop", ix, len(slice1))
	}

	// типо foreach в java. пробег по всех элементам из слайса
	for ind, val := range slice1 {
		println("range style by index and value", ind, val)
	}

	// можно по мапе
	mape := map[string]int{
		"hello": 1,
		"world": 2,
	}
	for key, value := range mape {
		println("range style by key value map", key, value)
	}

	// switch.
	switch "vasya" {
	case "petysa":
		{
			// в отличии от других языков не проваливаемся по умолчанию. break не нужен
			println("petysa")
		}
	case "vasya", "huyasa": // сразу несколько значений в кейсе
		{
			println("vasya")
			break // но никто не запрещает использовать  break что бы выйти из кейса
		}
		fallthrough // указываем что хотим провалиться
	default:
		{
			println("default")
		}
	}

}
