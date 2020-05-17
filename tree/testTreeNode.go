package main

import "fmt"

func main() {
	root := Node{Value: 3}
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
	fmt.Print("\n层次遍历: ")

	layer := root.layer()
	fmt.Print("\n树的深度: ", layer)

}
