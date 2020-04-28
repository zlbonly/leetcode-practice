package main

import (
	"fmt"
)

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

func (this *List) isEmpty() bool {
	if this.headNode == nil {
		return true
	}
	return false
}

func (this *List) length() int {
	cur := this.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (this *List) add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
	return node
}

func (this *List) append(data Object) {
	node := &Node{Data: data}
	if this.isEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (this *List) insert(index int, data Object) {
	if index < 0 {
		this.add(data)
	} else if index > this.length() {
		this.append(data)
	} else {
		pre := this.headNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}

		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

func (this *List) remove(data Object) {
	pre := this.headNode
	if pre.Data == data {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

func (this *List) removeByIndex(index int) {
	pre := this.headNode
	if index < 0 {
		this.headNode = pre.Next
	} else if index > this.length() {
		fmt.Println("超出链表长度")
	} else {
		count := 0

		for count != (index-1) && pre.Next != nil {
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

func (this *List) isContain(data Object) bool {
	cur := this.headNode

	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (this *List) showList() {
	if !this.isEmpty() {
		cur := this.headNode
		for {
			fmt.Printf("\t%v", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	}
}

/*func (this *List)showReverse()  {
	if!this.isEmpty(){
		cur := this.headNode

		for {
			fmt.Printf("\t%v",cur.Data)
			if
		}
	}
}*/

/**
链表反转 （cur + pre） 双指针
*/
func (this *List) reverseSimply() {
	pre := this.headNode.Next
	cur := this.headNode
	cur.Next = nil
	for pre != nil {
		// 局部反转
		temp := pre.Next
		pre.Next = cur
		cur = pre
		pre = temp
	}
}

/**
解题思路
假定链表是一个栈结构，从原链表头开始逐一取出数据视为出栈，
往新链表头部添加数据视为入栈，原栈数据全部存入新栈后，新栈数据跟原来是反向的。
*/
func (this *List) reverseStack() *Node {
	cur := this.headNode
	var res *Node
	for true {
		if cur == nil {
			break
		}
		res = &Node{cur.Data, res}
		cur = cur.Next
	}
	return res
}

/**
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。

示例:

输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8

解题思路： 双指针法

*/
func (this *List) partitionNode(partition Object) *Node {

	cur := this.headNode
	var res *Node
	var smaller *Node
	var bigger *Node
	for cur != nil {
		if cur.Data.(int) < partition.(int) {
			smaller = &Node{cur.Data, smaller}
		} else {
			bigger = &Node{cur.Data, bigger}
		}
		cur = cur.Next
	}
	for true {
		if bigger == nil {
			break
		}
		res = &Node{bigger.Data, res}
		bigger = bigger.Next
	}
	for true {
		if smaller == nil {
			break
		}

		res = &Node{smaller.Data, res}
		smaller = smaller.Next
	}

	return res
}

/*

题目：
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

解题思路： 采用双指针法
*/

func mergeTwoOrderList(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *Node
	if l1.Data.(int) < l2.Data.(int) {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	pre := head
	for l1 != nil && l2 != nil {
		if l1.Data.(int) < l2.Data.(int) {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return head
}

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.

解题思路：
迭代法：

我们把链表分为两部分，即奇数节点为一部分，偶数节点为一部分，A 指的是交换节点中的前面的节点，B 指的是要交换节点中的后面的节点。在完成它们的交换，我们还得用 prevNode 记录 A 的前驱节点。

算法：

1、firstNode（即 A） 和 secondNode（即 B） 分别遍历偶数节点和奇数节点，即两步看作一步。
2、交换两个节点：
 firstNode.next = secondNode.next
 secondNode.next = firstNode
3、还需要更新 prevNode.next 指向交换后的头。
	prevNode.next = secondNode
4、迭代完成后得到最终的交换结果。


*/
func swapPairs(head *Node) *Node {
	if head == nil {
		return nil
	}
	dump := new(Node) //  # Dummy node acts as the prevNode for the head node
	dump.Next = head
	prev := dump
	for head != nil {
		cur1 := head
		cur2 := head.Next
		prev.Next = cur2
		cur1.Next = cur2.Next
		cur2.Next = cur1
		prev = cur1
		head = cur1.Next
	}
	return dump.Next
}
