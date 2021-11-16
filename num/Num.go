package main

import (
	"math"
	"sort"
)

/**
1、整数反转
2、搜索二维矩阵I
3、搜索二维矩阵II
4、截取木头
5、是否回文数字
6、合并两个有序数组
*/
/*
	题目1：整数反转
	整数反转 12345 => 54321  ,-321 => 123。
	将12345 % 10 得到5，之后将12345 / 10
	2、将1234 % 10 得到4，再将1234 / 10
	3、将123 % 10 得到3，再将123 / 10
	4、将12 % 10 得到2，再将12 / 10
	5、将1 % 10 得到1，再将1 / 10
链接：https://leetcode-cn.com/problems/reverse-integer/solution/tu-jie-7-zheng-shu-fan-zhuan-by-wang_ni_ma/
*/
func reverseInt32(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10 // 求尾部数
		// maxInt32 最大值 2147483647
		if x > math.MaxInt32 || (x == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		// maxMinInt32  最小值 -2147483648
		if x < math.MinInt32 || (x == math.MinInt32/10 && pop < -8) {
			return 0
		}
		res = res*10 + pop
		x = x / 10
	}
	return res
}

/*
	2. 搜索二维矩阵I
	编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
	每行中的整数从左到右按升序排列。每行的第一个整数大于前一行的最后一个整数。

	输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
	输出：true
*/

func searchMatrixI(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && target <= matrix[i][len(matrix[i])-1] {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}

/**
3 搜索二维矩阵II
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
	!! 每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。

	示例：
		输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
		输出：true

解法说明： 上searchMatrixI 也可用  时间复杂度 m*n

优化解法：searchMatrixII
因为矩阵的行和列是排序的（分别从左到右和从上到下），所以在查看任何特定值时，我们可以修剪O(m)O(m)或O(n)O(n)元素。

算法：
首先，我们初始化一个指向矩阵左下角的 (row，col)(row，col) 指针。然后，直到找到目标并返回 true（或者指针指向矩阵维度之外的 (row，col)(row，col) 为止，
我们执行以下操作：如果当前指向的值大于目标值，则可以 “向上” 移动一行。 否则，如果当前指向的值小于目标值，则可以移动一列。
不难理解为什么这样做永远不会删减正确的答案；因为行是从左到右排序的，所以我们知道当前值右侧的每个值都较大。 因此，如果当前值已经大于目标值，
我们知道它右边的每个值会比较大。也可以对列进行非常类似的论证，因此这种搜索方式将始终在矩阵中找到目标（如果存在）。

时间复杂度：O(n+m)。
时间复杂度分析的关键是注意到在每次迭代（我们不返回 true）时，行或列都会精确地递减/递增一次。由于行只能减少 mm 次，而列只能增加 nn 次，因此在导致 while 循环终止之前，循环不能运行超过 n+mn+m 次。因为所有其他的工作都是常数，所以总的时间复杂度在矩阵维数之和中是线性的。
空间复杂度：O(1)，因为这种方法只处理几个指针，所以它的内存占用是恒定的。
*/

func searchMatrixII(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}

/**
4、题目描述 截取木头
	给定长度为n的数组，每个元素代表一个木头的长度，木头可以任意截断，从这堆木头中截出至少k个相同长度为m的木块。已知k，求max(m)。
	ps: 截断的长度必须是整数

	思路： 暴力破解法
	1、对nums 进行排序，获取nums 中最大值 maxLength
	2、从 1 ~ maxLength 遍历 ，内部遍历 nums ，记录 每个长度 nums 中元素可以截取的最大数量，如果>= k 则记录
	3、找出 1~ maxLenght 满足条件的最大值

	分析：上面的代码也比较容易理解，这里就不多展开说了。时间复杂度也很容易看出来是O(n * len),
		len为木头中最大的长度。容易想到遍历长度时可以从大到小遍历，if (cnt >= k)成立，
		则该值即为最终结果，可直接break，但最坏时间复杂度没变。
*/
func cutNum(nums []int, k int) int {
	sort.Ints(nums)
	maxLength := nums[len(nums)-1]
	res := 0
	for i := 1; i <= maxLength; i++ {
		cur := 0
		for j := 0; j < len(nums); j++ {
			cur += nums[j] / i
		}
		if cur >= k {
			res = max(res, i)
		}
	}
	return res
}

/**
5、题目描述：是否回文数字
	判断1个数字，是否是回文数例如 1221,12321
思路：倒序遍历数字，判断是否和原数字相等。

如果是负数则一定不是回文数，直接返回 false
如果是正数，则将其倒序数值计算出来，然后比较和原数值是否相等
如果是回文数则相等返回 true，如果不是则不相等 false
比如 123 的倒序 321，不相等；121 的倒序 121，相等

*/
func isHwNum(num int) bool {
	if num < 0 {
		return false
	}
	total := 0
	cur := num
	for cur != 0 {
		total = total*10 + cur%10
		cur = cur / 10
	}
	return total == num
}

/**
6. 合并两个有序数组
	给你两个按 ⾮非递减顺序 排列列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数⽬目。
	请你 合并 nums2 到 nums1 中，使合并后的数组同样按 ⾮非递减顺序 排列列。
	注意:最终，合并后数组不不应由函数返回，⽽而是存储在数组 nums1 中。为了了应对这种情况，nums1 的初始⻓长度为 m + n，
	其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略略。nums2 的⻓长度为 n 。
	示例例 1:
	输⼊入:nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3 输出:[1,2,2,3,5,6]
	解释:需要合并 [1,2,3] 和 [2,5,6] 。
	合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/
func mergeTwoNum(nums1 []int, m int, nums2 []int, n int) {
	right1, right2, tail := m-1, n-1, m+n-1
	var cur int
	for right1 >= 0 || right2 >= 0 {
		if right1 == -1 {
			cur = nums2[right2]
			right2--
		} else if right2 == -1 {
			cur = nums1[right1]
			right1--
		} else if nums2[right2] > nums1[right1] {
			cur = nums2[right2]
			right2--
		} else {
			cur = nums1[right1]
			right1--
		}
		nums1[tail] = cur
		tail--
	}
}
