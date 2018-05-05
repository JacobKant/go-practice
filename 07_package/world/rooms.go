package world

import (
	"fmt"
)

// экспортируется из пакета (с большой буквы)
func PrintHello() {
	fmt.Println("Hello")
}

// не экспортируется из пакета (с маленькой буквы)
func printWorld() {
	fmt.Println("World")
}
