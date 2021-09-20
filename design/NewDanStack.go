package main

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

// 暴力破解
func largestRectangleArea2(height []int) int {
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
func yushui1(height []int) int {
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
