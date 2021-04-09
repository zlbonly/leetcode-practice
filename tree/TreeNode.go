package main

import (
	"fmt"
	"math"
)

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

/**
124. 二叉树中的最大路径和
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

     1

2 		3
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

		-10

9				20


		15				7

输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

提示：
树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000



解题思路：
1、路径每走到一个节点，有3中选择：1：停在当前节点，2: 走到左子节点，3：走到右子节点
2、走到子节点，又面临3中选择，递归就是用来处理规模不一样的相同问题。
！！！ 不能走进一个分支又掉头回来走另一个分支，路径会重叠，不符合定义。


定义递归函数：
	1、对于一个节点而言，它只关心自己走入一个子树，能从中捞取的最大收益，不用管具体怎么走。
	2、定义dfs函数，返回当前子树向父节点 "提供"的最大路径和，即，一条从父节点延伸下来的路径，能在当前子树中获取到的最大收益
	分为三种情况：
		1、路径停在当前子树的跟节点，在这个子树中的收益 root.val
		2、走入左子树，在这个子树中的最大收益：root.val + dfs(root.left)
		3、走入右子树，在这个子树中的最大收益：root.val + dfs(root.right)


	对于当前父节点而言的三种选择，收益最大值：root.val + max(dfs(root.left),dfs(root.right))
再次提醒: 一条从父节点延伸下来的路径，不能走入左子树又掉头走右子树，不能两头收益，路径会重叠。

当遍历到null节点时，null 子树提供不了收益，返回 0。

如果某个子树 dfs 结果为负，走入它，收益不增反减，该子树应被忽略，杜绝选择走入它的可能，让它返回 0，像null一样如同砍掉。

子树中的内部路径要包含根节点
由题意可知，最大路径和可能产生于局部子树中，如下图左一。所以每递归一个子树，都求一下当前子树内部的最大路径和，见下图右一，从中比较出最大的。

注意: 一个子树内部的路径，要包含当前子树的根节点。如果不包含，那还算什么属于当前子树的路径，那就是当前子树的子树的内部路径了。

所以，一个子树内部的最大路径和 = 左子树提供的最大路径和 + 根节点值 + 右子树提供的最大路径和。即 dfs(root.left) + root.val + dfs(root.right)

时间复杂度 O(N)O(N)，每个节点都要遍历，空间复杂度是 O(H)O(H)，递归树的深度。

复盘总结
递归一个树，会对每个子树做同样的事（你写的处理逻辑）。
通过求出每个子树对外提供的最大路径和（return出来给父节点），从递归树底部向上，求出每个子树内部的最大路径和，后者是求解的目标，它的求解需要子树提供的值，理解清楚二者的关系。
每个子树的内部最大路径和，都挑战一下最大纪录，递归结束时，最大纪录就有了。
思考递归问题，不要纠结细节实现，结合求解的目标，自顶而下、屏蔽细节地思考。随着递归出栈，子问题自下而上地解决，最后解决了整个问题，内部细节是子递归帮你去做的。
你要做的只是写好递归的处理逻辑，怎么处理当前子树？需要返回东西吗？返回什么？再设置好递归的出口。其实就是——正确定义递归函数。

*/

func maxPathSum(root *Node) int {
	maxSum := math.MinInt32
	var dfs func(root *Node) int
	dfs = func(root *Node) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		innerMaxSum := left + root.Value + right

		maxSum = max(maxSum, innerMaxSum)

		outputMaxSum := root.Value + max(left, right)
		return max(outputMaxSum, 0)
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
