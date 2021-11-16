package main

import (
	"fmt"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	/*nums := []int{2, 3, 7, 11, 12}
	ret := twoNum(nums, 9)
	fmt.Println("返回他们的数组下标：", ret)*/

	/*nums := []int{-4, -1, -1, 0, 1, 2}

	ret := threeNum(nums)
	fmt.Printf("%v", nums)
	fmt.Println("")
	fmt.Printf("%v", ret)*/
	/*k := 5
	cutNum(nums, k)*/

	/*	nums := []int{2, 7, 3, 6}
		maxNumSwap(nums)*/

	s := "(1+(4+5+2)-13)+(6+18)"
	t := calculate(s)

	fmt.Printf("结果=%v", t)
}

/*
两数子和（I）
1、题目描述：
	 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
	你可以按任意顺序返回答案。

	示例1：
			输入：nums = [2,7,11,15], target = 9
			输出：[0,1]
			解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

创建map映射，用于存放目标数组的相关信息；
遍历目标数组，并获取目标值（target）与数组元素（nums[i]）的差值；
将差值当作map的key，目标数组的角标当作value；
判断map中是否包含，如果包含，则返回map的key为差值的value与i；
如果map中不包含，放入map中。
*/

func twoNum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

/**
21. 搜索二维矩阵I
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。每行的第一个整数大于前一行的最后一个整数。

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

*/

func searchMatrix1I(matrix [][]int, target int) bool {

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
	21 搜索二维矩阵II
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

func searchMatrixI1I(matrix [][]int, target int) bool {

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

/*
复杂度分析
时间复杂度：O(N)O(N)
空间复杂度：O(N)O(N)
*/

/**
	题目描述
给定长度为n的数组，每个元素代表一个木头的长度，木头可以任意截断，从这堆木头中截出至少k个相同长度为m的木块。已知k，求max(m)。

ps: 截断的长度必须是整数

思路： 暴力破解法
大概思路就是从1遍历到木棍最长的长度，每次遍历的长度作为m，如果可以将所有木头截出来k个长度为m的木块，则更新最大值，最后输出最大值即可

分析：上面的代码也比较容易理解，这里就不多展开说了。时间复杂度也很容易看出来是O(n * len), len为木头中最大的长度。容易想到遍历长度时可以从大到小遍历，if (cnt >= k)成立，则该值即为最终结果，可直接break，但最坏时间复杂度没变。
*/

func cutNum1(nums []int, k int) {

	//n := max(nums[0 ~ n-1])
	maxLength := 10

	res := 0
	for m := 1; m <= maxLength; m++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			cnt += nums[j] / m
		}
		if cnt >= k {
			res = max(res, m)
		}
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/**
33. 搜索旋转排序数组
	题目描述：
		整数数组 nums 按升序排列，数组中的值 互不相同 。在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
		使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
		例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
		（你可以设计一个时间复杂度为 O(log n) 的解决方案吗？）
	示例1：
		输入：nums = [4,5,6,7,0,1,2], target = 0
		输出：4
	示例2：
		输入：nums = [4,5,6,7,0,1,2], target = 3
		输出：-1

	解题思路：
		二分查找法：
		1、先判断 nums[0] <= nums[mid]  时 说明 nums【0】 到 nums[mid] 是有序递增的。则判断 nums[0] < target < nums[mid] 时 right 向right-1 查找
			否则 说明 target 在 nums[mid] ~ nums[n-1] 之间， left = mid+1
		2、如果nums[0] > nums[mid】 说明  nums[0] ~ nums[mid] 无序，旋转点 在 0 ～ nums 之间 ， 如果 target < nums[n-1]

	参考连接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/ji-jian-solution-by-lukelee/

*/
func searchXuanZhuanNums(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left := 0
	right := length - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] {
			if nums[0] < target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if left == right && nums[left] == target {
		return left
	} else {
		return -1
	}
}

func maxNumSwap(nums []int) []int {
	lastIndex := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		lastIndex[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		for d := 9; d > nums[i]; d-- {
			_, ok := lastIndex[d]
			if ok && lastIndex[d] > i {
				swap(nums, i, lastIndex[d])
			}
		}
	}

	fmt.Printf("nums %v", nums)
	return nums
}

func swap(nums []int, index int, d int) {
	temp := nums[index]
	nums[index] = nums[d]
	nums[d] = temp
}

/*
1、递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。
示例1:
	输入：A = 1, B = 10
	输出：10

解题思路：
	首先，求得A和B的最小值和最大值;
然后，可以对其中的最小值当做乘数（为什么选最小值，因为选最小值当乘数，可以算的少），将其拆分成2的幂的和
*/
func multiply(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}
	if A < B {
		return B + multiply(A-1, B)
	}
	return A + multiply(A, B-1)
}

/**
2、基本计算器
题目描述：
	给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值
	示例1：
		输入：s = "1 + 1"
		输出：2
	示例 2：
			输入：s = " 2-1 + 2 "
			输出：3
	示例 3：
		输入：s = "(1+(4+5+2)-3)+(6+8)"
		输出：23

	题目解析：
	方法一：括号展开 + 栈
	由于字符串除了数字与括号外，只有加号和减号两种运算符。因此，如果展开表达式中所有的括号，则得到的新表达式中，数字本身不会发生变化，只是每个数字前面的符号会发生变化。

	因此，我们考虑使用一个取值为 {−1,+1} 的整数
	sign 代表「当前」的符号。根据括号表达式的性质，它的取值：

	与字符串中当前位置的运算符有关；
	如果当前位置处于一系列括号之内，则也与这些括号前面的运算符有关：每当遇到一个以 -− 号开头的括号，则意味着此后的符号都要被「翻转」。
	考虑到第二点，我们需要维护一个栈 ops，其中栈顶元素记录了当前位置所处的每个括号所「共同形成」的符号。例如，对于字符串 1+2+(3-(4+5))：

	扫描到 1+2 时，由于当前位置没有被任何括号所包含，则栈顶元素为初始值 +1；
	扫描到 1+2+(3 时，当前位置被一个括号所包含，该括号前面的符号为 ++ 号，因此栈顶元素依然 +1；
	扫描到 1+2+(3-(4 时，当前位置被两个括号所包含，分别对应着 + 号和 − 号，由于 + 号和 - 号合并的结果为 -− 号，因此栈顶元素变为 -1−1。sign←ops.top()；如果遇到了遇到了 -− 号，则更新 sign←−ops.top()。

	然后，每当遇到 ( 时，都要将当前的 sign 取值压入栈中；每当遇到 ) 时，都从栈中弹出一个元素。这样，我们能够在扫描字符串的时候，即时地更新 ops 中的元素。

	https://leetcode-cn.com/problems/basic-calculator/solution/ji-ben-ji-suan-qi-by-leetcode-solution-jvir/
*/
func calculate(s string) int {
	ans := 0
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

/**
	3、旋转数组
	题目描述：
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
示例：
		输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]

	解题思路：三次反转（reverse）操作的过程如下图所示：

	原数组 ： 1，2，3，4，5，6，7
 	reverse(0,n)  7,6,5,4,3,2,1
	reverse(0,k) 5,6,7,4,3,2,1
	reverse(k,n) 5,6,7,1,2,3,4

空间复杂度 O（1）

*/
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n)
	reverse(nums, 0, k)
	reverse(nums, k, n)
}

func reverse(nums []int, begin int, end int) {
	for i, j := begin, end-1; i < j; i, j = i+1, j-1 {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
}

/**
4、快速幂
题目描述：
	实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。 （题目连接：https://leetcode-cn.com/problems/powx-n/）
示例；
输入：x = 2.00000, n = 10
输出：1024.00000

解题思路： 参考https://leetcode-cn.com/problems/powx-n/solution/powx-n-by-leetcode-solution/

*/
func myPow(x float64, n int) float64 {

	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

// 给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	if num > 0 {
		return dfs(num)
	}
	return "-" + dfs(-num)
}

func dfs(num int) string {
	if num == 0 {
		return ""
	}
	return dfs(num/2) + fmt.Sprintf("%d", num%2)
}

/**
剑指 Offer II 076. 数组中的第 k ⼤大的数字
给定整数数组 nums 和整数 k，请返回数组中第 k 个最⼤大的元素。
请注意，你需要找的是数组排序后的第 k 个最⼤大的元素，⽽而不不是第 k 个不不同的元素。
示例例 1:
输⼊入: [3,2,1,5,6,4] 和 k = 2 输出: 5
*/
func findKthLargest(nums []int, k int) int { // 1、先排序 quickSort(nums,0,len(nums)-1)
	// 2、取出第K⼤大的元素
	return nums[len(nums)-k]
}
func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	pivot := nums[(left+right)/2]
	for i <= j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	quickSort(nums, left, j)
	quickSort(nums, i, right)
}

/**
88. 合并两个有序数组
给你两个按 ⾮非递减顺序 排列列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数⽬目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 ⾮非递减顺序 排列列。
注意:最终，合并后数组不不应由函数返回，⽽而是存储在数组 nums1 中。为了了应对这种情况，nums1 的初始⻓长度为 m + n，
其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略略。nums2 的⻓长度为 n 。
示例例 1:
输⼊入:nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3 输出:[1,2,2,3,5,6]
解释:需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
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

/***
	二分查找I：
	给定⼀一个 n 个元素有序的(升序)整型数组 nums 和⼀一个⽬目标值 target ，写⼀一个函数搜索 nums 中的 target，如果⽬目标值存在返回 下标，否则返回 -1。
示例例 1:
输⼊入: nums = [-1,0,3,5,9,12], target = 9 输出: 4
解释: 9 出现在 nums 中并且下标为 4
*/
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

/**
	二分查找II
请实现有重复数字的升序数组的⼆二分查找。 输出在数组中第⼀一个⼤大于等于查找值的位置，如果数组中不不存在这样的数(指不不存在⼤大于等于查找值的数)，则输出数组⻓长度加⼀一。 示例例1
nums :{0,1,2,3,4,5,5,5,5,5,5,5,6,7,8} 输出:6 (从1开始的位置)
*/

func upper_bound_(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		num := nums[mid]
		if num < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left + 1
}

/**
判断1个数字，是否是回文数例如 1221,12321
思路：倒序遍历数字，判断是否和原数字相等。

如果是负数则一定不是回文数，直接返回 false
如果是正数，则将其倒序数值计算出来，然后比较和原数值是否相等
如果是回文数则相等返回 true，如果不是则不相等 false
比如 123 的倒序 321，不相等；121 的倒序 121，相等

*/
func isPaliadNum1(num int) bool {
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
