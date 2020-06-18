package main

import "fmt"

func main() {
	/*1、爬楼梯
	climbStairs(6)*/

	/* 2 最大子序列和
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	a := maxSubArray(arr)
	fmt.Println(a)*/

	// 3 买卖股票的最佳时机
	arr := []int{7, 1, 5, 3, 6, 4}
	a := maxProfit(arr)
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

/**
3、买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
注意：你不能在买入股票前卖出股票。

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

解决方法： 动态规划
1、维护2个变量：截至当前历史最低价，截止当前最大利差，迭代一遍求当天利差，与最大利差相比求较大值

复杂度分析
时间复杂度：O(n)，只需要遍历一次。
空间复杂度：O(1)，只使用了常数个变量。
*/
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {

		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
