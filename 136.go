package main

import (
	"fmt"
)

func main() {
	digits := [11]int{1, 2, 3, 4, 5, 2, 4, 100, 1, 5, 3}
	index, value := findUniqueDigit(digits)
	fmt.Printf("下标为%d的元素：%d\n即为没有重复的元素", index, value)
}

func findUniqueDigit(digits [11]int) (index, value int) {
	// 存储了哪个位置已经有相同的数字了
	countMap := make(map[int]bool)
	length := len(digits)
	for i := 0; i < length; i++ {
		_, e := countMap[i]
		if e {
			continue
		}
		exist := false
		for j := i + 1; j < length; j++ {
			if digits[i] == digits[j] {
				countMap[j] = true
				exist = true
				break
			}
		}
		if !exist {
			index = i
			value = digits[i]
			break
		}
	}
	return
}
