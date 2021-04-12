package lruCache

import (
	"fmt"
	"testing"
)

func TestLRUCache (t *testing.T) {
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	//fmt.Printf("Head: %v, Tail: %v\n", obj.nodeList.head.value, obj.nodeList.tail.value)
	fmt.Println(obj.Get(1))
	//fmt.Printf("Head: %v, Tail: %v\n", obj.nodeList.head.value, obj.nodeList.tail.value)
	obj.Put(3, 3)
	//fmt.Printf("Head: %v, Tail: %v\n", obj.nodeList.head.value, obj.nodeList.tail.value)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

func TestLRUCacheOverwriteValue (t *testing.T) {
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(1, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(3))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

func TestLRUCacheCapacity1 (t *testing.T) {
	obj := Constructor(1)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(3))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

func TestLRUCacheLC (t *testing.T) {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)

	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	//-1, 3
}

//["LRUCache","put","put","get","get","put","get","get","get"]
//[[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]
func TestLRUCacheLC2 (t *testing.T) {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(3, 2)
	fmt.Println(obj.Get(3))
	fmt.Printf("Head: %v, Tail: %v\n", obj.nodeList.head.value, obj.nodeList.tail.value)
	fmt.Println(obj.Get(2))
	fmt.Printf("Head: %v, Tail: %v\n", obj.nodeList.head.value, obj.nodeList.tail.value)
	obj.Put(4, 3)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}