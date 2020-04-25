package main

import (
	"fmt"
)

func main() {
	list := List{}
	list.append(1)
	list.append(2)
	list.append(3)
	/*list.append("a")
	list.append("b")
	list.append("c")*/

	fmt.Printf("list链表的长度：%d\n", list.length())
	fmt.Printf("list链表的数值")
	list.showList()
	fmt.Println()
	fmt.Println()

	revers := list.reverseStack()

	list1 := List{revers}
	list1.showList()

	/*
		fmt.Print("当前链表头部插入元素")
		list.add("before-add")
		fmt.Printf("list链表的数值")
		list.showList()
	*/
}
