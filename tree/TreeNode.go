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

func main() {
	grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}

	a := numIslands(grid)

	fmt.Printf("grid num =%v", a)
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
路径总和 I
 5、题目描述：
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。

解决方案：
最直接的方法就是利用递归，遍历整棵树：如果当前节点不是叶子，对它的所有孩子节点，
递归调用 hasPathSum 函数，其中 sum 值减去当前节点的权值；
如果当前节点是叶子，检查 sum 值是否为 0，也就是是否找到了给定的目标和。

题目连接：https://leetcode-cn.com/problems/path-sum/submissions/

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
路径总和II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

参考解题思路：https://mp.weixin.qq.com/s/g5uvxi1lyxmWC4LtP0Bdlw

回溯思想
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

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

解题连接：https://mp.weixin.qq.com/s/g5uvxi1lyxmWC4LtP0Bdlw
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	current := []int{}

	dfsSubsets(nums, 0, current, &res)
	return res
}

func dfsSubsets(nums []int, k int, current []int, res *[][]int) {
	if k == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, current)
		return
	}
	// 不选择第k个元素
	dfsSubsets(nums, k+1, current, res)
	// 选择第k个元素
	current = append(current, nums[k])
	dfsSubsets(nums, k+1, current, res)
	current = current[:len(current)-1]
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
题目描述：
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],
   3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]

解题思路：
显然这道题是广度优先遍历的变种，只需要在广度优先遍历的过程中，
把每一层的节点都添加到同一个数组中即可，问题的关键在于遍历同一层节点前，
必须事先算出同一层的节点个数有多少(即队列已有元素个数)，因为 BFS 用的是队列来实现的

*/

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for level := 0; 0 < len(queue); level++ {
		res = append(res, []int{})
		next := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			res[level] = append(res[level], node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		queue = next
	}
	return res
}

//二叉树遍历例题总结：

/**
	1、DFS
	DFS  全称Depth First Search 中文名为深度优先搜索。是一种以深度方向搜索某种数据结构的方法，常用栈来辅助DFS算法。深度优先搜索
	大多要以递归方式实现，所以要考虑递归爆栈的可能性。

	void traverse(TreeNode root) {
    // 判断 base case
    if (root == null) {
        return;
    }
    // 访问两个相邻结点：左子结点、右子结点
    traverse(root.left);
    traverse(root.right);
}

1、第一个要素是访问相邻结点。二叉树的相邻结点非常简单，只有左子结点和右子结点两个。二叉树本身就是一个递归定义的结构：
一棵二叉树，它的左子树和右子树也是一棵二叉树。那么我们的 DFS 遍历只需要递归调用左子树和右子树即可。

2、第二个要素是 判断 base case。一般来说，二叉树遍历的 base case 是 root == null。
这样一个条件判断其实有两个含义：一方面，这表示 root 指向的子树为空，不需要再往下遍历了。另一方面，
在 root == null 的时候及时返回，可以让后面的 root.left 和 root.right 操作不会出现空指针异常。


	2、BFS
	BFS 全称 Breadth First Search 中文名为广度优先搜索。是一种以宽度方向搜索某种数据结构的方法 常用队列辅助BFS算法。

*/

/**
 3、1 网格类问题的DFS遍历方法：
1、网格问题的基本概念
我们首先明确一下岛屿问题中的网格结构是如何定义的，以方便我们后面的讨论。

网格问题是由  个小方格组成一个网格，每个小方格与其上下左右四个方格认为是相邻的，要在这样的网格上进行某种搜索。

岛屿问题是一类典型的网格问题。每个格子中的数字可能是 0 或者 1。我们把数字为 0 的格子看成海洋格子，数字为 1 的格子看成陆地格子，
这样相邻的陆地格子就连接成一个岛屿。

2、网格类 DFS 的基本结构
网格结构要比二叉树结构稍微复杂一些，它其实是一种简化版的图结构。要写好网格上的 DFS 遍历，我们首先要理解二叉树上的 DFS 遍历方法，再类比写出网格结构上的 DFS 遍历。
我们写的二叉树 DFS 遍历一般是这样的：
 2.1 void traverse(TreeNode root) {
    // 判断 base case
    if (root == null) {
        return;
    }
    // 访问两个相邻结点：左子结点、右子结点
    traverse(root.left);
    traverse(root.right);
}

可以看到，二叉树的 DFS 有两个要素：「访问相邻结点」和「判断 base case」。

 1)第一个要素是访问相邻结点。二叉树的相邻结点非常简单，只有左子结点和右子结点两个。二叉树本身就是一个递归定义的结构：一棵二叉树，它的左子树和右子树也是一棵二叉树。那么我们的 DFS 遍历只需要递归调用左子树和右子树即可。

2) 第二个要素是 判断 base case。一般来说，二叉树遍历的 base case 是 root == null。这样一个条件判断其实有两个含义：一方面，这表示 root 指向的子树为空，不需要再往下遍历了。另一方面，在 root == null 的时候及时返回，可以让后面的 root.left 和 root.right 操作不会出现空指针异常。

对于网格上的 DFS，我们完全可以参考二叉树的 DFS，写出网格 DFS 的两个要素：

首先，网格结构中的格子有多少相邻结点？答案是上下左右四个。对于格子 (r, c) 来说（r 和 c 分别代表行坐标和列坐标），四个相邻的格子分别是 (r-1, c)、(r+1, c)、(r, c-1)、(r, c+1)。换句话说，网格结构是「四叉」的。

其次，网格 DFS 中的 base case 是什么？从二叉树的 base case 对应过来，应该是网格中不需要继续遍历、grid[r][c] 会出现数组下标越界异常的格子，也就是那些超出网格范围的格子。

3、这样，我们得到了网格 DFS 遍历的框架代码：
void dfs(int[][] grid, int r, int c) {
    // 判断 base case
    // 如果坐标 (r, c) 超出了网格范围，直接返回
    if (!inArea(grid, r, c)) {
        return;
    }
    // 访问上、下、左、右四个相邻结点
    dfs(grid, r - 1, c);
    dfs(grid, r + 1, c);
    dfs(grid, r, c - 1);
    dfs(grid, r, c + 1);
}

// 判断坐标 (r, c) 是否在网格中
boolean inArea(int[][] grid, int r, int c) {
    return 0 <= r && r < grid.length
         && 0 <= c && c < grid[0].length;
}

4、如何避免重复遍历
网格结构的 DFS 与二叉树的 DFS 最大的不同之处在于，遍历中可能遇到遍历过的结点。这是因为，网格结构本质上是一个「图」，我们可以把每个格子看成图中的结点，每个结点有向上下左右的四条边。在图中遍历时，自然可能遇到重复遍历结点。

这时候，DFS 可能会不停地「兜圈子」，永远停不下来，如下图所示
https://mmbiz.qpic.cn/mmbiz_gif/TKAD4axFcib8CSJjOnbGUamCj2B7OiclOvhNwXiaoJdXnDNUGpmtqqlHIbPnpejyAVqnQqODmYMIxovmlLcn0xEicA/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1
如何避免这样的重复遍历呢？答案是标记已经遍历过的格子。以岛屿问题为例，我们需要在所有值为 1 的陆地格子上做 DFS 遍历。每走过一个陆地格子，就把格子的值改为 2，这样当我们遇到 2 的时候，就知道这是遍历过的格子了。也就是说，每个格子可能取三个值：

0 —— 海洋格子
1 —— 陆地格子（未遍历过）
2 —— 陆地格子（已遍历过）
我们在框架代码中加入避免重复遍历的语句：
void dfs(int[][] grid, int r, int c) {
    // 判断 base case
    if (!inArea(grid, r, c)) {
        return;
    }
    // 如果这个格子不是岛屿，直接返回
    if (grid[r][c] != 1) {
        return;
    }
    grid[r][c] = 2; // 将格子标记为「已遍历过」

    // 访问上、下、左、右四个相邻结点
    dfs(grid, r - 1, c);
    dfs(grid, r + 1, c);
    dfs(grid, r, c - 1);
    dfs(grid, r, c + 1);
}

// 判断坐标 (r, c) 是否在网格中
boolean inArea(int[][] grid, int r, int c) {
    return 0 <= r && r < grid.length
         && 0 <= c && c < grid[0].length;
}
*/

/**
3.1 网格DFS例题1
	岛屿的最大面积:
	给定一个包含了一些 0 和 1 的非空二维数组 grid，一个岛屿是一组相邻的 1（代表陆地），这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
	你可以假设 grid 的四个边缘都被 0（代表海洋）包围着。
	找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

题目连接：https://leetcode-cn.com/problems/max-area-of-island/
maxAreaOfIsland（）
*/

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				a := area(grid, r, c)
				res = Max(res, a)
			}
		}
	}
	return res
}

/**
463. 岛屿的周长
	给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。

网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

题目地址：
https://leetcode-cn.com/problems/island-perimeter/

*/
func islandPerimeter(grid [][]int) int {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				// 题目限制只有一个岛屿，计算一个即可
				return dfsPerimeter(grid, r, c)
			}
		}
	}
	return 0
}

func dfsPerimeter(grid [][]int, r int, c int) int {
	// 函数因为坐标[r,c] 超出网格范围，返回对应一条黄色的边
	if !inArea(grid, r, c) {
		return 1
	}
	// 函数因为 当前格子是海洋格子 返回 对应一条蓝色的边
	if grid[r][c] == 0 {
		return 1
	}
	// 函数因为「当前格子是已遍历的陆地格子」返回，和周长没关系
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2

	return dfsPerimeter(grid, r-1, c) + dfsPerimeter(grid, r+1, c) + dfsPerimeter(grid, r, c-1) + dfsPerimeter(grid, r, c+1)
}

/**

	岛屿数量：
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
题目连接：https://leetcode-cn.com/problems/number-of-islands/
*/
func numIslands(grid [][]int) int {
	islandNum := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				dfsIslandNum(grid, r, c)
				islandNum++
			}
		}
	}
	return islandNum
}

func dfsIslandNum(grid [][]int, r int, c int) {
	// 1、判断是否在网格范围内
	if !inArea(grid, r, c) {
		return
	}
	// 2、判断是否重复遍历
	if grid[r][c] != 1 {
		return
	}
	// 3、标记已遍历
	grid[r][c] = 2
	// 4、遍历上下左右网格
	dfsIslandNum(grid, r-1, c)
	dfsIslandNum(grid, r+1, c)
	dfsIslandNum(grid, r, c-1)
	dfsIslandNum(grid, r, c+1)
}

// 获取最大值
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func area(grid [][]int, r int, c int) int {
	// 1、判断是否在网格范围内
	if !inArea(grid, r, c) {
		return 0
	}
	// 2、判断是否重复遍历
	if grid[r][c] != 1 {
		return 0
	}
	// 3、标记已遍历
	grid[r][c] = 2
	// 4、遍历上下左右网格
	return 1 + area(grid, r-1, c) + area(grid, r+1, c) + area(grid, r, c-1) + area(grid, r, c+1)
}

func inArea(grid [][]int, r int, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}

/**
	1、二叉树的最近公共祖先
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

题目连接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

解题思路：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
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
