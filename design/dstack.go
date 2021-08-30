package main

import "fmt"

func main() {
	/*height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := trap1(height)
	fmt.Printf("按列求解法：%v", ret)
	fmt.Println("")
	ret = trap2(height)
	fmt.Printf("动态规划求解：%v", ret)

	fmt.Println("")
	ret = trap3(height)
	fmt.Printf("递减栈求解法：%v", ret)*/

	height := []int{2, 1, 5, 6, 2, 3}
	ret := largestRectangleArea(height)
	fmt.Printf("柱状图中的最大矩形面积：%v", ret)

	fmt.Println()
	ret = largestRectangleAreaStack(height)

	fmt.Printf("柱状图中最大矩形面积，单调栈解法：%v", ret)

	//nums := []int{9, 8, 7, 3, 4, 2, 1}
	nums := []int{3, 3, 1}
	rets := byteDanceStack(nums)

	fmt.Printf("rets %v", rets)

}

/**
单调栈类似问题（ 接雨水 + 右侧 最小数 + 最大矩形）
解题参考连接：https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/
https://www.jianshu.com/p/33b2fb2b00bc
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
*/

/**
解法1： 按列求解
	求每一列的水，我们只需要关注当前列，以及左边最高的墙，右边最高的墙就够了。
装水的多少，当然根据木桶效应，我们只需要看左边最高的墙和右边最高的墙中较矮的一个就够了。
所以，根据较矮的那个墙和当前列的墙的高度可以分为三种情况。
1、较矮的墙的高度大于当前列的墙的高度
2、较矮的墙的高度小于当前列的墙的高度
3、较矮的墙的高度等于当前列的墙的高度。

时间复杂度：O(n²），遍历每一列需要 nn，找出左边最高和右边最高的墙加起来刚好又是一个 n，所以是 n²。
空间复杂度：O(1）。


*/
func trap1(height []int) int {
	length := len(height)
	sum := 0
	// 最两端的列不用考虑，因为一定不会有水。所以下标从 1 到 length - 1
	for i := 1; i < length-1; i++ {
		//找出左边最高
		maxLeft := 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}

		// 找出右边最高
		maxRight := 0
		for j := i + 1; j < length; j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}

		//找出两端较小的
		min := Min(maxRight, maxLeft)

		//只有较小的一段大于当前列的高度才会有水，其他情况不会有水
		if min > height[i] {
			sum += min - height[i]

		}
	}
	return sum
}

/**
解法2 ： 动态规划
ax_left [i] 代表第 i 列左边最高的墙的高度，max_right[i] 代表第 i 列右边最高的墙的高度。（一定要注意下，第 i 列左（右）边最高的墙，是不包括自身的，和 leetcode 上边的讲的有些不同）

对于 max_left我们其实可以这样求。
max_left [i] = Max(max_left [i-1],height[i-1])。它前边的墙的左边的最高高度和它前边的墙的高度选一个较大的，就是当前列左边最高的墙了。
对于 max_right我们可以这样求。
max_right[i] = Max(max_right[i+1],height[i+1]) 。它后边的墙的右边的最高高度和它后边的墙的高度选一个较大的，就是当前列右边最高的墙了。

这样，我们再利用解法二的算法，就不用在 for 循环里每次重新遍历一次求 max_left 和 max_right 了。

*/
func trap2(height []int) int {
	sum := 0
	length := len(height)
	arr_max_left := make([]int, length)
	arr_max_right := make([]int, length)
	for i := 1; i < length-1; i++ {
		arr_max_left[i] = Max(arr_max_left[i-1], height[i-1])
	}

	for i := length - 1 - 1; i >= 0; i-- {
		arr_max_right[i] = Max(arr_max_right[i+1], height[i+1])
	}

	for i := 1; i < length-1; i++ {
		min := Min(arr_max_left[i], arr_max_right[i])
		if min > height[i] {
			sum += min - height[i]
		}
	}
	return sum
}

/**
方法3：
	单调递减栈求解：

	该方法相对难以理解：不过提供了一个看问题的角度，前两种方法是逐列统计雨量的，而单调栈法则是分块逐层统计雨量
考虑一个单调递减的栈，即栈顶元素最小。因为当后入栈的柱子高度较小时，是不可能存下雨水的（见下图）。只有当即将入栈的柱子高度比栈顶的大时，才开始存下雨水。
参考图：https://github.com/zlbonly/leetcode-practice/blob/master/pic/raiin_stack.png

那具体存下了多少呢？从栈顶即将出栈的柱子来看（图中右侧第二个），雨量加上本身的高度不能超过左右柱子中较低者，故雨量为深蓝部分。
当此柱子出栈后，继续比较栈顶元素与即将入栈的柱子高度，发现仍然小，那仍然可以存下雨水，其雨量为（左右柱子中较低者-本身的高度）*2，
即浅蓝部分。为什么乘2？从图中容易看出，上一个出栈的元素并没有考虑到这一部分雨量。
直到栈顶柱子比即将入栈的柱子高或栈为空时，将此柱子入栈。总的来说，从图像可以看出，这是分层计算雨量，与之前的解法角度明显不同。实现代码如下：

*/
func trap3(height []int) int {
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

/**
题型二、下一个更大元素1
1、题目描述：
	给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
	请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
	nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

示例1：
输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
    对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
    对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。


参考连接：https://leetcode-cn.com/problems/next-greater-element-i/
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) != 0 && v > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, v)
	}

	for k, v := range nums1 {
		if value, ok := m[v]; ok {
			nums1[k] = value
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}

/**
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
数字 x 的下一个更大的元素是按数组遍历顺序，
这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。


示例1：
	输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。


复杂度分析
时间复杂度: O(n)O(n)，其中 nn 是序列的长度。我们需要遍历该数组中每个元素最多 22 次，每个元素出栈与入栈的总次数也不超过 44 次。
空间复杂度: O(n)O(n)，其中 nn 是序列的长度。空间复杂度主要取决于栈的大小，栈的大小至多为 2n-12n−1。
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}

/**
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
柱状图中的最大矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

参考图：https://github.com/zlbonly/leetcode-practice/blob/master/pic/Rec.png
示例1：
	输入: [2,1,5,6,2,3]
	输出: 10
*/

func largestRectangleArea(height []int) int {
	length := len(height)
	res := 0
	for i := 0; i < length; i++ {
		left := i
		curHeight := height[i]
		for left > 0 && height[left-1] >= curHeight {
			left--
		}
		right := i
		for right < length-1 && height[right+1] >= curHeight {
			right++
		}
		width := right - left + 1
		res = Max(res, width*curHeight)
	}
	return res
}

/**
1、该题目用栈求解的思路是
1.先将题目给定的数组左右各添加一个元素0，为了方便确定原有数组中第一个元素和最后一个元素能不能继续扩张；
2.然后开始从左到右依次遍历数组中的元素：
3-1.如果栈为空或者当前考察的新元素值比栈顶元素值大，表明以栈顶元素值为高的矩形面积暂不能确定，
所以就将当前考察的新元素入栈。在这个条件下，栈中的元素值从栈底到栈顶是依次递增的；
3-2.如果栈不为空且当前考察的新元素值比栈顶元素值小，表明以栈顶元素值为高的矩形的面积是可以确定的了。该矩形的高就是栈顶元素值，其右侧边界就是当前考察的新元素，
左侧边界是栈顶元素的前一个元素，因为，在上一步中我们知道栈中元素值从栈底到栈顶是依次递增的。 因此，矩形的宽是当前考察的元素索引与栈顶元素前一个元素的索引的差值减一。

这里需要注意的是，当栈顶元素出栈后，需要继续看当前考察的新元素值是否大于新的栈顶元素值，如果是，就继续将栈顶元素弹出，
然后计算以其值为高的矩形面积，直到当前考察的新元素值大于栈顶元素值时，当前考察元素入栈。
最后，由于最终计算矩形面积时，是用两个柱子的索引来确定矩形宽度的。因此，栈中存储的应该是给定数组的索引。

参考连接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/dong-hua-yan-shi-dan-diao-zhan-84zhu-zhu-03w3/


*/
func largestRectangleAreaStack(height []int) int {

	// 初始化最终结果0
	res := 0
	length := len(height)

	// 将给定的原数组左右各添加一个元素0
	newHeight := make([]int, length+2)
	newHeight[0] = 0
	newHeight[len(newHeight)-1] = 0

	for i := 1; i < length+1; i++ {
		newHeight[i] = height[i-1]
	}

	var stack []int
	for i := 0; i < len(newHeight); i++ {
		// 如果栈不为空且当前考察的元素值小于栈顶元素值，
		// 则表示以栈顶元素值为高的矩形面积可以确定
		for len(stack) > 0 && newHeight[i] < newHeight[stack[len(stack)-1]] {
			// 获取栈顶元素对应的高
			curHeight := newHeight[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			// 栈顶元素弹出后，新的栈顶元素就是其左侧边界
			leftIndex := stack[len(stack)-1]
			right := i                         // 右侧边界是当前考察的索引
			curWidth := right - leftIndex - 1  // 计算矩形宽度
			res = Max(res, curWidth*curHeight) // 计算面积

		}
		// 当前考察索引入栈
		stack = append(stack, i)
	}

	return res
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
func byteDanceStack(nums []int) []int {
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
