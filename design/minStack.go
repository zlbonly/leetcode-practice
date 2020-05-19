package main

/**
1、题目描述
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素

2、解题思路

i）辅助栈
要做出这道题目，首先要理解栈结构先进后出的性质。
对于栈来说，如果一个元素 a 在入栈时，栈里有其它的元素 b, c, d，那么无论这个栈在之后经历了什么操作，
只要 a 在栈中，b, c, d 就一定在栈中，因为在 a 被弹出之前，b, c, d 不会被弹出。
因此，在操作过程中的任意一个时刻，只要栈顶的元素是 a，那么我们就可以确定栈里面现在的元素一定是 a, b, c, d。
那么，我们可以在每个元素 a 入栈时把当前栈的最小值 m 存储起来。在这之后无论何时，如果栈顶元素是 a，我们就可以直接返回存储的最小值 m。

ii) 自定义栈结构， 同时存放最小值和元素

算法分析 ： 空间复杂度 O（n）
          时间复杂度 O（1）

*/
type MinStack struct {
	elems []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.elems = append(this.elems, x)
	if len(this.mins) == 0 || this.GetMin() >= x {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	elem := this.Top()
	this.elems = this.elems[:len(this.elems)-1]
	if elem <= this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {

	if len(this.elems) == 0 {
		panic("empty stack")
	}
	elem := this.elems[len(this.elems)-1]
	return elem
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		panic("empty stack")
	}
	elem := this.mins[len(this.mins)-1]
	return elem
}
