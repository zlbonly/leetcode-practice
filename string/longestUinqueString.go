package main

import (
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
	// 7、最长回文子序列（不连续）
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

思路不难，滑动窗口的基本操作而已。

1、初始化左右指针 left = right = 0，把索引区间 [left, right] 称为一个「窗口」。
2、不断地增加 right 指针扩大窗口 [left, right]，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。
此时，停止增加 right，转而不断增加 left 指针缩小窗口 [left, right]，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮结果。
3、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。
4、其中，第 2 步相当于在寻找一个「可行解」，然后第 3 步在优化这个「可行解」，最终找到最优解。左右指针轮流前进，窗口大小增增减减，窗口不断向右滑动。

*/

func minWindow(s string, t string) string {
	//  need数组记录t字符串内字符出现的频率
	//  window数组记录窗口内字符出现的频率
	need, window := map[byte]int{}, map[byte]int{}
	matchLen := 0       // 匹配长度
	minLength := len(s) // 最短匹配字符串长度
	// 窗口左边界，窗口右边界
	windowLeft, windowRight := -1, -1

	for i := range t {
		need[t[i]]++
	}
	// 依次遍历s，l为左指针，r为右指针
	for left, right := 0, 0; right < len(s); right++ {
		window[s[right]]++
		if _, ok := need[s[right]]; ok && window[s[right]] == need[s[right]] {
			matchLen++
		}

		// 达到匹配长度时，要缩减左、右指针之间的长度，得到最短匹配字符串
		for matchLen == len(need) {
			tempLength := right - left + 1 // 窗口长度
			if tempLength <= minLength {
				minLength = tempLength
				windowLeft, windowRight = left, right
			}
			// 左指针指向的字符在t里且数量一致时，匹配长度-1
			if _, ok := need[s[left]]; ok && window[s[left]] == need[s[left]] {
				matchLen--
			}
			// 左指针往前移动
			window[s[left]]--
			left++
		}
	}
	if windowLeft == -1 {
		return ""
	}
	return s[windowLeft : windowRight+1]
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

			  {			dp[i-1][j-1] + 1 ,  text1[i] = text[j]
dp[i][j]  =
              {			max(dp[i-1][j],dp[i][j-1]), text1[i]  != text2[j]

https://leetcode-cn.com/problems/longest-common-subsequence/solution/zui-chang-gong-gong-zi-xu-lie-tu-jie-dpz-6mvz/
*/
func longestCommonSubSequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m-1][n-1]
}

/*func longestCommonSubsequence(text1 string, text2 string) int {
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
}*/

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

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]


注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

排序 + 左右双指针

1、特判，对于数组长度 n，如果数组为 null 或者数组长度小于 3，返回 []。
2、对数组进行排序。
3、遍历排序后数组
	1、若 nums[i]>0因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
	2、对于重复元素：跳过，避免出现重复解
	3、令左指针 L=i+1，右指针 R=n-1，当L<R 时，执行循环：
	4、当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
	5、若和大于 0，说明 nums[R] 太大，R 左移
	6、若和小于 0，说明 nums[L] 太小，L 右移
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
			return ret
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

/***
7、最长回文子序列
	输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。

注意： 最长回文子串 不同的是，回文子序列不要求连续。
题目描述：给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
解题思路： 动态规划
1、于一个子序列而言，如果它是回文子序列，并且长度大于 2，那么将它首尾的两个字符去除之后，它仍然是个回文子序列。因此可以用动态规划的方法计算给定字符串的最长回文子序列
2、用 dp[i][j] 表示字符串 s 的下标范围 [i, j] 内的最长回文子序列的长度。假设字符串 s 的长度为 n，则只有当 0≤i≤j<n 时，才会有 dp[i][j]>0，否则 dp[i][j]=0。
3、由于任何长度为 1 的子序列都是回文子序列，因此动态规划的边界情况是，对任意0≤i<n，都有 dp[i][i]=1。
4、当 i<j 时，计算 dp[i][j] 需要分别考虑 s[i] 和 s[j] 相等和不相等的情况：
	1、如果 s[i]=s[j]，则首先得到 s 的下标范围 [i+1,j−1] 内的最长回文子序列，然后在该子序列的首尾分别添加 s[i] 和 s[j]，即可得到 s 的下标范围 [i, j] 内的最长回文子序列，因此 dp[i][j]=dp[i+1][j−1]+2；
	2、如果 s[i]和s[j]不相等，则 s[i] 和 s[j] 不可能同时作为同一个回文子序列的首尾，因此 dp[i][j]=max(dp[i+1][j],dp[i][j−1])。
	3、由于状态转移方程都是从长度较短的子序列向长度较长的子序列转移，因此需要注意动态规划的循环顺序
			最终得到 dp[0][n−1] 即为字符串 ss 的最长回文子序列的长度。

*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

/**
8、最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

注意： 最长回文子串要求一定是连续的。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

暴力破解法：
暴力求解，列举所有的子串，判断是否为回文串，保存最长的回文串。

时间复杂度：两层 for 循环O(n²），for 循环里边判断是否为回文 O(n），所以时间复杂度为O(n³）。
空间复杂度：O(1），常数个变量。



*/
func longestPalindrome(s string) string {
	length := len(s)
	maxLength := 0
	var ret string

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			temp := s[i : j+1]
			if isPalindromic(temp) && len(temp) > maxLength {
				maxLength = max(maxLength, len(temp))
				ret = temp
			}
		}
	}
	return ret
}

// 判断是否是回文字符串
func isPalindromic(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
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

/**
1、移除元素
	给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
	不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
