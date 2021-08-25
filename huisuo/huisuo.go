package huisuo

import "sort"

/**
回溯算法：
	回溯法 采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，
	它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。
	回溯法通常用最简单的递归方法来实现，在反复重复上述的步骤后可能出现两种情况：
  	 找到一个可能存在的正确的答案；在尝试了所有可能的分步方法后宣告该问题没有答案。

回溯和动态规划的区别：
共同点
		1、用于求解多阶段决策问题。多阶段决策问题即：

		2、求解一个问题分为很多步骤（阶段）；每一个步骤（阶段）可以有多种选择。
不同点
		1、动态规划只需要求我们评估最优解是多少，最优解对应的具体解是什么并不要求。因此很适合应用于评估一个方案的效果；
		2、回溯算法可以搜索得到所有的方案（当然包括最优解），但是本质上它是一种遍历算法，时间复杂度很高。

列题： 1和2 排列
      3 子集I ，4 子集II
      5 组合总和I，组合总和II

	复杂度 n * 2的n次方

*/

/**
1、题目描述： 全排列I
	给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
解题思路：回溯法
	从全排列问题开始理解回溯算法
我们尝试在纸上写 3 个数字、4 个数字、5 个数字的全排列，相信不难找到这样的方法。以数组 [1, 2, 3] 的全排列为例。

1、先写以 1 开头的全排列，它们是：[1, 2, 3], [1, 3, 2]，即 1 + [2, 3] 的全排列（注意：递归结构体现在这里）；
2、再写以 2 开头的全排列，它们是：[2, 1, 3], [2, 3, 1]，即 2 + [1, 3] 的全排列；
3、最后写以 3 开头的全排列，它们是：[3, 1, 2], [3, 2, 1]，即 3 + [1, 2] 的全排列。

总结搜索的方法：按顺序枚举每一位可能出现的情况，已经选择的数字在 当前 要选择的数字中不能出现。按照这种策略搜索就能够做到 不重不漏

参考链接：https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
*/

func permute(nums []int) [][]int {

	res := make([][]int, 0)
	path := make([]int, 0)                     // 初始化中间变量，用于存放临时结果
	isVistied := make([]bool, len(nums))       // 判断该位置数字是否用过
	dfsPermute(isVistied, path, nums, 0, &res) // 回溯函数
	return res
}

func dfsPermute(isVistied []bool, path []int, nums []int, depth int, res *[][]int) {
	// 回溯函数结束条件
	if len(nums) == depth {
		current := make([]int, len(path))
		copy(current, path)
		*res = append(*res, current)
		return
	}
	// 遍历数组中的数字，进行排列组合
	for i := 0; i < len(nums); i++ {
		// 减枝，当该位置数字已经使用过时则跳过
		if isVistied[i] {
			continue
		}
		// 添加数字
		path = append(path, nums[i])
		// 该位置数字设置为已访问过
		isVistied[i] = true
		// 继续搜索该支线
		dfsPermute(isVistied, path, nums, depth+1, res)
		// 回溯 恢复到之前的状态
		path = path[:len(path)-1]
		isVistied[i] = false
	}
}

/**
	题目描述：全排列II，
	 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
	输入：nums = [1,1,2]
	输出：
	[[1,1,2],
	 [1,2,1],
	 [2,1,1]]
考虑重复元素一定要优先排序，将重复的都放在一起，便于找到重复元素和剪枝！！！
推广至 --> 如果涉及考虑重复元素，或者大小比较的情况，对列表排序是一个不错的选择

	参考链接：https://leetcode-cn.com/problems/permutations-ii/solution/hot-100-47quan-pai-lie-ii-python3-hui-su-kao-lu-zh/
*/

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	// 1、先排序
	sort.Sort(sort.IntSlice(nums))
	path := []int{}
	isVistied := make([]bool, len(nums))
	dfsPermuteUnique(isVistied, path, nums, 0, &res)
	return res
}

func dfsPermuteUnique(isVistied []bool, path []int, nums []int, depth int, res *[][]int) {
	if len(nums) == depth {
		current := make([]int, len(path))
		copy(current, path)
		*res = append(*res, current)
		return
	}
	for i := 0; i < len(nums); i++ {
		if isVistied[i] {
			continue
		}

		// 2、为了去除重复，增加新的减枝条件
		if i > 0 && nums[i] == nums[i-1] && isVistied[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		isVistied[i] = true
		dfsPermuteUnique(isVistied, path, nums, depth+1, res)
		path = path[:len(path)-1]
		isVistied[i] = false
	}
}

/**

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

解题连接：https://mp.weixin.qq.com/s/g5uvxi1lyxmWC4LtP0Bdlw
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	current := []int{}

	dfsSubsets(nums, 0, current, &res)
	return res
}

func dfsSubsets(nums []int, k int, current []int, res *[][]int) {
	if k == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, current)
		return
	}
	// 不选择第k个元素
	dfsSubsets(nums, k+1, current, res)
	// 选择第k个元素
	current = append(current, nums[k])
	dfsSubsets(nums, k+1, current, res)
	current = current[:len(current)-1]
}

/**
5、题目描述
给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
对于给定的输入，保证和为 target 的唯一组合数少于 150 个。

输入: candidates = [2,3,6,7], target = 7
输出: [[7],[2,2,3]]
*/

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	dfsCombinationSum(candidates, 0, target, path, &res)
	return res
}

func dfsCombinationSum(candidates []int, depth int, target int, path []int, res *[][]int) {
	// target 为负数和 0 的时候不再产生新的孩子结点
	if target < 0 {
		return
	}
	// 重点理解这里从 begin 开始搜索的语意
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := depth; i < len(candidates); i++ {
		path = append(path, candidates[i])
		//  注意：由于每一个元素可以重复使用，下一轮搜索的起点依然是 i，这里非常容易弄错

		dfsCombinationSum(candidates, i, target-candidates[i], path, res)
		path = path[:len(path)-1]
	}
}

/**
6、题目描述：组合总和 （元素不能重复使用）

给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
注意：解集不能包含重复的组合。

 参考链接：https://leetcode-cn.com/problems/combination-sum-ii/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-3/

*/

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Sort(sort.IntSlice(candidates))
	dfsCombinationSum2(candidates, 0, target, path, &res)
	return res
}

func dfsCombinationSum2(candidates []int, depth int, target int, path []int, res *[][]int) {
	// target 为负数和 0 的时候不再产生新的孩子结点
	if target < 0 {
		return
	}
	// 重点理解这里从 begin 开始搜索的语意
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := depth; i < len(candidates); i++ {
		// 大剪枝：减去 candidates[i] 小于 0，减去后面的 candidates[i + 1]、candidates[i + 2] 肯定也小于 0，因此用 break
		if target-candidates[i] < 0 {
			break
		}
		// 小剪枝：同一层相同数值的结点，从第 2 个开始，候选数更少，结果一定发生重复，因此跳过，用 continue
		if i > depth && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])

		//  注意：由于每一个元素可以重复使用，下一轮搜索的起点依然是 i，这里非常容易弄错
		dfsCombinationSum(candidates, i+1, target-candidates[i], path, res)
		path = path[:len(path)-1]
	}
}
