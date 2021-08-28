package main

import "fmt"

func main() {
	/*1、爬楼梯
	climbStairs(6)*/

	/* 2 最大子序列和
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	a := maxSubArray(arr)
	fmt.Println(a)*/

	// 3 买卖股票的最佳时机(1)
	/*arr := []int{7, 1, 5, 3, 6, 4}
	a := maxProfit(arr)*/

	// 4、 买卖股票最佳时机（2）
	/*arr := []int{1, 2, 3, 4, 5}
	b := maxProfitTwo(arr)
	fmt.Println(b)*/

	/*
		// 5、不同路径（1）
		path := uniquePaths(7,3)
		fmt.Println(path)
	*/

	/*
		6、不同路径 （2）
		path := [][]int {{0,0,0},{0,1,0},{0,0,0}}
		path_num := uniquePathsWithObstacles(path)
		fmt.Println(path_num)
	*/

}

/**
 1题目描述：假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
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
func maxProfit1(prices []int) int {
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

/**
	4、买卖股票（2）
	题目描述：
		给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
		设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
		注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

 	输入: [7,1,5,3,6,4]
	输出: 7
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3

	解题方案：
	这个题用暴力法就有点太复杂了，得稍微转换一下思路，最困扰我们的地方就是当1，2，3，4，5这类型的情况，我们不知道是需要卖出再买还是等最高价卖出，
	但是换个思路想一下，5-1等于（2-1）+（3-2）+（4-3）+（5-4），也就是说这种情况只要后面数字大于前面数字，那就叠加利润，
、	最后一定是最大利润，如果中间有一个数是小的，那么直接跳过即可，相当于分成了几个这样的数列。
*/
func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

/**
	4、题目描述：买卖股票（3）
	给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
	设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	题目链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

	解题思路：
		方法一：动态规划
		思路与算法
		由于我们最多可以完成两笔交易，因此在任意一天结束之后，我们会处于以下五个状态中的一种：
		a、未进行过任何操作
		b、只进行过一次买操作
		c、进行了一次买操作和一次卖操作，即完成了一笔交易
		d、在完成了一笔交易的前提下，进行了第二次买操作
		e、完成了全部两笔交易

	由于第一个状态的利润显然为 0，因此我们可以不用将其记录。对于剩下的四个状态，
	我们分别将它们的最大利润记为:buy1,sell1,buy2,sell2

	1、对于 buy1而言，在第 i 天我们可以不进行任何操作，保持不变，也可以在未进行任何操作的前提下以prices[i] 的价格买入股票，那么buy1
  的状态转移方程即为： buy1 = max(buy1[i-1] -prices[i])
	2、对于sell1而言 在第 i 天我们可以不进行任何操作，保持不变，也可以在只进行过一次买操作的前提下以 prices[i] 的价格卖出股票，
	那么sell1的状态转移方程为：
		sell1 = max(sell1[i-1],buy1[i-1]+price[i-1=])

	同理得 buy2 和 sell2
			buy2 = max(buy2[i-1],sell1[i-1]-price[i])
			sell2 = max(sell2[i-1]+buy2[i-1]+price[i])

		解题链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iii-by-wrnt/
*/
func maxProfit3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

/**
1、题目描述：买卖股票的最佳时机含手续费、
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。你可以无限次地完成交易，但是你每笔交易都需要付手续费。
如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。返回获得利润的最大值。
注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

题目链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

解题思路：方法一：动态规划
考虑到「不能同时参与多笔交易」，因此每天交易结束后只可能存在手里有一支股票或者没有股票的状态。
定义状态dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润，dp[i][1] 表示第 i天交易完后手里持有一支股票的最大利润（i 从  开始）。
1、
	考虑 dp[i][0] 的转移方程，如果这一天交易完后手里没有股票，那么可能的转移状态为前一天已经没有股票，即dp[i−1][0]，或者前一天结束的时候手里持有一支股票，即 dp[i−1][1]，这时候我们要将其卖出，并获得 prices[i] 的收益，但需要支付fee 的手续费。
	因此为了收益最大化，我们列出如下的转移方程：
	dp[i][0]=max{dp[i−1][0],dp[i−1][1]+prices[i]−fee}
2、再来按照同样的方式考虑 dp[i][1] 按状态转移，那么可能的转移状态为前一天已经持有一支股票
	即 dp[i−1][1]，或者前一天结束时还没有股票，即 dp[i−1][0]，这时候我们要将其买入，并减少prices[i] 的收益。可以列出如下的转移方程：
	dp[i][1]=max{dp[i−1][1],dp[i−1][0]−prices[i]}
	对于初始状态，根据状态定义我们可以知道第 00 天交易结束的时候有 dp[0][0]=0 以及 dp[0][1]=−prices[0]。

解题链接参考：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-han-sh-rzlz/
*/
func maxProfit5(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, 0-prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return max(dp[n-1][0], dp[n-1][1])
}

/**
	题目描述： 最佳买卖股票时机含冷冻期
	给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
	设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
	你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

	输入: [1,2,3,0,2]
	输出: 3
	解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

	题目链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
	解题思路：
		方法一：动态规划
		1、我们用 f[i] 表示第 i 天结束之后的「累计最大收益」。根据题目描述，由于我们最多只能同时买入（持有）一支股票，并且卖出股票后有冷冻期的限制，因此我们会有三种不同的状态：
			a、我们目前持有一支股票，对应的「累计最大收益」记为 f[i][0]；
			b、我们目前不持有任何股票，并且处于冷冻期中，对应的「累计最大收益」记为 f[i][1]；
			c、我们目前不持有任何股票，并且不处于冷冻期中，对应的「累计最大收益」记为 f[i][2]。
		2、对于 f[i][0]，我们目前持有的这一支股票可以是在第 i-1天就已经持有的，对应的状态为 f[i−1][0]；或者是第 ii 天买入的，那么第 i-1 天就不能持有股票并且不处于冷冻期中，
           对应的状态为 f[i−1][2] 加上买入股票的负收益 prices[i]。
				因此状态转移方程为：
				f[i][0]=max(f[i−1][0],f[i−1][2]−prices[i])
		3、对于 f[i][1]，我们在第 i 天结束之后处于冷冻期的原因是在当天卖出了股票，那么说明在第 i−1 天时我们必须持有一支股票，
			对应的状态为 f[i−1][0] 加上卖出股票的正收益 prices[i]。因此状态转移方程为：
			f[i][1]=f[i−1][0]+prices[i]
		4、对于 f[i][2]，我们在第 i 天结束之后不持有任何股票并且不处于冷冻期，说明当天没有进行任何操作，即第 i−1 天时不持有任何股票：如果处于冷冻期，对应的状态为 f[i−1][1]；
			如果不处于冷冻期，对应的状态为 f[i−1][2]。因此状态转移方程为：
			f[i][2]=max(f[i−1][1],f[i−1][2])

		这样我们就得到了所有的状态转移方程。如果一共有 nn 天，那么最终的答案即为
		max(f[n−1][0],f[n−1][1],f[n−1][2])

	解题链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/zui-jia-mai-mai-gu-piao-shi-ji-han-leng-dong-qi-4/
*/
func maxProfit6(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

/**

 */

/**
题目描述： 不同路径 （1）
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右

解题思路：
	我们令 dp[i][j] 是到达 i, j 最多路径
动态方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
注意，对于第一行 dp[0][j]，或者第一列 dp[i][0]，由于都是在边界，所以只能为 1
算法分析：时间复杂度：O(m*n)，空间复杂度：O(m * n)

*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

/*
 题目描述：  不同路径（II）
位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。
说明：m 和 n 的值均不超过 100。
示例 1:
输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右

题目分析：
	算法

如果第一个格点 obstacleGrid[0,0] 是 1，说明有障碍物，那么机器人不能做任何移动，我们返回结果 0。
否则，如果 obstacleGrid[0,0] 是 0，我们初始化这个值为 1 然后继续算法。
遍历第一行，如果有一个格点初始值为 1 ，说明当前节点有障碍物，没有路径可以通过，设值为 0 ；否则设这个值是前一个节点的值 obstacleGrid[i,j] = obstacleGrid[i,j-1]。
遍历第一列，如果有一个格点初始值为 1 ，说明当前节点有障碍物，没有路径可以通过，设值为 0 ；否则设这个值是前一个节点的值 obstacleGrid[i,j] = obstacleGrid[i-1,j]。
现在，从 obstacleGrid[1,1] 开始遍历整个数组，如果某个格点初始不包含任何障碍物，就把值赋为上方和左侧两个格点方案数之和 obstacleGrid[i,j] = obstacleGrid[i-1,j] + obstacleGrid[i,j-1]。
如果这个点有障碍物，设值为 0 ，这可以保证不会对后面的路径产生贡献。

*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	matrix := obstacleGrid
	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[0][0] == 1 {
		return 0
	}
	// 处理第一行
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			dp[j] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ { // 从第二行开始
		if matrix[i][0] == 1 { // 处理每一个的第一列
			dp[0] = 0
		}
		for j := 1; j < n; j++ { // 从第二列开始
			if matrix[i][j] == 1 { // 遇到1则到达第j列置为0
				dp[j] = 0
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划类习题

/**
	1、打家劫舍I
		你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
		给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
示例 1：
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

动态转移方程：
dp[k] = Max (dp[k-1],dp[k-2]+1)

空间复杂度O(n)
*/
func robI(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	N := len(nums)
	dp := make([]int, N+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= N; i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[N]
}

/**
打家劫舍I 优化  空间复杂度O（1）

*/
func rob(nums []int) int {
	first, second := nums[0], Max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, Max(first+v, second)
	}
	return second
}

/*
	2、打家劫舍II
	题目描述：
	你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

解题思路： 因为是一个环，因此


总体思路：
此题是 198. 打家劫舍 的拓展版： 唯一的区别是此题中的房间是环状排列的（即首尾相接），而 198.198. 题中的房间是单排排列的；而这也是此题的难点。
环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃，因此可以把此环状排列房间问题约化为两个单排排列房间子问题：

1、在不偷窃第一个房子的情况下（即 nums[1:]），最大金额是 p1
2、在不偷窃最后一个房子的情况下（即 nums[:n-1]），最大金额是 p2
综合偷窃最大金额： 为以上两种情况的较大值，即 max(p1,p2) 。
下面的任务则是解决 单排排列房间（即 198. 打家劫舍） 问题。推荐可以先把 198.198. 做完再做这道题。

*/
func robII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob(nums[:n-1]), rob(nums[1:]))
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
