package main

func main() {
	i := 5
	b := &i // получаем ссылку(указатель) на переменную
	println(i)
	*b = 10
	println(i)

}
