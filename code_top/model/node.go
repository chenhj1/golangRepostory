package model

type Node struct {
	Value int
	Next  *Node
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}
