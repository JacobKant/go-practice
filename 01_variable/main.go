package main

func main() {
	// интеджеры
	var i int = 10
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1
	var unsignedInt uint = 100500
	var unsignedBigInt uint64 = 1<<64 - 1
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// с правающей точкой
	var p float32 = 3.14
	println("float: ", p)

	var b = true
	println("bool variable ", b)

	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	// символы в одинарных ' ковычках
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)

	// создание новое переменной сокращенно
	meaningOfLive := 42
	println("meaning of life is ", meaningOfLive)

	// приведение типов
	var u1 uint = 17
	var s1 int = 23
	// приведение типа int(), string() приведение к стринг (48 номер символа в таблице символов)
	println(s1 + int(u1))
	println("int to string ", string(48))

	// комплексные числа
	z := 2 + 3i
	println("complex number: ", z)

	// конкатинация строк
	st1 := "Ven9"
	st2 := "Sentybrev"
	fullName := st1 + " " + st2
	println("full name is: ", fullName, len(fullName))

	// ингнор спец символов (строка как есть)
	escaping := `Hello fullName
			\n\t World`
	println(escaping)

	// значения по умолчанию
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values: ", defaultInt, defaultFloat, defaultString, defaultBool)

	// сразу несколько значений
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 = 12
		m2 = "string"
		m3 = 23
	)
	println(m0, m2, m3)
	return
}
