package main

// 每个结点包括两个部分： 一个是存储数据元素的数据域，另一个是存储下一个结点地址的指针域。
// 头结点的数据域可以不存储任何信息，头结点的指针域存储指向第一个结点的指针（即第一个元素结点的存储位置）。头结点的作用是使所有链表（包括空表）的头指针非空，并使对单链表的插入、删除操作不需要区分是否为空表或是否在第一个位置进行，从而与其他位置的插入、删除操作一致。
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
1、题目描述：链表翻转
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。

复杂度分析
时间复杂度：O(n)，其中 nn 是链表的长度。需要遍历链表一次。
空间复杂度：O(1)
pre + curr
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/**
递归实现 链表翻转
*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
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
	2、合并两个排序的链表
	输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
	解题思路：递归：
	1、判断两个链表都是有序的 ，如果某个链表为空，那就直接返回另外一个有序链表
	2、首先，我们定义一个result指针，比较两个链表的第一个元素哪个比较小，result指向小的那个链表
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//如果有一条链是nil，直接返回另外一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 定义一个结果节点
	var res *ListNode
	// 当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}

/**
3 题目描述 合并k个有序链表
解题思路
将链表拆分为lists[:len(lists)/2]和lists[len(lists)/2:]两部分，
分别使用mergeKLists进行递归直至lists中只有一个链表时返回，
然后返回的结果两两合并，最终合并为一个链表。流程大致如下所示

方案：
1、分治
2、优先队列 （使用priority_queue队列库） 先把把所有list的节点入队列，然后 使用尾插法。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}
	// 分治法
	return mergeTwoLists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

/**
	4、题目
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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

/*
5、题目描述
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

/**
	6 、分割链表
	问题描述： 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
		你应当 保留 两个分区中每个节点的初始相对位置。

示例：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

参考解题连接：https://leetcode-cn.com/problems/partition-list/solution/fen-ge-lian-biao-by-leetcode-solution-7ade/
*/

func partition(head *Node, x int) *Node {
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

/***
	7、删除链表的倒数第 N 个结点
	我们也可以在遍历链表的同时将所有节点依次入栈。根据栈「先进后出」的原则，我们弹出栈的第 nn 个节点就是需要删除的节点，并且目前栈顶的节点就是待删除节点的前驱节点。这样一来，删除操作就变得十分方便了。
复杂度分析
时间复杂度：O(L)，其中 L 是链表的长度。
空间复杂度：O(L)，其中 L是链表的长度。主要为栈的开销。
参考链接：	https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

/*
8、题目描述：对链表进行 插入排序

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

数组的插入排序，是两重循环将 n 个元素的数列分为已有序和无序两个部分，第一重是遍法一：定义一个新节点，每次从原链表中取出要排序的节点后与链表中的进行比较
为了第一个结点和后边结点插入方法相同，创建一个新的带表头结点的链表，即 ListNode dummyHead = new ListNode(0)。每次取原链表中的一个值，然后在新链表中找到该值的插入位置。

此时应注意，需要找到小于当前值的最后一个。找到之后，将当前值插入到新链表，直到原链表为空

链接：https://leetcode-cn.com/problems/insertion-sort-list/solution/dui-lian-biao-jin-xing-cha-ru-pai-xu-lia-raix/：


*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 创建带表头结点的新链表
	dummyHead := &ListNode{}

	// 用于依次取原链表的值
	cur := head
	// 用于寻找插入位置
	prev := dummyHead

	for cur != nil {
		// 提前记录下当前 p 指针指向的节点，防止原链表断链，取当前值的下一个
		next := cur.Next
		prev = dummyHead
		// 寻找插入点
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return dummyHead.Next
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
func middleNode(head *ListNode) *ListNode {
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

/*
	 题目描述：奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，
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

时间复杂度： O(n) 。总共有 nn个节点，我们每个遍历一次。

空间复杂度： O(1)。我们只需要 4 个指针。
参考解题链接：https://leetcode-cn.com/problems/odd-even-linked-list/solution/qi-ou-lian-biao-by-leetcode-solution/

*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = even.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
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

解题步骤：
	1、判断极端case
	2、使用快慢指针找到链表节点
	3、reverse 逆序后半部分
	4、check 从头，中点，开始比较是否相同

参考链接：https://leetcode-cn.com/problems/aMhZSa/submissions/
*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow = reverseList(slow.Next)
	for slow != nil {
		if head.Val != slow.Val {
			return false
		}

		head = head.Next
		slow = slow.Next
	}
	return true
}

/*

1、问题描述：  重排链表
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

1、解决方案I
	方法一：线性表
	因为链表不支持下标访问，所以我们无法随机访问链表中任意位置的元素。
	因此比较容易想到的一个方法是，我们利用线性表存储该链表，然后利用线性表可以下标访问的特点，直接按顺序访问指定元素，重建该链表即可。

	复杂度分析
	时间复杂度：O(N)，其中 N 是链表中的ß点数。
	空间复杂度：O(N)，其中 N 是链表中的节点数。主要为线性表的开销。

2、解决方案II

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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	start, end := 0, len(nodes)-1
	for start < end {
		nodes[start].Next = nodes[end]
		start++
		if start == end {
			break
		}
		nodes[end].Next = nodes[start]
		end--
	}
	nodes[start].Next = nil
}

/**

环形链表
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

题目地址：https://leetcode-cn.com/problems/linked-list-cycle/
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
