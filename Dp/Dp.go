package main

import "fmt"

func main() {
	//1、爬楼梯
	climbStairs(6)
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
