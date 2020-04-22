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
