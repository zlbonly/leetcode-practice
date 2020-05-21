package main

import "fmt"

/**
1、题目描述
	用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

2、解题思路
	题目只要求实现 加入队尾appendTail() 和 删除队首deleteHead() 两个函数的正常工作，
    因此我们可以设计栈 A 用于加入队尾操作，栈 B 用于将元素倒序，从而实现删除队首元素。

加入队尾 appendTail()函数： 将数字 val 加入栈 A 即可。
删除队首deleteHead()函数： 有以下三种情况。
当栈 B 不为空： B中仍有已完成倒序的元素，因此直接返回 B 的栈顶元素。
否则，当 A 为空： 即两个栈都为空，无元素，因此返回 -1−1 。
否则： 将栈 A 元素全部转移至栈 B 中，实现元素倒序，并返回栈 B 的栈顶元素

3、算法分析
时间复杂度： appendTail()函数为 O(1) ；deleteHead() 函数在 NN 次队首元素删除操作中总共需完成 NN 个元素的倒序。
空间复杂度 O(N)： 最差情况下，栈 A 和 B 共保存 NN 个元素。
*/

type stack []int

// 入栈
func (s *stack) Push(value int) {
	*s = append(*s, value)
}

// 出栈
func (s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n-1]
	*s = (*s)[:n-1]
	return res
}

/**** 实现队列****/

type MyQueue struct {
	in  stack
	out stack
}

func New() MyQueue {
	return MyQueue{}
}

// 入队

func (this *MyQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *MyQueue) DeleteHead() int {
	if this.out != nil {
		this.out.Pop()
	} else if len(this.in) != 0 {
		for len(this.in) != 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}

	return -1
}

func main() {
	myQueue := New()
	myQueue.AppendTail(1)
	myQueue.AppendTail(2)
	myQueue.AppendTail(3)
	a := myQueue.DeleteHead()
	fmt.Print(a)
}
