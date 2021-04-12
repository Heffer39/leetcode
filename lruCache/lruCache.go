package lruCache

type LRUCache struct {
	indexMap map[int]*Node
	nodeList LinkedList
	capacity int
}

type LinkedList struct {
	head, tail *Node
}

type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	imap := make(map[int]*Node, capacity)
	nodeList := LinkedList{}
	lru := LRUCache{indexMap : imap, capacity : capacity, nodeList : nodeList}
	return lru
}


func (this *LRUCache) Get(key int) int {
	if val, ok := this.indexMap[key]; ok {
		this.moveToFront(val)
		return val.value
	}
	return -1
}

func (this *LRUCache) moveToFront(val *Node) {
	if val.prev != nil {
		val.prev.next = val.next
	}
	if val.next != nil {
		val.next.prev = val.prev
	}
	if val.next != nil && this.nodeList.head.value == val.value {
		this.nodeList.head = val.next
	}
	this.nodeList.tail.next = val
	val.prev = this.nodeList.tail
	this.nodeList.tail = val
}

func (this *LRUCache) Put(key int, value int)  {
	if val, ok := this.indexMap[key]; ok {
		val.value = value
		this.moveToFront(val)
		return
	}
	// If at max capacity, need to evict LRU key/node
	if len(this.indexMap) == this.capacity {
		delete(this.indexMap, this.nodeList.head.key)
		if len(this.indexMap) > 1 {
			this.nodeList.head.next.prev = nil
		}
		this.nodeList.head = this.nodeList.head.next
	}
	newNode := &Node{key : key, value : value, prev : this.nodeList.tail}
	if len(this.indexMap) == 0 {
		this.nodeList.head = newNode
	} else {
		this.nodeList.tail.next = newNode
	}
	this.nodeList.tail = newNode
	this.indexMap[key] = newNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
