package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
1、树的前序遍历
2、树的中序遍历
3、树的后序遍历
4、树的深度
5、判断两个树是否相同
6、重建二叉树
7、二叉树的层次遍历
8、二叉树的左视图和右视图
9、二叉树的锯齿形层序遍历
10、翻转二叉树
11、二叉树的最近公共祖先
12、N叉树层次遍历
13、二叉树的完整性校验（校验完全二叉树）
14、校验搜索二叉树
15、二叉树转双向链表
16、二叉树路径总和I
17、二叉树路径综和II
18、二叉树最大路径
19、二叉树最小的深度
20、二叉树的直径
*/

/*
	1、前序遍历 根-左-右
*/
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

/**
2、左=》根-》右
*/
func middleOrder(root *TreeNode) {
	if root == nil {
		return
	}
	middleOrder(root.Left)
	fmt.Printf("%d", root.Val)
	middleOrder(root.Right)
}

/**
后序遍历
3、左 -》右=》 根
*/
func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	middleOrder(root.Left)
	middleOrder(root.Right)
	fmt.Printf("%d", root.Val)
}

/**
4、树的最大深度
解题思路：对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1。层数(递归实现)
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

/**
5、判断两个二叉树是否相同
解题思路：递归
	两个二叉树相同，当且仅当两个二叉树的结构完全相同，且所有对应节点的值相同。因此，可以通过搜索的方式判断两个二叉树是否相同。
	如果两个二叉树都为空，则两个二叉树相同。如果两个二叉树中有且只有一个为空，则两个二叉树一定不相同。
	如果两个二叉树都不为空，那么首先判断它们的根节点的值是否相同，若不相同则两个二叉树一定不同，若相同，
	再分别判断两个二叉树的左子树是否相同以及右子树是否相同。这是一个递归的过程，因此可以使用深度优先搜索，递归地判断两个二叉树是否相同。
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/**
 6、重建二叉树
1.题目概述
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

2.解题思路
前序排列顺序为 根-左-右，中序排列为左-根-右。
那么如题根为1。
则根据中序遍历序列则可以得到左子树{4,7,2,}和右子树{5,3,8,6}。
又根据前序遍历则可以得到左子树的根为2，右子树的根为3。
重复3,4步。
直到左右子树皆为空时即可重建二叉树如图
https://upload-images.jianshu.io/upload_images/1441907-d31e4c6898f6c3ef.png?imageMogr2/auto-orient/strip|imageView2/2/w/366/format/webp
*/
func buildTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preOrder[0],
		Left:  nil,
		Right: nil,
	}
	i := 0
	for ; i < len(inOrder); i++ {
		if inOrder[i] == preOrder[0] {
			break
		}
	}
	root.Left = buildTree(preOrder[1:len(inOrder[:i])+1], inOrder[:i])
	root.Right = buildTree(preOrder[len(inOrder[:i])+1:], inOrder[i+1:])
	return root
}

/*

7、题目描述：二叉树的层次遍历
二叉树：[3,9,20,null,null,15,7],
返回其层序遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
解题思路：使用队列先进先出的特性，对二叉树进行层次遍历。
显然这道题是广度优先遍历的变种，只需要在广度优先遍历的过程中，
把每一层的节点都添加到同一个数组中即可，问题的关键在于遍历同一层节点前，
必须事先算出同一层的节点个数有多少(即队列已有元素个数)，因为 BFS 用的是队列来实现的
*/

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			first := queue[0]
			temp = append(temp, first.Val)
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			queue = queue[1:]
		}
		res = append(res, temp)
	}
	return res
}

/**
8、题目描述 【二叉树的左视图和右视图】
解题思想：
	1、层次遍历二叉树，
		左视图 取二叉树 每层遍历的第一个元素。
		右视图，取二叉树，每层遍历的第二个元素。
	2、使用队列来遍历进行二叉树的层次队列。先把跟节点添加到队列中，然后循环把每层的左右节点，分别加入队列中。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			first := queue[0]
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}

			/*if i == 0 { // 左视图
				res = append(res,first.Val)
			}*/
			if i == length-1 { // 右视图
				res = append(res, first.Val)
			}
			queue = queue[1:]
		}
	}
	return res
}

/**
9、题目描述	二叉树的锯齿形层序遍历
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
题目描述：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
  解题思路： 二叉树的层次遍历BFS，对奇数层进行翻转
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		res := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			first := queue[0]
			res = append(res, first.Val)
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			queue = queue[1:]
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(res); i < n/2; i++ {
				res[i], res[n-1-i] = res[n-1-i], res[i]
			}
		}
		ans = append(ans, res)
		level++
	}
	return ans
}

/**
10、翻转二叉树 （镜像二叉树）
	1、题目描述： 翻转二叉树 （镜像二叉树）
	2、解题方案： 使用递归
		https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
   3、算法分析
	既然树中的每个节点都只被访问一次，那么时间复杂度就是 O(n)，
	其中 n 是树中节点的个数。在反转之前，不论怎样我们至少都得访问每个节点至少一次，因此这个问题无法做地比 O(n)更好了。
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}

/**
	11、题目描述：【二叉树的最近公共祖先】
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
解题思路：
直接拿本身这个函数进行递归，本身这个函数的含义是在root这棵树找到p和q的最近公共祖先

1、若当前节点root == p，则表示q点一定在root的左右子树其中一处，则最近的公共结点肯定是root

2、若当前节点root == q，则表示p点一定在root的左右子树其中一处，则最近的公共结点肯定是root

3、若1和2情况都不是，则p和q的最近公共祖先要么在root的左子树，要么在root的右子树，则直接递归到root.left和root.right进行搜索，若递归完 后，左子树返回null表示没找到，那答案肯定是在右子树，同理，右子树返回null表示没找到，那答案肯定是在左子树

4、若3情况中左右子树都找不到p和q的最近公共祖先，则表示p点和q点分别在不同的左右子树，则root就是他们的最近公共祖先


题目连接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

解题思路：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2/
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

/**
12 题目描述：N叉树层次遍历
*/
func levelOrderN(root *TreeNode) [][]int {
	resut := make([][]int, 0)

	if root == nil {
		return resut
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		temp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			first := queue[0]
			for j := 0; j < len(first.Children); j++ {
				queue = append(queue, first.Children[j])
			}
			temp = append(temp, first.Val)
			queue = queue[1:]
		}
		resut = append(resut, temp)
	}
	return resut
}

/**
13、题目描述：二叉树的完整性校验
给定一个二叉树，确定它是否是一个完全二叉树。
百度百科中对完全二叉树的定义如下：
若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）

输入：[1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。


解题思路：
完全二叉树是指最后一层左边是满的，右边可能满也不能不满，然后其余层都是满的，根据这个特性，利用层遍历。如果我们当前遍历到了NULL结点，如果后续还有非NULL结点，说明是非完全二叉树。
参考链接：
	https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/solution/ceng-xu-bian-li-pan-duan-ji-ke-jie-jue-b-xcml/
*/
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	lastIsNil := false
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if lastIsNil == true && cur != nil {
			return false
		}

		if cur == nil {
			lastIsNil = true
			continue
		}
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return true
}

/**
	14、验证二叉搜索树
	给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
	有效 二叉搜索树定义如下：
		节点的左子树只包含 小于 当前节点的数。
		节点的右子树只包含 大于 当前节点的数。
		所有左子树和右子树自身必须也是二叉搜索树。

	输入：root = [2,1,3]
	输出：true

	解题思路：
		1、搜索二叉树的中序遍历是一个 升序的数组，可以借助这一个特性，来判断是不是 搜索二叉树。
 参考链接：https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/
*/

var preValue = math.MinInt32

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= preValue {
		return false
	}
	preValue = root.Val
	return isValidBST(root.Right)
}

/*
	题目描述15：二叉搜索树转成双向链表
	输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
	参考链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
*/

var pre *TreeNode
var head *TreeNode

// 生成双向链表
func treeToDoubleList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	pre = nil
	head = nil
	inOrder(root)

	pre.Right = head
	head.Left = pre

	return head
}

// 中序遍历
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)

	if pre == nil {
		head = root
	} else {
		pre.Right = root
		root.Left = pre
	}
	pre = root
	inOrder(root.Right)
}

/**
 16、题目描述：路径总和 I
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。

解决方案：
最直接的方法就是利用递归，遍历整棵树：如果当前节点不是叶子，对它的所有孩子节点，
递归调用 hasPathSum 函数，其中 sum 值减去当前节点的权值；
如果当前节点是叶子，检查 sum 值是否为 0，也就是是否找到了给定的目标和。

题目连接：https://leetcode-cn.com/problems/path-sum/submissions/

*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

/**
17、路径总和II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

参考解题思路：https://mp.weixin.qq.com/s/g5uvxi1lyxmWC4LtP0Bdlw

而在 Path Sum II 中，为了得到路径，我们需要从遍历的视角来看二叉树问题。

为了最终输出所有可能的路径，我们需要在遍历时记录当前路径，当发现路径满足条件时，就将路径保存下来。

路径遍历的顺序实际上就是 DFS 的顺序。当 DFS 进入一个结点时，路径中就增加一个结点；
当 DFS 从一个结点退出时，路径中就减少一个结点。下面的 GIF 动图展示了这个过程

回溯思想
回溯法采用试错的思想，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，
它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。—— 回溯法 - 维基百科[3]

*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	var path []int
	dfsPathSumII(root, targetSum, path, &res)
	return res
}

func dfsPathSumII(root *TreeNode, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			temp := make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
		}
	}

	targetSum = targetSum - root.Val
	dfsPathSumII(root.Left, targetSum, path, res)
	dfsPathSumII(root.Right, targetSum, path, res)
	path = path[:len(path)-1]
}

/**
18. 二叉树中的最大路径和
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
二叉树 abc，a 是根结点（递归中的 root），bc 是左右子结点（代表其递归后的最优解）。
最大的路径，可能的路径情况：
	a
 b		c

1、b + a + c。
2、b + a + a 的父结点。
3、a + c + a 的父结点。

其中情况 1，表示如果不联络父结点的情况，或本身是根结点的情况。
这种情况是没法递归的，但是结果有可能是全局最大路径和。
情况 2 和 3，递归时计算 a+b 和 a+c，选择一个更优的方案返回，也就是上面说的递归后的最优解啦。

另外结点有可能是负值，最大和肯定就要想办法舍弃负值（max(0, x)）（max(0,x)）。
但是上面 3 种情况，无论哪种，a 作为联络点，都不能够舍弃。

代码中使用 val 来记录全局最大路径和。
ret 是情况 2 和 3。
lmr 是情况 1。

所要做的就是递归，递归时记录好全局最大和，返回联络最大和。

参考链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/solution/er-cha-shu-zhong-de-zui-da-lu-jing-he-by-ikaruga/
*/

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain
		// 更新答案
		maxSum = max(maxSum, priceNewPath)
		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

/**
19、题目描述： 二叉树的最小深度
给定一个二叉树，找出其最小深度。最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
解题思路：基于层次遍历
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			temp := queue[0]
			if temp.Left == nil && temp.Right == nil {
				return depth
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
			queue = queue[1:]
		}
		depth++
	}
	return depth
}

/**
	20、二叉树的直径
	题目描述：给定一棵二叉树，你需要计算它的直径长度。
			一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 		  1
         / \
        2   3
       / \
      4   5
	返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

*/

/**
  思路：每个节点的最大直径路径 = 做孩子深度+右孩子深度
      但是 因为可以不通过跟节点，因此需要将每个节点最大直径(左子树深度+右子树深度)当前最大值比较并取大者

*/
func diameterOfBinaryTree(root *TreeNode) int {
	dialimeter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)
		dialimeter = max(dialimeter, leftDepth+rightDepth)
		return 1 + max(leftDepth, rightDepth)
	}
	dfs(root)
	return dialimeter
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
