package main

import (
	"fmt"
	"strings"
)

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，
标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/

func main() {
	s := " hello world! zlb!"
	ret := reverseWords(s)
	fmt.Println("单词反转后的结果：", ret)

	s = "leetcode"
	wordDict := []string{"leet", "code"}
	result := wordBreak(s, wordDict)
	fmt.Printf("单词拆分，result:%v", result)
}

/*
解题思路
小技巧，首先通过Split拆分，然后逆序遍历，将每个元素去除多余的空格之后添加到strArr2中，然后join起来。
*/
func reverseWords(s string) string {
	strArr1 := strings.Split(s, " ")
	strArr2 := make([]string, 0)
	for i := len(strArr1) - 1; i > 0; i-- {
		if strArr1[i] != "" {
			strArr2 = append(strArr2, strings.TrimSpace(strArr1[i]))
		}
	}
	return strings.Join(strArr2, " ")
}

/**
	2、单词拆分
	题目描述：给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
	说明：拆分时可以重复使用字典中的单词。
		你可以假设字典中没有重复的单词。

	示例 1：
	输入: s = "leetcode", wordDict = ["leet", "code"]
	输出: true
	解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

	示例 2：

	输入: s = "applepenapple", wordDict = ["apple", "pen"]
	输出: true
	解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。

	解法：动态规划
	""	 	l 		e 		e 		t		c 		o		d		e
	true 	false	false   false	true	false	false	false	true

	1、初始化 dp = [false ...false] 长度为n+1,n为字符串长度，dp[i]表死s的前i位是否可以用wordDict中的单词表示
	2、初始化dp[0] = true  空字符串 可以被表示
	3、遍历字符串的所有子串，遍历开始索引i,遍历区间【0，n)
		1) 遍历结束索引j,遍历区间【i+1,n+1）
		2） 若 dp[i] = true 且 s[i...j)在wordlist中，dp[j] = true
			dp[i] = true 说明s的前i位可以用wordDict表示，则s[i...j) 出现在wordDict中，说明s的前j位可以表示

	参考题目解析：https://leetcode-cn.com/problems/word-break/solution/dong-tai-gui-hua-ji-yi-hua-hui-su-zhu-xing-jie-shi/

*/

func wordBreak(s string, wordDict []string) bool {
	word_dict := make(map[string]int)
	for _, v := range wordDict {
		word_dict[v] = 1
	}
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := i + 1; j < length+1; j++ {
			_, ok := word_dict[s[i:j]]
			if dp[i] && ok {
				dp[j] = true
			}
		}
	}
	return dp[length]
}
