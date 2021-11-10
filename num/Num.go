package main

import "math"

/**
1、整数反转
2、
3、

*/
/*
	题目1：整数反转
	整数反转 12345 => 54321  ,-321 => 123。
	将12345 % 10 得到5，之后将12345 / 10
	2、将1234 % 10 得到4，再将1234 / 10
	3、将123 % 10 得到3，再将123 / 10
	4、将12 % 10 得到2，再将12 / 10
	5、将1 % 10 得到1，再将1 / 10
链接：https://leetcode-cn.com/problems/reverse-integer/solution/tu-jie-7-zheng-shu-fan-zhuan-by-wang_ni_ma/
*/
func reverseInt32(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10 // 求尾部数
		// maxInt32 最大值 2147483647
		if x > math.MaxInt32 || (x == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		// maxMinInt32  最小值 -2147483648
		if x < math.MinInt32 || (x == math.MinInt32/10 && pop < -8) {
			return 0
		}
		res = res*10 + pop
		x = x / 10
	}
	return res
}
