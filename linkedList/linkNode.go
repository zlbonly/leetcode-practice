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

func reverseNode(head *Node) *Node {
	if head.Next == nil {
		return head
	}
	last := reverseNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseN(head *Node, n int) *Node {
	successor := &Node{}
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

func createCycle(head *Node, pos int) *Node {
	if head == nil {
		return nil
	}
	num := 0

	temp := &Node{}
	for head.Next != nil {
		if pos == num {
			temp = head
		}
		head = head.Next
		num++
	}

	fmt.Println(temp.Data)
	head.Next = temp
	return head
}

/**
	双指针法 求相交连表的节点
因为如果链表A和链表B相交于D的话,那么说明D结点即在A上又在B上,而D之后的元素自然也就均在A和B上了,因为他们是通过next指针相连的.

如果有相交的结点D的话,每条链的头结点先走完自己的链表长度,然后回头走另外的一条链表,那么两结点一定为相交于D点,因为这时每个头结点走的距离是一样的,都是 AD + BD + DC,而他们每次又都是前进1,所以距离相同,速度又相同,固然一定会在相同的时间走到相同的结点上,即D点

*/
func getIntersectionNode(heada *Node, headb *Node) *Node {
	if heada == nil || headb == nil {
		return nil
	}
	t1 := heada
	t2 := headb
	for t1 != t2 {
		if t1 == nil {
			t1 = headb
		}
		t1 = t1.Next

		if t2 == nil {
			t2 = heada
		}
		t2 = t2.Next
	}

	return t1
}

/**
题目：
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

不过代码上做了合并优化，特别是while循环里面的处理应该好好琢磨琢磨。
特别注意的是这里返回的是dummyHead.next也蛮巧妙的，
省去了链表开头或结尾的情况判断；以及carry进位标记的设置。
*/

func addTwoNumbers(head1 *Node, head2 *Node) *Node {
	dummyHead := &Node{}

	curr := dummyHead
	p1 := head1
	p2 := head2
	carry := 0
	for p1 != nil && p2 != nil {
		x := 0
		if p1 != nil {
			x = p1.Data.(int)
		}

		y := 0
		if p2 != nil {
			y = p2.Data.(int)
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &Node{}
		curr.Next.Data = sum % 10
		curr = curr.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if carry > 0 {
		curr.Next = &Node{}
		curr.Next.Data = carry
	}
	return dummyHead.Next
}
