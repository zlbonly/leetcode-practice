package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func createNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) setValue(val int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Value = val
}

/*
	前序遍历 根-左-右
*/

func (node *Node) preOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.preOrder()
	node.Right.preOrder()
}

/**
左=》根-》右
*/
func (node *Node) middleOrder() {
	if node == nil {
		return
	}
	node.Left.middleOrder()
	node.Print()
	node.Right.middleOrder()
}

/**
后序遍历
左 -》右=》 根
*/
func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.Left.postOrder()
	node.Right.postOrder()
	node.Print()
}

/**
树 的深度
//层数(递归实现)
//对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
*/
func (node *Node) layer() int {
	if node == nil {
		return 0
	}
	leftLayer := node.Left.layer()
	rightLayer := node.Right.layer()

	if leftLayer > rightLayer {
		return leftLayer + 1
	} else {
		return rightLayer + 1
	}

}
