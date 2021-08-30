package main

/*
什么是单调栈
单调栈就是栈里面存放的数据都是有序的，所以可以分为单调递增栈和单调递减栈两种。
1、单调递增栈就是从栈底到栈顶是从大到小
2、单调递减栈就是从栈底到栈顶是从小到大
栈基础知识：
1、栈是一种特殊的线性表，仅允许在表的一端进行插入和删除。 这一端被称为栈顶。 （先进后出）
2、单调栈

    所谓单调栈 则是在栈的 先进后出 的基础之上额外新增一个特性： 从栈顶到栈底的元素是严格递增（递减）

    具体进栈过程如下：

    1、对于单调递增栈，若当前进栈元素为 e，从栈顶开始遍历元素，把小于 e 或者等于 e 的元素弹出栈，直接遇到一个大于 e 的元素或者栈为空为止，然后再把 e 压入栈中。

    2、对于单调递减
*/

/**
单调栈典型例题：
1、
*/

/*
1、接雨水
	题目描述：给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
	参考：图片连接：
	https://github.com/zlbonly/leetcode-practice/blob/master/pic/stack_pic.png

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

单调递减栈求解：
该方法相对难以理解：不过提供了一个看问题的角度，前两种方法是逐列统计雨量的，而单调栈法则是分块逐层统计雨量
考虑一个单调递减的栈，即栈顶元素最小。因为当后入栈的柱子高度较小时，是不可能存下雨水的（见下图）。只有当即将入栈的柱子高度比栈顶的大时，才开始存下雨水。
参考图：https://github.com/zlbonly/leetcode-practice/blob/master/pic/raiin_stack.png

那具体存下了多少呢？从栈顶即将出栈的柱子来看（图中右侧第二个），雨量加上本身的高度不能超过左右柱子中较低者，故雨量为深蓝部分。
当此柱子出栈后，继续比较栈顶元素与即将入栈的柱子高度，发现仍然小，那仍然可以存下雨水，其雨量为（左右柱子中较低者-本身的高度）*2，
即浅蓝部分。为什么乘2？从图中容易看出，上一个出栈的元素并没有考虑到这一部分雨量。
直到栈顶柱子比即将入栈的柱子高或栈为空时，将此柱子入栈。总的来说，从图像可以看出，这是分层计算雨量，与之前的解法角度明显不同。
实现代码如下：

我们用栈保存每堵墙。
当遍历墙的高度的时候，如果当前高度小于栈顶的墙高度，说明这里会有积水，我们将墙的高度的下标入栈。
如果当前高度大于栈顶的墙的高度，说明之前的积水到这里停下，我们可以计算下有多少积水了。计算完，就把当前的墙继续入栈，作为新的积水的墙。
总体的原则就是，
当前高度小于等于栈顶高度，入栈，指针后移。
当前高度大于栈顶高度，出栈，计算出当前墙和栈顶的墙之间水的多少，然后计算当前的高度和新栈的高度的关系，重复第 2 步。直到当前墙的高度不大于栈顶高度或者栈空，然后把当前墙入栈，指针后移。
。
*/
func trapRain(height []int) int {
	length := len(height)
	stack := make([]int, 0)
	sum := 0
	current := 0
	for current < length {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			h := height[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			min := Min(height[stack[len(stack)-1]], height[current])
			sum += distance * (min - h)
		}

		stack = append(stack, current)
		current++
	}
	return sum
}

/*
	前几天被问到一道字节的面试题：
找到数组中, 比左边所有数字都小, 比右边所有数字都大的 数字

先从左到右遍历，求一个leftMin数组，记录nums[i]左边的最小值；
再从右到左遍历求一个rightMax数组，记录nums[i]右边的最大值。
然后从左到右遍历每个元素，如果某元素满足leftMin[i]>nums[i]>rightMax[i]，则求得该数

示例1：
	input: 9,8,7,3,4,2,1
	output: 9,8,7,2,1

	input: 3,3,1
	output: 1

*/
func byteDanceStack2(nums []int) []int {
	length := len(nums)
	var stack []int
	var ret []int
	for i := 0; i < length; i++ {

		if len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			for len(stack) > 0 && nums[i] <= stack[len(stack)-1] {
				ret = append(ret, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, nums[i])
			if i == length-1 && len(stack) > 0 && nums[i] <= stack[len(stack)-1] {
				ret = append(ret, stack[len(stack)-1])
			}
		}
	}
	return ret
}

func Min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func Max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}
