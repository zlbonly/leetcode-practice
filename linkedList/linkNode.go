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

/**
合并k个有序链表
解题思路
将链表拆分为lists[:len(lists)/2]和lists[len(lists)/2:]两部分，
分别使用mergeKLists进行递归直至lists中只有一个链表时返回，
然后返回的结果两两合并，最终合并为一个链表。流程大致如下所示
*/
func mergeKLists(lists []*Node) *Node {
	n := len(lists)

	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}
	// 分治法
	return mergeTwoOrderList(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
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

/*func reverseBetween(head *Node,m int,n int) *Node   {

}*/

/*

1、问题描述
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.


2、解决方案

1 -> 2 -> 3 -> 4 -> 5 -> 6
第一步，将链表平均分成两半
1 -> 2 -> 3
4 -> 5 -> 6

第二步，将第二个链表逆序
1 -> 2 -> 3
6 -> 5 -> 4

第三步，依次连接两个链表
1 -> 6 -> 2 -> 5 -> 3 -> 4
*/

func reorderList(head *Node) *Node {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	a := head
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next

	slow.Next = nil
	// 将第二个链表反转
	newHead = reverseNode(newHead)
	// 合并两个连链表
	for newHead != nil {
		temp := newHead.Next
		newHead.Next = head.Next
		head.Next = newHead
		head = newHead.Next
		newHead = temp
	}
	return a
}

/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

解决方案：
时间复杂度：O(n)O(n)，其中 nn 指的是链表的大小。
空间复杂度：O(1)O(1)，我们是一个接着一个的改变指针，我们在堆栈上的堆栈帧不超过 O(1)O(1)。

*/
func isPalindrome(head *Node) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	// 1、找到链表的中点，链表长度奇偶不影响
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//2、将slow之后链表进行断开且反转，最后翻转完成之后pre指向反转链表的头节点
	pre := &Node{
		Data: nil,
		Next: nil,
	}
	for slow != nil {
		temp := slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}
	//3、前后链表进行比较，注意若为奇数链表，后半部分回比前部分多1一个节点，然而我们只比较相同长度的节点值，巧妙地避开这点判断

	fmt.Print(head.Data)
	fmt.Print(pre.Data)
	fmt.Println()

	for head != nil && pre.Next != nil {
		fmt.Println(pre.Data)
		if head.Data != pre.Data {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}

/**

题目地址：https://leetcode-cn.com/problems/linked-list-cycle/
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。




进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

创建环形链表

*/
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

// 快慢指针 判断链表是否有环
func hasCycle(head *Node) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

/**
题解链接：
https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
 	环行链表入口节点：

	解题思路：
这类链表题目一般都是使用双指针法解决的，例如寻找距离尾部第K个节点、寻找环入口、寻找公共尾部入口等。
算法流程：
双指针第一次相遇： 设两指针 fast，slow 指向链表头部 head，fast 每轮走 22 步，slow 每轮走 11 步；

第一种结果： fast 指针走过链表末端，说明链表无环，直接返回 null；

TIPS: 若有环，两指针一定会相遇。因为每走 11 轮，fast 与 slow 的间距 +1+1，fast 终会追上 slow；
第二种结果： 当fast == slow时， 两指针在环中 第一次相遇 。下面分析此时fast 与 slow走过的 步数关系 ：

设链表共有 a+ba+b 个节点，其中 链表头部到链表入口 有 aa 个节点（不计链表入口节点）， 链表环 有 bb 个节点（这里需要注意，aa 和 bb 是未知数，例如图解上链表 a=4a=4 , b=5b=5）；设两指针分别走了 ff，ss 步，则有：
fast 走的步数是slow步数的 22 倍，即 f = 2sf=2s；（解析： fast 每轮走 22 步）
fast 比 slow多走了 nn 个环的长度，即 f = s + nbf=s+nb；（ 解析： 双指针都走过 aa 步，然后在环内绕圈直到重合，重合时 fast 比 slow 多走 环的长度整数倍 ）；
以上两式相减得：f = 2nbf=2nb，s = nbs=nb，即fast和slow 指针分别走了 2n2n，nn 个 环的周长 （注意： nn 是未知数，不同链表的情况不同）。
目前情况分析：

如果让指针从链表头部一直向前走并统计步数k，那么所有 走到链表入口节点时的步数 是：k=a+nb（先走 aa 步到入口节点，之后每绕 11 圈环（ bb 步）都会再次到入口节点）。
而目前，slow 指针走过的步数为 nbnb 步。因此，我们只要想办法让 slow 再走 aa 步停下来，就可以到环的入口。
但是我们不知道 aa 的值，该怎么办？依然是使用双指针法。我们构建一个指针，此指针需要有以下性质：此指针和slow 一起向前走 a 步后，两者在入口节点重合。那么从哪里走到入口节点需要 aa 步？答案是链表头部head。
双指针第二次相遇：

slow指针 位置不变 ，将fast指针重新 指向链表头部节点 ；slow和fast同时每轮向前走 11 步；
TIPS：此时 f = 0f=0，s = nbs=nb ；
当 fast 指针走到f = af=a 步时，slow 指针走到步s = a+nbs=a+nb，此时 两指针重合，并同时指向链表环入口 。
返回slow指针指向的节点。

复杂度分析：
时间复杂度 O(N)O(N) ：第二次相遇中，慢指针须走步数 a < a + ba<a+b；第一次相遇中，慢指针须走步数 a + b - x < a + ba+b−x<a+b，其中 xx 为双指针重合点与环入口距离；因此总体为线性复杂度；
空间复杂度 O(1)O(1) ：双指针使用常数大小的额外空间。

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

/*
	链表反转II
	题目描述：
			给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
	示例1：
		1 -> 2 -> 3 -> 4 -> 5
		1 -> 4 -> 3 -> 2 -> 5
	输入：head = [1,2,3,4,5], left = 2, right = 4
	输出：[1,4,3,2,5]

	解题思路：
		1、我们定义两个指针，分别是 g(guard 守卫)  和 p (ponit) ，我们首先根据方法的参数 m 确定 g 和 p 的位置。
		将  g 移动到第一个要反转的节点的前面，将 p 移动到第一个 要反转的节点的 位置上 。。
		2、将p 后面的元素删除，然后添加到g的后面去，也即头插法
		3、根据 m 和 n 的重复步骤（2）
		4、返回dummyHead.next
*/
func reverseBetween(head *Node, left int, right int) *Node {
	dummyHead := &Node{Data: -1}
	dummyHead.Next = head

	// 初始化指针
	g := dummyHead
	p := dummyHead.Next

	// 将指针移到相应的位置
	for i := 0; i < left-1; i++ {
		g = g.Next
		p = p.Next
	}

	for i := 0; i < right-left; i++ {
		removed := p.Next
		p.Next = p.Next.Next
		removed.Next = g.Next
		g.Next = removed
	}
	return dummyHead.Next
}

/**
	12 、分割链表
	问题描述： 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
		你应当 保留 两个分区中每个节点的初始相对位置。

示例：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

参考解题连接：https://leetcode-cn.com/problems/partition-list/solution/fen-ge-lian-biao-by-leetcode-solution-7ade/
*/

func partion(head *Node, x int) *Node {
	small := &Node{}
	large := &Node{}
	smallHead := small
	largeHead := large

	for head != nil {
		if head.Data.(int) < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	largeHead = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
