package main

import (
	"fmt"
)

func main() {
	str := "编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。"
	strrunes := []rune(str)
	length := len(strrunes)
	for i := 0; i < length/2; i++ {
		temp := strrunes[i]
		strrunes[i] = strrunes[length-i-1]
		strrunes[length-i-1] = temp
	}

	fmt.Println(string(strrunes))
}
