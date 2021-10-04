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

/**
题目描述1：数字转中文大写
1、今天和同事讨论，头条面试的算法题，他说问了一个金额大小写转换的，即1024 转成壹千零贰拾肆圆整，
然后自己尝试一下了（不考虑小数，金额数字小数点只能到两位了）
2、具体思路如下：逐个读取每个数字，然后对应取中文数字，并加上中文单位，
	最后把一些需要处理的描述，比如零零，换成零，把零万，换成零等。
参考链接：https://www.jianshu.com/p/7c03606dab04
*/
func transfer(num int) string {
	chineseMap := []string{"圆整", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseNum := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	listNum := []int{}
	for ; num > 0; num = num / 10 {
		listNum = append(listNum, num%10)
	}
	n := len(listNum)
	chinese := ""
	//注意这里是倒序的
	for i := n - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	//注意替换顺序
	for {
		copychinese := chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零十", "零", 1)
		copychinese = strings.Replace(copychinese, "零百", "零", 1)
		copychinese = strings.Replace(copychinese, "零千", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零圆", "圆", 1)

		if copychinese == chinese {
			break
		} else {
			chinese = copychinese
		}
	}
	return chinese
}

/**
题目描述2，整数转罗马数字
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，
而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。

解题思路：
	参考解题链接： https://leetcode-cn.com/problems/integer-to-roman/solution/zheng-shu-zhuan-luo-ma-shu-zi-by-leetcod-75rs/
*/

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
