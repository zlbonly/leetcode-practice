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
题解链接：
https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
 	环行链表入口节点：

	解题思路：
这类链表题目一般都是使用双指针法解决的，例如寻找距离尾部第K个节点、寻找环入口、寻找公共尾部入口等。
算法流程：
双指针第一次相遇： 设两指针 fast，slow 指向链表头部 head，fast 每轮走 2 步，slow 每轮走 1 步；

第一种结果： fast 指针走过链表末端，说明链表无环，直接返回 null；

TIPS: 若有环，两指针一定会相遇。因为每走 1 轮，fast 与 slow 的间距 +1，fast 终会追上 slow；
第二种结果： 当fast == slow时， 两指针在环中 第一次相遇 。下面分析此时fast 与 slow走过的 步数关系 ：

设链表共有 a+b 个节点，其中 链表头部到链表入口 有 a 个节点（不计链表入口节点）， 链表环 有 b 个节点（这里需要注意，a 和 b 是未知数，例如图解上链表 a=4 , b=5）；设两指针分别走了 f，s 步，则有：
fast 走的步数是slow步数的 2 倍，即 f = 2s（解析： fast 每轮走 2 步）
fast 比 slow多走了 n 个环的长度，即 f = s + nb（ 解析： 双指针都走过 a 步，然后在环内绕圈直到重合，重合时 fast 比 slow 多走 环的长度整数倍 ）；
以上两式相减得：f = 2nb，s = nb，即fast和slow 指针分别走了 2n，n 个 环的周长 （注意： n 是未知数，不同链表的情况不同）。
目前情况分析：

如果让指针从链表头部一直向前走并统计步数k，那么所有 走到链表入口节点时的步数 是：k=a+nb（先走 a 步到入口节点，之后每绕 1 圈环（ b 步）都会再次到入口节点）。
而目前，slow 指针走过的步数为 nb 步。因此，我们只要想办法让 slow 再走 a 步停下来，就可以到环的入口。
但是我们不知道 a 的值，该怎么办？依然是使用双指针法。我们构建一个指针，此指针需要有以下性质：此指针和slow 一起向前走 a 步后，两者在入口节点重合。那么从哪里走到入口节点需要 a 步？答案是链表头部head。
双指针第二次相遇：

slow指针 位置不变 ，将fast指针重新 指向链表头部节点 ；slow和fast同时每轮向前走 1 步；
TIPS：此时 f = 0，s = nb ；
当 fast 指针走到f = a 步时，slow 指针走到步s = a+nb 此时 两指针重合，并同时指向链表环入口 。
返回slow指针指向的节点。

复杂度分析：
时间复杂度 O(N) ：第二次相遇中，慢指针须走步数 a < a + b；第一次相遇中，慢指针须走步数 a + b - x < a + b，其中 x 为双指针重合点与环入口距离；因此总体为线性复杂度；
空间复杂度 O(1)：双指针使用常数大小的额外空间。

*/
func detectCycle(head *Node) *Node {
	if head.Next == nil || head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	fast = head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	return fast
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
