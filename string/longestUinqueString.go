package main

import (
	"strings"
)

func main() {
	// 1、 无重复字符的最长子串
}

/*
题目：无重复字符的最长子串
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
