package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	Num  int
	Name string
}

// алиас типа
type MyInt int
type withFiles bool

// так можно прилепить функцию к типу создать подобие "метода"
func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

type mySliceStruct []MyStruct

// прокачиваем слайс
func (m mySliceStruct) Less(i int) bool { return false }
func (m mySliceStruct) Len() int        { return 1 }
func (m mySliceStruct) Swap(i, j int)   {}

func main() {
	i := MyInt(0)

	i.add(3)
	i.showYourSelf()
}

func sorter(sl []MyStruct) []MyStruct {
	sort.Slice(sl, func(i, j int) bool { return true })
	return sl
}
