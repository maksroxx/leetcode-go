package lrucache

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToFront(node)
	} else {
		if len(this.cache) >= this.capacity {
			lru := this.tail.prev
			this.removeNode(lru)
			delete(this.cache, lru.key)
		}
		newNode := &Node{key: key, value: value}
		this.addToFront(newNode)
		this.cache[key] = newNode
	}
}

func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addToFront(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
