package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := trap1(height)
	fmt.Printf("按列求解法：%v", ret)
	fmt.Println("")
	ret = trap2(height)
	fmt.Printf("动态规划求解：%v", ret)

}

/**
单调栈类似问题（ 接雨水 + 右侧 最小数 + 最大矩形）
*/

//

/*
1、接雨水
	题目描述：给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
	参考：图片连接：

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
参考图：

那具体存下了多少呢？从栈顶即将出栈的柱子来看（图中右侧第二个），雨量加上本身的高度不能超过左右柱子中较低者，故雨量为深蓝部分。
当此柱子出栈后，继续比较栈顶元素与即将入栈的柱子高度，发现仍然小，那仍然可以存下雨水，其雨量为（左右柱子中较低者-本身的高度）*2，
即浅蓝部分。为什么乘2？从图中容易看出，上一个出栈的元素并没有考虑到这一部分雨量。
直到栈顶柱子比即将入栈的柱子高或栈为空时，将此柱子入栈。总的来说，从图像可以看出，这是分层计算雨量，与之前的解法角度明显不同。实现代码如下：

*/
func trap3() {

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
