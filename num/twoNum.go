package main

import (
	"fmt"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	/*nums := []int{2, 3, 7, 11, 12}
	ret := twoNum(nums, 9)
	fmt.Println("返回他们的数组下标：", ret)*/

	nums := []int{4, 7, 2, 10, 5}
	k := 5
	cutNum(nums, k)
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

/**
	题目描述
给定长度为n的数组，每个元素代表一个木头的长度，木头可以任意截断，从这堆木头中截出至少k个相同长度为m的木块。已知k，求max(m)。

ps: 截断的长度必须是整数

思路： 暴力破解法
大概思路就是从1遍历到木棍最长的长度，每次遍历的长度作为m，如果可以将所有木头截出来k个长度为m的木块，则更新最大值，最后输出最大值即可

分析：上面的代码也比较容易理解，这里就不多展开说了。时间复杂度也很容易看出来是O(n * len), len为木头中最大的长度。容易想到遍历长度时可以从大到小遍历，if (cnt >= k)成立，则该值即为最终结果，可直接break，但最坏时间复杂度没变。
*/

func cutNum(nums []int, k int) {

	//n := max(nums[0 ~ n-1])
	maxLength := 10

	res := 0
	for m := 1; m <= maxLength; m++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			cnt += nums[j] / m
		}

		fmt.Println(cnt, m)

		if cnt >= k {
			res = max(res, m)
		}
	}

	fmt.Println(res)

}

/**
切木棍二分法

*/
func cutNumBinarySearch(nums []int, k int) int {

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
