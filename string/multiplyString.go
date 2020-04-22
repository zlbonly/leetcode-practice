package main

import (
	"fmt"
	"strconv"
)

/*给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

func main() {
	s1 := "123"
	s2 := "456"
	ret := multiplay(s1, s2)
	fmt.Println("字符串s1 * s2 :", ret)
}

/*方法一：普通竖式
思路

竖式运算思想，以 num1 为 123，num2 为 456 为例分析：
图片地址：https://pic.leetcode-cn.com/d24bf3174a878890e1273fbe35426ecdfa932c33efb464ed3602f4d149ed343a

遍历 num2 每一位与 num1 进行相乘，将每一步的结果进行累加。

注意：

num2 除了第一位的其他位与 num1 运算的结果需要 补0
计算字符串数字累加其实就是 415. 字符串相加
实现
*/

/**
 * 计算形式
 *    num1
 *  x num2
 *  ------
 *  result
 */
func multiplay(s1 string, s2 string) string {

	if len(s1) == 0 || len(s2) == 0 {
		return "0"
	}
	res := "0" // 保存计算结果
	// num2 逐位与 num1 相乘
	for i := len(s2) - 1; i >= 0; i-- {
		carry := 0 // 保存 num2 第i位数字与 num1 相乘的结果
		temp := "" // 补0
		for j := 0; j < len(s2)-1-i; j++ {
			temp += "0"
		}
		n2, _ := strconv.Atoi(string(s2[i]))

		// num2 的第 i 位数字 n2 与 num1 相乘
		for j := len(s1) - 1; j >= 0; j-- {
			n1, _ := strconv.Atoi(string(s1[j]))
			product := (n1*n2 + carry) % 10
			temp += strconv.Itoa(product)
			carry = (n1*n2 + carry) / 10
		}
		fmt.Println(temp)
		fmt.Println(res)
		// 将当前结果与新计算的结果求和作为新的结果
		res = sumString(res, reverseString(temp))
	}
	return res
}

/**
 * 对两个字符串数字进行相加，返回字符串形式的和
 */
func sumString(s1 string, s2 string) string {

	temp := ""
	carray := 0

	for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		tempa := 0
		tempb := 0
		if i >= 0 {
			tempa, _ = strconv.Atoi(string(s1[i]))
		}
		if j >= 0 {
			tempb, _ = strconv.Atoi(string(s2[j]))
		}
		sum := (tempa + tempb + carray) % 10
		temp += strconv.Itoa(sum)
		carray = (tempa + tempb + carray) / 10
	}
	return reverseString(temp)
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
