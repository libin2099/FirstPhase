package main

import (
	"fmt"
)

func main() {
	houses := [10]int{20, 300, 7, 60, 2000, 2001, 1500, 40, 39, 115}

	totalAmount := houses[0]
	curIndex := 0
	length := len(houses)
	i := 1
	for {
		if i < length {
			if curIndex+1 == i {
				i++
				continue
			} else {
				if i == length-1 {
					totalAmount += houses[length-1]
					break
				} else {
					if houses[i] > houses[i+1] {
						totalAmount += houses[i]
						curIndex = i
					} else {
						totalAmount += houses[i+1]
						curIndex = i + 1
					}
					i = curIndex + 1
				}
			}
		} else {
			break
		}
	}
	fmt.Println(totalAmount)
}
