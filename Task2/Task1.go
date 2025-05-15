package main

import (
	"fmt"
)

func main() {
	num := 1
	changeValueWithPointer(&num)
	fmt.Printf("after changed the num's value is %d\n", num)

	someSlice := []int{1, 2, 3, 4, 5}
	changeValueWithSlice(&someSlice)
	fmt.Println(someSlice)

}

func changeValueWithPointer(num *int) {
	*num += 10
}

func changeValueWithSlice(someSlice *[]int) {
	for i, v := range *someSlice {
		(*someSlice)[i] = v * 2
	}
}
