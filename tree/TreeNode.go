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

/**
 5、题目描述：
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。

解决方案：
最直接的方法就是利用递归，遍历整棵树：如果当前节点不是叶子，对它的所有孩子节点，
递归调用 hasPathSum 函数，其中 sum 值减去当前节点的权值；
如果当前节点是叶子，检查 sum 值是否为 0，也就是是否找到了给定的目标和。
*/
func hasPathSum(root *Node, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Value
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

/**
	1、题目描述： 翻转二叉树
	2、解题方案： 使用递归
	反转一颗空树结果还是一颗空树。对于一颗根为 rr，左子树为 \mbox{right}，
	右子树为 \mbox{left} 的树来说，它的反转树是一颗根为 rr，
	左子树为 \mbox{right} 的反转树，右子树为 \mbox{left} 的反转树的树。
   3、算法分析
	既然树中的每个节点都只被访问一次，那么时间复杂度就是 O(n)，
	其中 n 是树中节点的个数。在反转之前，不论怎样我们至少都得访问每个节点至少一次，因此这个问题无法做地比 O(n)更好了。

*/
func invertTreeNode(root *Node) *Node {
	if root == nil {
		return nil
	}
	right := invertTreeNode(root.Right)
	left := invertTreeNode(root.Left)

	root.Right = left
	root.Left = right
	return root
}
