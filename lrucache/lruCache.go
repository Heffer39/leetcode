// Package lrucache represents a data structure that follows the constraints of a Least Recently Used (LRU) Cache
package lrucache

import "container/list"

type element = list.Element

// LRUCache keeps track of key, value pairs, as well as the order in which they were most recently accessed
type LRUCache struct {
	indexMap map[int]*element
	nodeList list.List
	capacity int
}

// Node represents a key, value pair for each element in a nodeList
type Node struct {
	key, value int
}

// Constructor instantiates an LRUCache
func Constructor(capacity int) LRUCache {
	imap := make(map[int]*element, capacity)
	nodeList := list.List{}
	lru := LRUCache{indexMap: imap, capacity: capacity, nodeList: nodeList}
	return lru
}

// Get returns the value associated with a given key
// This method accesses a value in the cache, which makes it the most recently used value,
// moving it to the front of the list
func (cache *LRUCache) Get(key int) int {
	if val, ok := cache.indexMap[key]; ok {
		cache.nodeList.MoveToBack(val)
		return val.Value.(Node).value
	}
	return -1
}

// Put adds a new value to the LRU
// This method adds a new node to the end of the LRU list,
// as well as stores the association from the key to an element in the LRU map
func (cache *LRUCache) Put(key int, value int) {
	// If key already exists, remove and add new value
	if val, ok := cache.indexMap[key]; ok {
		cache.nodeList.Remove(val)
	}
	newNode := Node{key: key, value: value}
	cache.indexMap[key] = cache.nodeList.PushBack(newNode)
	// If at max capacity, need to evict LRU key/node
	if len(cache.indexMap) > cache.capacity {
		delete(cache.indexMap, cache.nodeList.Front().Value.(Node).key)
		cache.nodeList.Remove(cache.nodeList.Front())
	}
}
