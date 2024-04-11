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
		Cache: make(map[int]*Node),
		Left:  &Node{},
		Right: &Node{},
	}
	res.Left.Next, res.Right.Prev = res.Right, res.Left
	return res
}

func (c *LRUCache) Remove(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

func (c *LRUCache) Insert(node *Node) {
	prev, next := c.Right.Prev, c.Right
	next.Prev = node
	prev.Next = next.Prev
	node.Next, node.Prev = next, prev
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
	if _, ok := c.Cache[key]; ok {
		c.Remove(c.Cache[key])
	}

	c.Cache[key] = &Node{
		Key: key,
		Val: value,
	}

	c.Insert(c.Cache[key])

	if len(c.Cache) > c.Cap {
		lru := c.Left.Next

		c.Remove(lru)
		delete(c.Cache, lru.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
