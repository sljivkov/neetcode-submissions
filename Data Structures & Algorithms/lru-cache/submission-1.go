type Node struct {
	key, val int

	prev *Node
	next *Node
}

type LRUCache struct {
	cache       map[int]*Node
	capacity    int
	left, right *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    make(map[int]*Node),
		capacity: capacity,
		left:     &Node{},
		right:    &Node{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

func (this *LRUCache) remove(node *Node) {
	next, prev := node.next, node.prev
	prev.next = node.next
	next.prev = node.prev
}

func (this *LRUCache) insert(node *Node) {
	next, prev := this.right, this.right.prev
	next.prev = node
	prev.next = node
	node.next = next
	node.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.remove(node)
		delete(this.cache, node.key)
	}
	node := &Node{
		key: key,
		val: value,
	}

	this.cache[key] = node

	this.insert(node)

	if this.capacity < len(this.cache) {

		temp := this.left.next

		this.remove(temp)
		delete(this.cache, temp.key)

	}
}
