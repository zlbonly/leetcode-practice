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

/*
插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

func insertionSortList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dump := new(Node)
	for head != nil {
		prev := dump
		next := head.Next
		// 寻找插入点
		for prev.Next != nil && prev.Next.Data.(int) < head.Data.(int) {
			prev = prev.Next
		}
		head.Next = prev.Next
		prev.Next = head
		head = next
	}
	return dump.Next
}

/**

1、题目
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，
即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，
它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

示例：
给定一个链表: 1->2->3->4->5, 和 k = 2.
返回链表 4->5.


1、解决思路
快慢指针
定义两个指针，快指针 fastfast， 慢指针 lowlow .
让 fastfast 先向前移动 kk 个位置，然后 lowlow 和 fastfast 再一起向前移动 .
当 fastfast 到达链表尾部，返回 lowlow .

*/
func getKthFromEnd(head *Node, k int) *Node {

	fmt.Println(k)
	slow := head
	fast := head

	for fast != nil {
		fast = fast.Next
		if k == 0 {
			slow = slow.Next
		} else {
			k--
		}
	}
	return slow
}

/**

给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。
示例 1：

输入：[1,2,3,4,5]
输出：此列表中的结点 3 (序列化形式：[3,4,5])
返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
注意，我们返回了一个 ListNode 类型的对象 ans，这样：
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
示例 2：

输入：[1,2,3,4,5,6]
输出：此列表中的结点 4 (序列化形式：[4,5,6])
由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。

解题思路： 快慢指针
本题可使用快慢双指针解决:
初始时两个指针均指向头指针。
如果快指针为空或者快指针的next为空，那么此时慢指针即为答案，移动结束。
每轮移动，快指针需要移动两次，慢指针需要移动一次。跳转步骤二。


*/
func middleNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	prev := head
	cur := head
	for cur != nil && cur.Next != nil {
		prev = prev.Next
		cur = cur.Next.Next
	}
	return prev
}

/*给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，
这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，
时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。


解决思路：

将奇节点放在一个链表里，偶链表放在另一个链表里。然后把偶链表接在奇链表的尾部。
算法
这个解法非常符合直觉思路也很简单。但是要写一个精确且没有 bug 的代码还是需要进行一番思索的。
一个 LinkedList 需要一个头指针和一个尾指针来支持双端操作。
我们用变量 head 和 odd 保存奇链表的头和尾指针。 evenHead 和 even 保存偶链表的头和尾指针。
算法会遍历原链表一次并把奇节点放到奇链表里去、偶节点放到偶链表里去。遍历整个链表我们至少需要一个指针作为迭代器。
这里 odd 指针和 even 指针不仅仅是尾指针，也可以扮演原链表迭代器的角色。

复杂度分析

时间复杂度： O(n)O(n) 。总共有 nn 个节点，我们每个遍历一次。

空间复杂度： O(1)O(1) 。我们只需要 4 个指针。
*/

func oddEvenList(head *Node) *Node {
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
