package main

import "fmt"

func main() {
	//1、爬楼梯
	//climbStairs(6)
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	a := maxSubArray(arr)
	fmt.Println(a)

}

/**
 题目描述：假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 解题方法： 动态规划
不难发现，这个问题可以被分解为一些包含最优子结构的子问题，即它的最优解可以从其子问题的最优解来有效地构建，我们可以使用动态规划来解决这一问题。
第 ii 阶可以由以下两种方法得到：
在第 (i-1) 阶后向上爬一阶。
在第 (i-2) 阶后向上爬 22 阶。
所以到达第 ii 阶的方法总数就是到第 (i-1) 阶和第 (i-2)阶的方法数之和。
令 dp[i] 表示能到达第 ii 阶的方法总数：
dp[i]=dp[i-1]+dp[i-2]
dp[i]=dp[i−1]+dp[i−2]

 算法分析：
	复杂度分析

时间复杂度：O(n)，单循环到 nn 。
空间复杂度：O(n)O，dp 数组用了 nn 的空间。

*/
func climbStairs(n int) {
	if n == 1 || n == 2 {
		fmt.Println(n)
		return
	}
	var stairs []int = make([]int, n)
	stairs[0] = 1
	stairs[1] = 2
	for i := 2; i < n; i++ {
		stairs[i] = stairs[i-1] + stairs[i-2]
	}
	fmt.Println(stairs[n-1])
}

/**
	题目描述：
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
解决方案：动态规划
于是可以写出这样的动态规划转移方程：
f(i) = max { f(i - 1) + a[i], a[i]}

算饭分析：
复杂度
时间复杂度：O(n)，其中 n 为 nums 数组的长度。我们只需要遍历一遍数组即可求得答案。
空间复杂度：O(1)。我们只需要常数空间存放若干变量。
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
