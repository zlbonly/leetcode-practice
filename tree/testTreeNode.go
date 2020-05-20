package main

import "fmt"

func main() {

	/*root := Node{Value: 3}
	root.Left = &Node{}
	root.Left.setValue(0)
	root.Left.Right = createNode(2)
	root.Right = &Node{5, nil, nil}
	root.Right.Left = createNode(4)

	fmt.Print("\n前序遍历: ")
	root.preOrder()
	fmt.Print("\n中序遍历: ")
	root.middleOrder()
	fmt.Print("\n后序遍历: ")
	root.preOrder()
	fmt.Print("\n层次遍历: ")*/

	/*layer := root.layer()
	fmt.Print("\n树的深度: ", layer)*/

	/**
	给定如下二叉树，以及目标和 sum = 22， 是否存在该路径
	*/

	/*root := Node{Value: 5}
	root.Left = createNode(4)
	root.Left.Left = createNode(11)
	root.Left.Left.Left = createNode(7)
	root.Left.Left.Right = createNode(2)

	root.Right = createNode(8)
	root.Right.Left = createNode(13)
	root.Right.Right = createNode(4)
	root.Right.Right.Right = createNode(1)
	hasPathSum := hasPathSum(&root, 22)
	if hasPathSum {
		fmt.Println("存在sum =22 的路径")
	} else {
		fmt.Println("不存在sum =22 的路径")
	}*/
	// 3、反转二叉树
	root := Node{Value: 5}
	root.Left = createNode(2)
	root.Left.Left = createNode(8)
	root.Left.Right = createNode(10)
	root.Right = createNode(3)
	root.Right.Left = createNode(4)
	root.Right.Right = createNode(5)
	fmt.Print("\n前序遍历: ")
	root.preOrder()
	invertRoot := invertTreeNode(&root)
	fmt.Println("\n翻转后前序遍历")
	invertRoot.preOrder()

}
