package main

// https://leetcode.cn/problems/lru-cache/solution/lruhuan-cun-ji-zhi-by-leetcode-solution/
// LRUCache
/*
关键点：
1. 双链表。优势：删除元素时，可以直接找到pre
2. 用map[int]*DNode 标记元素是否已经存在于LRUCache中
*/

func main() {

}

type DNode struct {
	Key   int
	Value int

	Pre  *DNode
	Next *DNode
}

type LRUCache struct {
	Head     *DNode
	Tail     *DNode
	Capacity int // 最大容量
	Size     int // 当前容量

	M map[int]*DNode // 检查某个值是否存在
}

func NewLRUCache(capacity int) *LRUCache {
	head := &DNode{
		Value: 0,
	}
	tail := &DNode{
		Value: 0,
	}
	head.Next = tail
	tail.Pre = head

	return &LRUCache{
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
		M:        map[int]*DNode{},
	}
}

func (c *LRUCache) Put(key int, val int) {
	if node, ok := c.M[key]; ok {
		c.moveToHead(node)
	} else {
		dNode := &DNode{
			Value: val,
			Pre:   nil,
			Next:  nil,
		}

		if c.Size >= c.Capacity {
			removedDNode := c.removeTail()
			delete(c.M, removedDNode.Key)

			c.addToHead(dNode)
		} else {
			c.addToHead(dNode)
			c.Size++
			c.M[key] = dNode
		}
	}

}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.M[key]; ok {
		c.moveToHead(node)
		return node.Value
	}
	return -1
}

func (c *LRUCache) removeNode(node *DNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (c *LRUCache) moveToHead(node *DNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) addToHead(node *DNode) {
	node.Pre = c.Head
	node.Next = c.Head.Next

	c.Head.Next = node
	c.Head.Next.Pre = node
}

func (c *LRUCache) removeTail() *DNode {
	node := c.Tail.Pre
	c.removeNode(node)
	return node
}
