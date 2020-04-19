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
	"strings"
)

func main() {
	//s1 := []string{"flower","flow","flight"}

	s1 := []string{"flow", "flo", "flight"}

	ret := testCommonPrefix(s1)
	fmt.Println("最长子前缀：%s", ret)

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
