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

func cutNum(nums []int, k int) {

	//n := max(nums[0 ~ n-1])
	maxLength := 10

	res := 0
	for m := 1; m <= maxLength; m++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			cnt += nums[j] / m
		}

		fmt.Println(cnt, m)

		if cnt >= k {
			res = max(res, m)
		}
	}

	fmt.Println(res)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/**
	排序 + 双指针
本题的难点在于如何去除重复解。
算法流程：
特判，对于数组长度 nn，如果数组为 nullnull 或者数组长度小于 33，返回 [][]。
对数组进行排序。
遍历排序后数组：
若 nums[i]>0nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 00，直接返回结果。
对于重复元素：跳过，避免出现重复解
令左指针 L=i+1L=i+1，右指针 R=n-1R=n−1，当 L<RL<R 时，执行循环：
当 nums[i]+nums[L]+nums[R]==0nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,RL,R 移到下一位置，寻找新的解
若和大于 00，说明 nums[R]nums[R] 太大，RR 左移
若和小于 00，说明 nums[L]nums[L] 太小，LL 右移
复杂度分析
时间复杂度：O\left(n^{2}\right)O(n
2
 )，数组排序 O(N \log N)O(NlogN)，遍历数组 O\left(n\right)O(n)，双指针遍历 O\left(n\right)O(n)，总体 O(N \log N)+O\left(n\right)*O\left(n\right)O(NlogN)+O(n)∗O(n)，O\left(n^{2}\right)O(n
2
 )
空间复杂度：O(1)O(1)

作者：wu_yan_zu
链接：https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
func threeNum(nums []int) [][]int {
	len := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < len; i++ {
		L := i + 1
		R := len - 1
		if nums[i] > 0 && nums[i] == nums[i-1] {
			continue
		}
		for L < R {
			temp := nums[i] + nums[L] + nums[R]
			if temp == 0 {
				arr := []int{nums[i], nums[L], nums[R]}
				ret = append(ret, arr)
				for L < R && nums[L] == nums[i+1] {
					L += 1
				}
				for L < R && nums[R] == nums[R-1] {
					R -= 1
				}
				L = L + 1
				R = R - 1
			} else if temp < 0 {
				L = L + 1
			} else {
				R = R - 1
			}

		}
	}
	return ret
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
