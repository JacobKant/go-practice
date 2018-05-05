package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {}

//ReturnInt Типо умею в инты
func ReturnInt() int {
	return 1
}

//ReturnFloat Типо умею в флоуты
func ReturnFloat() float32 {
	return 1.1
}

//ReturnIntArray Типо сделать массив
func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

//ReturnIntSlice Типо умею сделать слайс
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

//IntSliceToString жесть...
func IntSliceToString(numbers []int) string {
	valuesText := []string{}
	for i := range numbers {
		number := numbers[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	return strings.Join(valuesText, "")
}

//MergeSlices объединение двух массивов
func MergeSlices(slice1 []float32, slice2 []int32) []int {
	result := make([]int, 0, len(slice1)+len(slice2))
	for _, value := range slice1 {
		result = append(result, int(value))
	}

	for _, value := range slice2 {
		result = append(result, int(value))
	}
	return result
}

//GetMapValuesSortedByKey получаем сортированный слайс вэлью по кеям
func GetMapValuesSortedByKey(mm map[int]string) []string {
	keys := []int{}

	for key := range mm {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result := []string{}
	for _, value := range keys {
		result = append(result, mm[value])
	}

	return result

}
