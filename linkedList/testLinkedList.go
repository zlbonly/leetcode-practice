package main

import "fmt"

func main() {
	/*list := List{}
	list.append(9)
	list.append(2)
	list.append(3)
	list.append(10)
	list.append(5)
	list.append(5)
	list.append(8)
	list.append(1)*/

	/*list.append("a")
	list.append("b")
	list.append("c")*/

	/*fmt.Printf("list链表的长度：%d\n", list.length())
	fmt.Printf("list链表的数值")
	list.showList()
	fmt.Println()
	fmt.Println()*/

	// 链表反转
	//revers := list.reverseStack()
	//list1 := List{revers}

	// 链表分割
	/*partitionNode := list.partitionNode(5)
	list1 := List{partitionNode}
	list1.showList()*/

	/*
		fmt.Print("当前链表头部插入元素")
		list.add("before-add")
		fmt.Printf("list链表的数值")
		list.showList()
	*/

	/*
		合并两个有序链表

		list := List{}
		list.append(1)
		list.append(5)
		list.append(6)
		fmt.Printf("list1链表的数值")
		list.showList()

		fmt.Println()

		list2 := List{}
		list2.append(4)
		list2.append(9)
		list2.append(10)
		list2.append(11)

		fmt.Printf("list2链表的数值")
		list2.showList()

		fmt.Println()
		fmt.Printf("合并后list3链表的数值")

		 node := mergeTwoOrderList(list.headNode,list2.headNode)

		list3 := List{node}
		list3.showList()
	*/

	// 链表相邻节点交换

	/*list := List{}
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	list.append(6)

	fmt.Printf("交换前的链表")
	list.showList()
	fmt.Println()

	fmt.Printf("交换后的链表")

	node := swapPairs(list.headNode)
	list2 := List{node}
	list2.showList()*/

	/*list := List{}
	list.append(4)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(7)
	list.append(5)

	fmt.Printf("交换前的链表")
	list.showList()
	fmt.Println()

	fmt.Printf("交换后的链表")

	node := insertionSortList(list.headNode)
	list2 := List{node}
	list2.showList()*/

	/*list := List{}
	list.append(4)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(7)
	list.append(5)

	fmt.Printf("交换前的链表")
	list.showList()
	fmt.Println()

	fmt.Printf("交换后的链表")

	fmt.Println(list.headNode.Data)

	node := getKthFromEnd(list.headNode, 5)

	fmt.Println(node.Data)

	list2 := List{node}
	list2.showList()*/

	// 链表中间节点
	/*list := List{}
	list.append(4)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(7)
	list.append(5)

	fmt.Printf("初始化链表")
	list.showList()
	fmt.Println()

	fmt.Printf("中间节点：")
	node := middleNode(list.headNode)

	fmt.Println(node.Data)*/

	/*list := List{}
	list.append(4)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(7)
	list.append(5)

	fmt.Printf("初始化链表")
	list.showList()
	fmt.Println()

	fmt.Printf("奇偶节点：")
	node := oddEvenList(list.headNode)
	list2 := List{node}
	list2.showList()*/

	/*list := List{}
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	list.append(6)

	fmt.Printf("初始化链表")
	list.showList()
	fmt.Println()

	fmt.Printf(" 重排链表：")
	node := reorderList(list.headNode)
	list2 := List{node}
	list2.showList()*/

	/*// 回文链表
	list := List{}
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(3)
	list.append(2)
	list.append(1)

	fmt.Printf("初始化链表")
	list.showList()
	fmt.Println()

	fmt.Printf(" 重排链表：")
	ispalind := isPalindrome(list.headNode)
	fmt.Println(ispalind)*/

	/*// 环形链表
	list := List{}
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(3)
	list.append(2)
	list.append(1)

	fmt.Printf("初始化链表")
	list.showList()
	fmt.Println()

	fmt.Printf(" 常见环形连标：")
	cycleNode := createCycle(list.headNode, 1)
	hasCycle := hasCycle(cycleNode)
	fmt.Println(hasCycle)*/

	/*// 两个链表相加
	list1 := List{}
	list1.append(2)
	list1.append(4)
	list1.append(3)

	// 两个链表相加
	list2 := List{}
	list2.append(5)
	list2.append(6)
	list2.append(4)

	fmt.Printf("初始化链表1")
	list1.showList()
	fmt.Println()

	fmt.Printf("初始化链表1")
	list2.showList()
	fmt.Println()
	fmt.Printf(" 常见环形连标：")

	node := addTwoNumbers(list1.headNode, list2.headNode)
	fmt.Println()
	fmt.Printf(" 相加后链表：")

	list3 := List{node}
	list3.showList()*/

	// 环形链表入口点
	/*list := List{}
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	list.append(6)
	list.append(7)
	list.append(8)
	list.append(9)

	fmt.Printf("初始化链表")
	list.showList()
	fmt.Println()

	fmt.Printf(" 常见环形连标：")
	cycleNode := createCycle(list.headNode, 5)
	hasCycle := hasCycle(cycleNode)
	fmt.Println(hasCycle)

	detectNode := detectCycle(cycleNode)
	fmt.Printf(" 入口节点：")
	fmt.Print(detectNode.Data)*/

	//合并K有序链表

	list := List{}
	list.append(1)
	list.append(5)
	list.append(6)
	fmt.Printf("list1链表的数值")
	list.showList()

	fmt.Println()
	list1 := List{}
	list1.append(4)
	list1.append(9)
	list1.append(10)
	list1.append(11)

	fmt.Printf("list2链表的数值")
	list1.showList()
	fmt.Println()

	list3 := List{}
	list3.append(2)
	list3.append(7)
	list3.append(13)
	list3.append(20)

	fmt.Printf("list3链表的数值")
	list3.showList()
	fmt.Println()
	fmt.Printf("合并后链表：")

	arrList := []*Node{list.headNode, list1.headNode, list3.headNode}
	node := mergeKLists(arrList)
	list4 := List{node}
	list4.showList()

}
