package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-level-order-traversal/solution/er-cha-shu-de-ceng-xu-bian-li-by-leetcode-solution/
// 二叉树遍历
/*
关键点：
1. golang与java语法特性有区别，go没有队列，因此，go采用数组来替代队列。
2. java中，本层与下层节点，复用同一个队列，但是用length表示遍历到何处是一层。而go语言，可以直接定义"本层"、"下层"两个数组，直接隔离开了。
3. go先读取本层数组处理本层，顺便把孩子节点存储到下层数组。然后，再跳到下层数组，继续这个过程。
*/

func main() {
	root := &TreeNode{Value: 1}
	n2 := &TreeNode{Value: 2}
	n3 := &TreeNode{Value: 3}
	n4 := &TreeNode{Value: 4}
	n5 := &TreeNode{Value: 5}

	root.Left = n2
	root.Right = n3

	n2.Left = n4
	n2.Right = n5

	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	currLevelNodes := []*TreeNode{root} // 只存本层的节点。

	for i := 0; len(currLevelNodes) > 0; i++ { // 遍历多少次层。i的取值是有多少层

		ret = append(ret, []int{})
		nextLevelNodes := []*TreeNode{} // 用作预存储下一层

		for j := 0; j < len(currLevelNodes); j++ {

			// 处理本层
			node := currLevelNodes[j]
			ret[i] = append(ret[i], node.Value)

			// 把本层的左右孩子节点，放到下一层的列表。
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		currLevelNodes = nextLevelNodes // 已经处理完了本层，于是跳到下一层。
	}
	return ret
}
