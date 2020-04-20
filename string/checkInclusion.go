package main

import "fmt"

/*给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False

注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间*/

func main() {
	s1 := "ba"
	s2 := "idbaooo"
	ret := checkInclusion(s1, s2)

	if ret {
		fmt.Println("s1 是 s2 的子串")
	} else {
		fmt.Println("s1 不是 s2 的子串")
	}
}

/*解决思路：
1、暴力破解法（求出所有的子排列，判断s1的所有 子排列是否在s2中）
2、滑动窗口方法
	1)参考了官方解题的思路也是滑动窗口。
但区别是s1,s2的字符统计不用分2个数组了，先根据s1初始化一个数组，统计字符出现的次数。然后滑动窗口处理s2的时候，把对应的字符减1。
最后用来统计的数组全部是0的时候，就说明他们完全一样了。
	2)
这道题是求s2是否包含s1的排列，也就是说s1中所有的元素可以随机组合。
因此，我们只能通过比较s1长度的s2子串和s1中字符的个数是否相同来进行判断（只要元素个数相同，
s1就可以组合成当前s2子串的样子）。那么我们可以定义一个26字母的字母表数组来存储个数。两个数组相减如果都为0，
那么肯定相同；如果不为0，那么此时子链不相同，将s1整体在s2的位置右移一位，也就是说将s2此时与s1对应的头部的元素的个数减一，将要对应的下一位元素的个数加一，这样就完成的s1数组对于s2的滑动比较。
*/

func checkInclusion(s1 string, s2 string) bool {
	size := len(s1)
	charDiff := make([]int, 26)
	for _, r := range s1 {
		charDiff[r-'a'] += 1
	}
	for index, newChar := range s2 {
		charDiff[newChar-'a'] -= 1
		if index-size >= 0 {
			outChar := s2[index-size]
			charDiff[outChar-'a'] += 1
		}
		if charDiff[newChar-'a'] == 0 {
			same := true
			for _, diff := range charDiff {
				if diff != 0 {
					same = false
					break
				}
			}
			if same {
				return true
			}
		}
	}
	return false
}
