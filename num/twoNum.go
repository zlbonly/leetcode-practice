package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	nums := []int{2, 3, 7, 11, 12}
	ret := twoNum(nums, 9)
	fmt.Println("返回他们的数组下标：", ret)
}

/*
创建map映射，用于存放目标数组的相关信息；
遍历目标数组，并获取目标值（target）与数组元素（nums[i]）的差值；
将差值当作map的key，目标数组的角标当作value；
判断map中是否包含，如果包含，则返回map的key为差值的value与i；
如果map中不包含，放入map中。
*/

func twoNum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

/*
复杂度分析
时间复杂度：O(N)O(N)
空间复杂度：O(N)O(N)
*/
