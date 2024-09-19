package main

import "container/list"

// 想要快，就抛弃掉List，自己设计链表和结点struct
// 目前的性能消耗在element.Value的断言上

type LRUCache struct {
	capacity int
	mapIndex map[int]*list.Element
	linklist *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		mapIndex: make(map[int]*list.Element, capacity),
		linklist: list.New(),
	}
}

type Pair struct {
	K, V int
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.mapIndex[key]; ok {
		this.linklist.MoveToFront(ele)
		return ele.Value.(Pair).V
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	//if this.linklist.Len() == this.capacity {
	//
	//}

	if ele, ok := this.mapIndex[key]; ok {
		//ele.Value.(Pair).V = value
		ele.Value = Pair{K: key, V: value}
		this.linklist.MoveToFront(ele)
	} else {

		this.mapIndex[key] = this.linklist.PushFront(
			Pair{K: key, V: value},
		)
	}
	if len(this.mapIndex) > this.capacity {
		delete(this.mapIndex, this.linklist.Back().Value.(Pair).K)
		this.linklist.Remove(this.linklist.Back())
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
