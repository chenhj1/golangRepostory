package main

import "fmt"

// https://leetcode.cn/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/
// 链表反转
/*
关键点：
1. 定义pre、curr
2. for循环的第一步 next = curr.next
*/

func main() {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	ret := reverse(node1)
	printNode(ret)
}

type Node struct {
	Value int
	Next  *Node
}

func reverse(root *Node) *Node {
	var pre *Node = nil
	curr := root

	for curr != nil {
		next := curr.Next
		curr.Next = pre

		pre = curr
		curr = next
	}
	return pre
}

func printNode(n *Node) {
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
}
