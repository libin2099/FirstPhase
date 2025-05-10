package main

import "fmt"

type LinkedList struct {
	pre   *LinkedList
	value int
	next  *LinkedList
}

func main() {
	// 构造第一个链表数据
	currentFirst := &LinkedList{nil, 1, nil}
	current1 := &LinkedList{currentFirst, 3, nil}
	currentFirst.next = current1
	current2 := &LinkedList{current1, 7, nil}
	current1.next = current2
	current1 = &LinkedList{current2, 8, nil}
	current2.next = current1
	current2 = &LinkedList{current1, 9, nil}
	current1.next = current2
	current1 = &LinkedList{current2, 998, nil}
	current2.next = current1

	// 构造第二个链表数据
	currentSecond := &LinkedList{nil, 2, nil}
	current3 := &LinkedList{currentSecond, 4, nil}
	currentSecond.next = current3
	current4 := &LinkedList{current3, 5, nil}
	current3.next = current4
	current3 = &LinkedList{current4, 6, nil}
	current4.next = current3
	current4 = &LinkedList{current3, 10, nil}
	current3.next = current4
	current3 = &LinkedList{current4, 100, nil}
	current4.next = current3
	current4 = &LinkedList{current3, 999, nil}
	current3.next = current4

	// 对两个链表混合并升序排序
	result := sort(currentFirst, currentSecond)
	// ***特别注意：如果一个指针变量被分配内存地址，那么即使将它赋值给另一个变量或做为参数传入函数，都会报错，这点真的不同于java
	//var currentFirst1 *LinkedList
	//result := sort(currentFirst1, currentSecond)
	// 打印合并并排序后的结果
	for cur := result; cur != nil; cur = cur.next {
		fmt.Println(cur.value)
	}

}

func sort(list1, list2 *LinkedList) (result *LinkedList) {

	var current *LinkedList = nil
	current1 := list1
	current2 := list2

	// 是否首次进入，为了是将第一个元素的指针用于返回值
	isFirst := true
	for {
		if current1 == nil || current2 == nil {
			if current1 == nil && current2 != nil {
				current.next = current2
				current2.pre = current
			} else if current1 != nil && current2 == nil {
				current.next = current1
				current1.pre = current
			}
			break
		}
		if current1.value <= current2.value {
			if current == nil {
				current = current1
			} else {
				current.next = current1
				current1.pre = current
				current = current.next
			}
			current1 = current1.next
		} else {
			if current == nil {
				current = current2
			} else {
				current.next = current2
				current2.pre = current
				current = current.next
			}

			current2 = current2.next
		}

		if isFirst {
			isFirst = false
			result = current
		}
	}

	return
}
