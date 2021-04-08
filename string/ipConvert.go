package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
给你一个IP4的地址，请转成十进制整数 （2021-03 腾讯-PCG-前端）
手撕算法: ip地址转整数 （2021-03 腾讯-TEG-后端）
ip字符串转整型（2021-02 快手-效率工程-后端）

题目描述
ip地址与整数的转换。

例如，ip地址为10.0.3.193，把每段拆分成一个二进制形式组合起来为00001010 00000000 00000011 11000001，然后把这个二进制数转变成十进制整数就是167773121。

题目分析
借助位运算实现。如IP10.0.3.193，将10左移24位，0左移16位，3左移8位，193左移0位。4个seg或运算，即为结果。




*/

func IpConvertInt(ip string) int {
	strArr := strings.Split(ip, ".")
	sLen := len(strArr)
	ret := 0
	for i := 0; i < sLen; i++ {
		temp, _ := strconv.Atoi(strArr[i])
		ret = ret<<8 | temp
	}
	fmt.Println(ret)
	return ret
}

func IntConvertIp(ipInt int) string {

	res := make([]string, 4)

	for i := 0; i < 4; i++ {
		temp := strconv.Itoa(ipInt & 255)
		fmt.Println(temp)
		res[i] = temp
		ipInt = ipInt >> 8
	}

	str := ""

	for i := 4 - 1; i >= 0; i-- {
		if i == 0 {
			str += res[i]
		} else {
			str += res[i] + "."
		}
	}
	fmt.Println(str)
	return str

}

func main() {
	ip := "10.0.3.193"
	ret := IpConvertInt(ip)

	IntConvertIp(ret)

}
