package main

import "fmt"

func main() {
	minstack := MinStack{}
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	min := minstack.GetMin()
	fmt.Print("最小值：")
	fmt.Println(min)
	minstack.Pop()
	top := minstack.Top()
	fmt.Print("栈顶元素：")
	fmt.Println(top)
	fmt.Print("最小值：")
	min = minstack.GetMin()
	fmt.Println(min)

}
