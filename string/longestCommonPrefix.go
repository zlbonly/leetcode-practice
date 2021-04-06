package main

/*
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

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//s1 := []string{"flower","flow","flight"}

	s1 := []string{"flow", "flo", "flight"}

	ret := testCommonPrefix(s1)
	fmt.Println("最长子前缀：%s", ret)

	s2 := "ADOBECODEBANC"
	s3 := "ABC"
	fmt.Println(minWindow(s2, s3))
}

/*
	解题思路：
	1、长公共前缀，一定是公共的。假定我们现在从第一个元素中寻找公共前缀。那么首先，
		我们将第一个元素设置为基准元素 x0。在这个例子里，就是 flow。[“flow”,"flower","flight”]
 	2、那么很容易得到一个条件。最长公共前缀的长度，一定小于 flow。
	3、然后我们依次将基准元素和后面的元素（假定后面的元素依次为 x0,x1….），进行一一匹配对比。
	4、
		1）如果 strings.Index(x1,x) == 0，则直接跳过（因为此时 x 就是 x1 的最长公共前缀），
		对比下一个元素。（如 flower 和 flow 进行比较）

		2）如果 strings.Index(x1,x) != 0，则截取掉基准元素 x 的最后一个元素，再次和 x1 进行比较，
			直至满足 string.Index(x1,x) == 0，则此时截取后的x为x和x1的最长公共前缀。
			（如 flight 和 flow 进行比较，依次截取出 flow-flo-fl，
			直到 fl 被截取出，此时 fl 为 flight 和 flow 的最长公共前缀）
	5、重复第三步，直至对比完最后一个元素x(N)，
	6、最后，记得处理一下临界条件。。
*/
func testCommonPrefix(strs []string) string {

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

本问题要求我们返回字符串 ss 中包含字符串 tt 的全部字符的最小窗口。我们称包含 tt 的全部字母的窗口为「可行」窗口。

我们可以用滑动窗口的思想解决这个问题。在滑动窗口类型的问题中都会有两个指针，一个用于「延伸」现有窗口的 rr 指针，
和一个用于「收缩」窗口的 ll 指针。在任意时刻，只有一个指针运动，而另一个保持静止。
我们在 ss 上滑动窗口，通过移动 rr 指针不断扩张窗口。当窗口包含 tt 全部所需的字符后，如果能收缩，我们就收缩窗口直到得到最小窗口。

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
