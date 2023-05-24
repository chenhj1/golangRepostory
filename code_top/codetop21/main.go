package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// https://leetcode.cn/problems/merge-two-sorted-lists/solution/he-bing-liang-ge-you-xu-lian-biao-by-leetcode-solu/
// 合并两个有序链表
/*
关键点：
1. 初始化preHead，永远当做链表头，这样可以最后返回preHead。
2. 定义节点pre，（或者叫curr，命名无所谓），作用是作为动态合并链表时，最新合并的节点。
*/

type Node struct {
	Value int
	Next  *Node
}

func mergeList(l1 *Node, l2 *Node) *Node {
	var preHead = &Node{Value: -1}
	var pre = preHead

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	// 有一个链表已经走到末尾，判断是l1还是l2。把未走到末尾的链表，续到pre之后就可以了
	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}

	return preHead
}

func Stringify(obj interface{}) string {
	ret, _ := jsoniter.MarshalToString(obj)
	return ret
}

func main() {
	n11 := &Node{Value: 1}
	n12 := &Node{Value: 3}
	n13 := &Node{Value: 5}
	n14 := &Node{Value: 7}
	n11.Next = n12
	n12.Next = n13
	n13.Next = n14

	n21 := &Node{Value: 2}
	n22 := &Node{Value: 4}
	n23 := &Node{Value: 6}
	n24 := &Node{Value: 8}
	n25 := &Node{Value: 10}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n25

	ret := mergeList(n11, n21)

	fmt.Println(Stringify(ret))
}
