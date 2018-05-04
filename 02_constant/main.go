package main

const someInt = 1
const typeInt int32 = 17
const fullName = "Vasily"

const (
	flagKey1 = 1
	flagKey2 = 2
)

// iota лютая дичь типо enum с автоинкрементом (whaaaatttt....)
const (
	first = iota // = 0
	one          // = 1
	two          // = 2
	_            // пустая переменная О__О (пропуск 3)
	four         // = 4
)

const (
	_ = iota
	// тута инкремент подставляется в мат выражение и далее выражение как копируется бы с iota + 1
	KB uint = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ошибка ! переполнение типа
	// ZB
)

func main() {
	pi := 3.14
	println(pi + someInt)
	println(first, one, two, four)
	println(KB, MB, GB, TB, PB, EB)
	return
}
