package main

/**
1、 链表翻转（非递归和非递归）
2、 链表翻转（从指定位置）
3、 合并两个有序链表
4、 合并N个有序链表
5、 链表中倒数第k个节点
6、 两两交换链表中的节点
7、 分割链表
8、 删除链表的倒数第 N 个结点
9、 对链表进行插入排序
10、链表的中间结点
11、奇偶链表
12、链表是否为回文链表
13、重排链表
14、是否有环
15、环入口
16、相交链表节点
17、删除排序链表中的重复元素
18、删除排序链表中的重复元素II（只保留没有出现过的元素）
19、链表求和
20、链表反转（中间节点链表反转）
21、随机链表（深拷贝）
*/
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

在遍历链表时，将当前节点的 next 指针改为指向前一个节点。
由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。
在更改引用之前，还需要存储后一个节点。最后返回新的头引用。

复杂度分析
时间复杂度：O(n)，其中 nn 是链表的长度。需要遍历链表一次。
空间复杂度：O(1)

链表反转解题步骤：
1、定一个空头节点prev
2、存储下一个next := head.next 节点
3、head.next 指向空prev节点
4、head 赋值给prev节点
5、head 节点下移
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

/**
2、递归实现 链表翻转
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
	2、链表反转II
	题目描述：
			给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
	示例1：
		1 -> 2 -> 3 -> 4 -> 5
		1 -> 4 -> 3 -> 2 -> 5
	输入：head = [1,2,3,4,5], left = 2, right = 4
	输出：[1,4,3,2,5]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dumpNode := &ListNode{Val: -1}
	dumpNode.Next = head
	pre := dumpNode

	// 第一步 从虚拟头节点走 left-1步，来到left节点的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第2步 从pre再走right-left+1步 来到right节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第三步 切断出一个子链表（截取left ~ right 总节点为一个新链表）
	betWeenNode := pre.Next
	curr := rightNode.Next

	// 注意切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第四步，中间的链表子区间进行反转
	reverseNode := reverseList(betWeenNode)

	// 第五步 把反转后的链表接回到原链表中
	pre.Next = reverseNode
	for reverseNode.Next != nil {
		reverseNode = reverseNode.Next
	}
	reverseNode.Next = curr
	return dumpNode.Next
}

/**

	3、合并两个排序的链表
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
 4、题目描述 合并k个有序链表
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
	5、题目 链表中倒数第k个节点
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
6、题目描述  两两交换链表中的节点
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
	7 、分割链表
	问题描述： 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
		你应当 保留 两个分区中每个节点的初始相对位置。

示例：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

参考解题连接：https://leetcode-cn.com/problems/partition-list/solution/fen-ge-lian-biao-by-leetcode-solution-7ade/
*/

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	smallHead := small
	largeHead := large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

/***
	8、删除链表的倒数第 N 个结点
	1、双指针解决
	时间复杂度：O(L)，其中 L 是链表的长度。
	空间复杂度：O(1)。
时间复杂度：O(L)，其中 L 是链表的长度。
空间复杂度：O(L)，其中 L是链表的长度。主要为栈的开销。
参考链接：	https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dump := &ListNode{Val: -1, Next: head}
	first, slow := head, dump
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dump.Next
}

/*
9、题目描述：对链表进行插入排序

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
10、链表的中间结点
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
11、题目描述：奇偶链表
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
	nodeA := &ListNode{Val: -1}
	nodeB := &ListNode{Val: -1}
	headA, headB := nodeA, nodeB
	i := 1
	for head != nil {
		if i%2 == 1 {
			headA.Next = head
			headA = headA.Next
		} else {
			headB.Next = head
			headB = headB.Next
		}
		i++
		head = head.Next
	}
	headB.Next = nil
	headA.Next = nodeB.Next
	return nodeA.Next
}

/**
12、链表是否为回文链表

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

快指针走到末尾，慢指针刚好到中间。其中慢指针将前半部分反转。然后比较。 思路果然666


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

13、问题描述：  重排链表
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
func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
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
	return nodes[0]
}

/**
	14、判断是否有环 （面试时，可以直接给面试官讲 获取环的入口 ，因为肯定要先判断是否有环）
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
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return true
}

/**
15、环的入口
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

概括一下：
根据：
f=2s （快指针每次2步，路程刚好2倍）
f = s + nb (相遇时，刚好多走了n圈）
推出：s = nb
从head结点走到入环点需要走 ： a + nb， 而slow已经走了nb，那么slow再走a步就是入环点了。
如何知道slow刚好走了a步？ 从head开始，和slow指针一起走，相遇时刚好就是a步

复杂度分析：
时间复杂度 O(N) ：第二次相遇中，慢指针须走步数 a < a + b；第一次相遇中，慢指针须走步数 a + b - x < a + b，其中 x 为双指针重合点与环入口距离；因此总体为线性复杂度；
空间复杂度 O(1)：双指针使用常数大小的额外空间。

*/

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
}

/**
16. 相交链表
面试的时候可以先询问面试官 要的时间复杂度。然后给出这两种方案。
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
题目链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists

1、解决思路：1、方法一：哈希集合
思路和算法

1、判断两个链表是否相交，可以使用哈希集合存储链表节点。

2、首先遍历链表headA，并将链表 headA 中的每个节点加入哈希集合中。然后遍历链表 headB，对于遍历到的每个节点，判断该节点是否在哈希集合中：
	如果当前节点不在哈希集合中，则继续遍历下一个节点；
	如果当前节点在哈希集合中，则后面的节点都在哈希集合中，即从当前节点开始的所有节点都在两个链表的相交部分，因此在链表headB 中遍历到的第一个在哈希集合中的节点就是两个链表相交的节点，返回该节点。
	如果链表headB 中的所有节点都不在哈希集合中，则两个链表不相交，返回 null。

	时间复杂度 ： o（m+n）
	空间复杂度 n
	时间复杂度：O(m+n)，其中 m 和 n 是分别是链表 headA 和headB 的长度。需要遍历两个链表各一次。
	空间复杂度：O(m)，其中 m 是链表headA 的长度。需要使用哈希集合存储链表 headA 中的全部节点。

3、双指针
解题步骤：
	1、设长链表A长度为LA，短链表长度LB；
	2、由于速度相同，则在长链表A走完LA长度时，短链表B已经反过头在A上走了LA-LB的长度，剩余要走的长度为LA-(LA-LB) = LB；
	3、之后长链表A要反过头在B上走，剩余要走的长度也是LB；
	4、也就是说目前两个链表“对齐”了。因此，接下来遇到的第一个相同节点便是两个链表的交点。

	复杂度分析
	时间复杂度：O(m+n)其中 m 和 n 是分别是链表 headA 和 headB 的长度。两个指针同时遍历两个链表，每个指针遍历两个链表各一次。
	空间复杂度：O(1)
	参考解题链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/tu-jie-xiang-jiao-lian-biao-by-user7208t/
*/
func getIntersectionNodeI(headA, headB *ListNode) *ListNode {
	vis := make(map[*ListNode]bool, 0)
	for temp := headA; temp != nil; temp = temp.Next {
		vis[temp] = true
	}
	for temp := headB; temp != nil; temp = temp.Next {
		if vis[temp] {
			return temp
		}
	}
	return nil
}

func getIntersectionNodeII(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

/**
17、题目描述：删除排序链表中的重复元素
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。返回同样按升序排列的结果链表。
输入：head = [1,1,2]
输出：[1,2]
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

/**
18、题目描述：删除排序链表中的重复元素
存在一个按升序排列的链表，给你这个链表的头节点 head ，
请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
返回同样按升序排列的结果链表。

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
*/
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dumpNode := &ListNode{Val: -1}
	dumpNode.Next = head

	pre, cur := dumpNode, head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			pre = pre.Next
			cur = cur.Next
		} else {
			for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next
			cur = cur.Next
		}
	}
	return dumpNode.Next
}

/**
题目描述：19、链表求和
给定两个用链表表示的整数，每个节点包含一个数位。这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。
示例 反向存放
	输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
	输出：2 -> 1 -> 9，即912

解题思路：
反向存放：从两个链表头开始相加，处理进位（单位大于10的问题），创建新的链表节点，

正向存放：
示例：
	输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
	输出：9 -> 1 -> 2，即912
解题思路：
	1、利用栈先进后出、计算每一位的和，累加过程处理与反向一致，每一步处理链接节点摆放位置不同
	2、先将输入链表反转，求完和之后在翻转一次。从而使得最终的和也满足正向存放。
*/
func addTwoNumbersI(l1 *ListNode, l2 *ListNode) *ListNode {
	dumpHead := &ListNode{Val: -1}
	cur := dumpHead
	carry := 0
	sum := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		temp := &ListNode{Val: sum % 10}
		carry = sum / 10
		cur.Next = temp
		cur = cur.Next
	}
	return dumpHead.Next
}

/**
题目描述：20 链表反转II （中间节点链表反转）
		给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
		请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
示例：
	输入：head = [1,2,3,4,5], left = 2, right = 4
	输出：[1,4,3,2,5]
解题思路：
	1、定义指针 分别为 g（guard守卫） 和 p(point)，根据m和n的参数，确定g 和 p 的位置。
	将g移动到第一个要反转的节点的前面，将p移动到第一个要反转的节点的位置上。
	例如：
	  1   ->   2  ->  3   -> 4 	-> 5   -> NULL
	  dump ->  1   ->   2  ->  3   -> 4 -> 5  -> NULL
	  dump ->  1   ->   2  ->  3   -> 4 -> 5  -> NULL
				g       p
	2、将p后面的元素删除，然后添加到 g的后面，也即头插法
	3、根据m和n 重复步骤（2）
	4、返回dumpHead.next结点

	链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
*/
func reverseBetWeen(head *ListNode, m int, n int) *ListNode {
	dump := &ListNode{Val: -1, Next: head}
	pre := dump
	g, p := pre, pre.Next

	// 将指针移动到相应位置
	for i := 0; i < m-1; i++ {
		p = p.Next
		g = g.Next
	}
	// 头插法插入结点
	for i := 0; i < n-m; i++ {
		removed := p.Next
		p.Next = p.Next.Next
		removed.Next = g.Next
		g.Next = removed
	}
	return dump.Next
}

/***
	21、题目描述：随机链表深拷贝
		1、给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
		构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。
       新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
		例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
		返回复制链表的头节点。
	链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer
*/

type ListNode2 struct {
	Val    int
	Next   *ListNode2
	Random *ListNode2
}

func copyRandomList(head *ListNode2) *ListNode2 {
	if head == nil {
		return nil
	}
	// 1、在每个原节点后面创建一个新节点
	cur := head
	for cur != nil {
		cur.Next = &ListNode2{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur = cur.Next.Next
	}
	// 2、设置新节点的随机节点
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 3、将两个链表分离
	dump := &ListNode2{Val: -1}
	cur, prev := head, dump
	for cur != nil {
		prev.Next = cur.Next
		prev = prev.Next
		cur = cur.Next.Next
	}
	return dump.Next
}
