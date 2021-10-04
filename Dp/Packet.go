package main

import (
	"math"
)

/**
一、背包问题总结
	1、背包问题：
		背包问题是动态规划非常重要的一类问题，它有很多变种，
		但题目千变万化都离不开我根据力扣上背包问题的题解和一些大佬的经验总结的解题模板

	2、背包定义：
		那么什么样的问题可以被称作为背包问题？换言之，我们拿到题目如何透过题目的不同包装形式看到里面背包问题的不变内核呢？
		我对背包问题定义的理解：
		给定一个背包容量target，再给定一个数组nums(物品)，能否按一定方式选取nums中的元素得到target

		注意：
		1、背包容量target和物品nums的类型可能是数，也可能是字符串
		2、target可能题目已经给出(显式)，也可能是需要我们从题目的信息中挖掘出来(非显式)(常见的非显式target比如sum/2等)
		3、选取方式有常见的一下几种：每个元素选一次/每个元素选多次/选元素进行排列组合
		那么对应的背包问题就是下面我们要讲的背包分类

	3、背包问题分类：
		常见的背包类型主要有以下几种：
			1、0/1背包问题：每个元素最多选取一次
			2、完全背包问题：每个元素可以重复选择
			3、组合背包问题：背包中的物品要考虑顺序
			4、分组背包问题：不止一个背包，需要遍历每个背包

		而每个背包问题要求的也是不同的，按照所求问题分类，又可以分为以下几种：
			1、最值问题：要求最大值/最小值
			2、存在问题：是否存在…………，满足…………
			3、组合问题：求所有满足……的排列组合

		因此把背包类型和问题类型结合起来就会出现以下细分的题目类型：
			1、0/1背包最值问题
			2、0/1背包存在问题
			3、0/1背包组合问题
			4、完全背包最值问题
			5、完全背包存在问题
			6、完全背包组合问题
			7、分组背包最值问题
			8、分组背包存在问题
			9、分组背包组合问题
			这九类问题我认为几乎可以涵盖力扣上所有的背包问题

	4、背包问题解题模板
		背包问题大体的解题模板是两层循环，分别遍历物品nums和背包容量target，然后写转移方程，
		根据背包的分类我们确定物品和容量遍历的先后顺序，根据问题的分类我们确定状态转移方程的写法

			首先是背包分类的模板：
			1、0/1背包：外循环nums,内循环target,target倒序且target>=nums[i];
			2、完全背包：外循环nums,内循环target,target正序且target>=nums[i];
			3、组合背包(考虑顺序)：外循环target,内循环nums,target正序且target>=nums[i];
			4、分组背包：这个比较特殊，需要三重循环：外循环背包bags,内部两层循环根据题目的要求转化为1,2,3三种背包类型的模板

		然后是问题分类的模板：
		1、最值问题: dp[i] = max/min(dp[i], dp[i-nums]+1)或dp[i] = max/min(dp[i], dp[i-num]+nums);
		2、存在问题(bool)：dp[i]=dp[i]||dp[i-num];
		3、组合问题：dp[i]+=dp[i-num];

		这样遇到问题将两个模板往上一套大部分问题就可以迎刃而解
		下面看一下具体的题目分析
*/

/**
题目描述1、零钱兑换
	给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
	计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
	你可以认为每种硬币的数量是无限的
	示例 1：
	输入：coins = [1, 2, 5], amount = 11
	输出：3
	解释：11 = 5 + 5 + 1

	解题思路：完全背包最值问题：外循环coins,内循环amount正序,应用状态方程1
题目链接：https://leetcode-cn.com/problems/coin-change/
*/

func coinChangeI(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt8
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := 0; i <= amount; i++ {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt8 {
		return -1
	} else {
		return dp[amount]
	}
}

/**
题目描述2：零钱兑换 II
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数
示例 1：
	输入：amount = 5, coins = [1, 2, 5]
	输出：4
	解释：有四种方式可以凑成总金额：
	5=5
	5=2+2+1
	5=2+1+1+1
	5=1+1+1+1+1

解题思路：
	完全背包

*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 0; i <= amount; i++ {
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}

/**
	题目描述：组合总和 Ⅳ
		给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。
		请你从 nums 中找出并返回总和为 target 的元素组合的个数。
		题目数据保证答案符合 32 位整数范围。

示例 1：
	输入：nums = [1,2,3], target = 4
	输出：7
	解释：
	所有可能的组合为：
	(1, 1, 1, 1)
	(1, 1, 2)
	(1, 2, 1)
	(1, 3)
	(2, 1, 1)
	(2, 2)
	(3, 1)
	请注意，顺序不同的序列被视作不同的组合。

	输入：nums = [9], target = 3
	输出：0

	解题思路：
		考虑顺序的组合问题：外循环target，内循环nums，应用状态方程3

*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
