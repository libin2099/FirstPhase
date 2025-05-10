package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(&result, nums, 0)
	return result
}

func backtrack(result *[][]int, nums []int, start int) {
	if start == len(nums) {
		// 复制当前排列到结果中
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		// 交换元素
		nums[start], nums[i] = nums[i], nums[start]
		// 递归处理下一个位置
		backtrack(result, nums, start+1)
		// 回溯，恢复交换
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func main() {
	// 测试用例
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	fmt.Println(permutations)

	// 另一个测试用例
	nums2 := []int{0, 1}
	permutations2 := permute(nums2)
	fmt.Println(permutations2)
}
