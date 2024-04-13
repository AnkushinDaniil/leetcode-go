package lruCache

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	Cap         int
	Cache       map[int]*Node
	Left, Right *Node
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*Node, capacity),
		Left:  &Node{},
		Right: &Node{},
	}
	res.Left.Next = res.Right
	res.Right.Prev = res.Left
	return res
}

func (c *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) Insert(node *Node) {
	node.Prev = c.Right.Prev
	node.Next = c.Right
	c.Right.Prev.Next = node
	c.Right.Prev = node
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.Cache[key]; ok {
		c.Remove(node)
		c.Insert(node)
		return node.Val
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	node := &Node{
		Key: key,
		Val: value,
	}

	if _, ok := c.Cache[key]; ok {
		c.Remove(c.Cache[key])
	} else if len(c.Cache) == c.Cap {
		delete(c.Cache, c.Left.Next.Key)
		c.Remove(c.Left.Next)
	}

	c.Insert(node)
	c.Cache[key] = node
}
