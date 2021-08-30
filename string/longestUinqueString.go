package main

import (
	"math"
	"sort"
	"strings"
)

func main() {
	// 1、 无重复字符的最长子串
	// 2、数组中最长公共前缀
	// 3、 最小覆盖子串
	// 4、最长公共子序列（两个字符串）
	// 5、题目描述：两数子和（I）
	// 6、三数之和
}

/*
1、题目：无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串

解题思路
窗口可以在两个边界移动一开始窗口大小为0
随着数组下标的前进窗口的右侧依次增大
每次查询窗口里的字符，若窗口中有查询的字符
窗口的左侧移动到该字符加一的位置
每次记录窗口的最大程度
重复操作直到数组遍历完成
返回最大窗口长度即可

流程：
	我们不妨以示例一中的字符串 abcabcbb 为例，找出从每一个字符开始的，不包含重复字符的最长子串，那么其中最长的那个字符串即为答案。对于示例一中的字符串，我们列举出这些结果，其中括号中表示选中的字符以及最长的字符串：

以 (a)bcabcbb 开始的最长字符串为 (abc)abcbb；
以 a(b)cabcbb 开始的最长字符串为 a(bca)bcbb；
以 ab(c)abcbb 开始的最长字符串为 ab(cab)cbb；
以 abc(a)bcbb 开始的最长字符串为 abc(abc)bb；
以 abca(b)cbb 开始的最长字符串为 abca(bc)bb；
以 abcab(c)bb 开始的最长字符串为 abcab(cb)b；
以 abcabc(b)b 开始的最长字符串为 abcabc(b)b；
以 abcabcb(b) 开始的最长字符串为 abcabcb(b)。

链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-by-leetc-2/


总的时间复杂度（O(n*n)） (strings 包中也有for循环)
*/

func lengthOfLongestSubstring(s string) int {
	length, left, right := 0, 0, 0
	s1 := s[left:right]
	for ; right < len(s); right++ {
		index := strings.IndexByte(s1, s[right])
		if index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > length {
			length = len(s1)
		}
	}
	return length
}

/*
2、题目描述：数组中最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] // 可以写个方法求出 最短的字符串，作为第一个
	for _, val := range strs {
		for strings.Index(string(val), prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return string(prefix)
}

/**
3、题目描述： 最小覆盖子串
最小覆盖子串 （重点）
题目描述：
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
	注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
	示例1：
		输入：s = "ADOBECODEBANC", t = "ABC"
		输出："BANC"

	示例2：
		输入：s = "a", t = "a"
		输出："a"

	思路：
	方法一：滑动窗口
思路和算法

本问题要求我们返回字符串 ss 中包含字符串 tt 的全部字符的最小窗口。我们称包含 t 的全部字母的窗口为「可行」窗口。

我们可以用滑动窗口的思想解决这个问题。在滑动窗口类型的问题中都会有两个指针，一个用于「延伸」现有窗口的 r 指针，
和一个用于「收缩」窗口的 l指针。在任意时刻，只有一个指针运动，而另一个保持静止。
我们在 s 上滑动窗口，通过移动 r 指针不断扩张窗口。当窗口包含 t 全部所需的字符后，如果能收缩，我们就收缩窗口直到得到最小窗口。

参考流程图分析：
https://github.com/zlbonly/simple-go-algorithm/blob/master/pics/minWindow.gif

*/
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	asrL, asrR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}

		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				asrL, asrR = l, l+len
			}

			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}

	if asrL == -1 {
		return ""
	}
	return s[asrL:asrR]
}

/**
	4、最长公共子序列（两个字符串）
最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例1：
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。

示例2：
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。


动态转移方程：

			  {			dp[i-1][j-1] + 1 ,  text1[i-1] = text[j-1]
dp[i][j]  =
              {			max(dp[i-1][j],dp[i][j-1]), text1[i-1]  != text2[j-1]

https://leetcode-cn.com/problems/longest-common-subsequence/solution/zui-chang-gong-gong-zi-xu-lie-tu-jie-dpz-6mvz/
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

/*
5、题目描述：两数子和（I）
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

func twoNum(nums []int, target int) []int {
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
6、题目描述：三数之和
	排序 + 双指针
本题的难点在于如何去除重复解。
算法流程：
特判，对于数组长度 n，如果数组为 null 或者数组长度小于 3，返回 []。
对数组进行排序。
遍历排序后数组：
若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
对于重复元素：跳过，避免出现重复解
令左指针 L=i+1，右指针 R=n−1，当L<R 时，执行循环：
当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
若和大于 0，说明 nums[R] 太大，R 左移
若和小于 0，说明 nums[L] 太小，L 右移
复杂度分析
时间复杂度O(n2)，数组排序 O(NlogN)，遍历数组 O(n)，双指针遍历 O(n)，总体 N*N
空间复杂度：O(1)

作者：wu_yan_zu
链接：https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	len := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < len; i++ {
		L := i + 1
		R := len - 1
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for L < R {
			temp := nums[i] + nums[L] + nums[R]
			if temp == 0 {
				arr := []int{nums[i], nums[L], nums[R]}
				ret = append(ret, arr)
				for L < R && nums[L] == nums[L+1] {
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
题目描述： 删掉一个元素以后全为 1 的最长子数组
给你一个二进制数组 nums ，你需要从中删掉一个元素。请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。如果不存在这样的子数组，请返回 0
提示 1：
输入：nums = [1,1,0,1]
输出：3
解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。

解题思路：
1、设置两个指针 left 和 right 让其分别指向窗口的两端。其中 right 指针是主动移动的，而 left 指针是被迫移动的。其中，count 为窗口里 0 的个数，初始化为 0。

2、首先，看 right 指针位置上的元素是否为 0：若为 0，则 count + 1。当滑动窗口中 0 的个数 count 超过了 1 时(因为只让删掉一个元素)，则就应该将 left 指针向右移动一位了，并且要看 left 指针位置上的值是否为 0，若为 0 则在 left 指针向右移动时，应先将 count - 1 ，以保证滑动窗口中 0 的个数能够及时更新。
然后，判断完大小后，再使 right 指针向右移动一位，进行下一个元素的判断。

3、因为右指针是主动移动的，右指针向右移动之前，若 nums[right] == 0，则count 的值应该 + 1.
4、因为左指针是被动移动的，左指针向右移动之前，若 nums[left] == 0，则 count 的值应该 - 1.

*/
func longestSubarray(nums []int) int {

	// 左指针，右指针，窗口中0的个树，最大值
	left, right, zeroCount, maxLength := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}

		// 只让删除一个元素
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		// 因为要删除一个元素，所以元素的长度不需要+1
		maxLength = max(maxLength, right-left)
		right++
	}
	return maxLength
}

/**
题目描述： 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

demo：
	输入：target = 9
	输出：[[2,3,4],[4,5]]

解题思路：findContinuousSequence 及其注释

*/

func findContinuousSequence(target int) [][]int {

	// plow，phigh 两个起点，相当于动态窗口的两边，根据其窗口内的值来确定窗口的位置和大小
	plow, phigh, cur := 1, 2, 0
	result := make([][]int, 0)
	for plow < phigh {
		// 由于是连续的，差为1的等差序列，那么求和公式（a0 + an）*n /2
		cur = (plow + phigh) * (phigh - plow + 1) / 2
		// 相等，那么将窗口范围内的所有数添加进结果集。
		if cur == target {
			temp := make([]int, 0)
			for i := plow; i <= phigh; i++ {
				temp = append(temp, i)
			}
			result = append(result, temp)
			plow++
		} else if cur < target { // 如果当前窗口内的值和小于target,那么右边窗口移动一下
			phigh++
		} else { // 如果当前窗口内的值和大于sum，那么左边的窗口移动一下
			plow++
		}
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
