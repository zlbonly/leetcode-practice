package main

import "fmt"

/*type Student struct {
	Name string
	Age int
	Address []int
}

type Student2 struct {
	Name string
	Age int
}*/
func main() {

	/*node1 := &ListNode{Val:1}
	node2 := &ListNode{Val:2}
	node3 := &ListNode{Val:3}
	node4 := &ListNode{Val:3}
	node5 := &ListNode{Val:4}
	node6 := &ListNode{Val:4}
	node7 := &ListNode{Val:6}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	t := removenode(node1)*/

	node1 := &ListNode2{
		Val:    1,
		Next:   nil,
		Random: nil,
	}
	node2 := &ListNode2{
		Val:    2,
		Next:   nil,
		Random: nil,
	}
	node3 := &ListNode2{
		Val:    3,
		Next:   nil,
		Random: nil,
	}

	node1.Next = node2
	node2.Next = node3
	node1.Random = node3
	node2.Random = node1

	t := copyRandomList(node1)
	for t != nil {
		if t.Random != nil {
			fmt.Printf("Val=%v,random=%v ", t.Val, t.Random.Val)

		}
		t = t.Next
	}
}

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

func removenode(head *ListNode) *ListNode {
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

func removeUnique(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dump := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev, cur := dump, head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev = prev.Next
			cur = cur.Next
		}
	}
	return dump.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseNearBy(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dump := &ListNode{
		Val: -1,
		//Next: head,
	}
	prev := dump

	for head != nil {
		cur1 := head
		cur2 := head.Next
		cur1.Next = cur2.Next
		cur2.Next = cur1
		prev.Next = cur2
		prev = cur1
		head = head.Next
	}
	return dump.Next
}
