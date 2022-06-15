package lru

import (
	"container/list"
)

type Cache struct {
	list *list.List
	keys map[int]*list.Element
	cap  int
}

type node struct {
	Key   int
	Value int
}

func NewCache(capacity int) *Cache {

	return &Cache{
		list: list.New(),
		keys: make(map[int]*list.Element),
		cap:  capacity,
	}
}

func (l *Cache) Get(key int) int {
	ele, ok := l.keys[key]
	if !ok {
		return -1
	}
	v := ele.Value.(*node).Value
	l.list.Remove(ele)
	ele = l.list.PushBack(&node{
		Key: key, Value: v,
	})
	l.keys[key] = ele
	return v
}

func (l *Cache) Put(key int, value int) {
	ele, ok := l.keys[key]
	// 只要存在，无论是否满了，都插入最后
	if ok {
		l.list.Remove(ele)
		ele = l.list.PushBack(&node{
			Key:   key,
			Value: value,
		})
		l.keys[key] = ele
		return
	}

	if len(l.keys) < l.cap {
		ele = l.list.PushBack(&node{
			Key:   key,
			Value: value,
		})
		l.keys[key] = ele
		return
	}

	ele = l.list.Front()
	k := ele.Value.(*node).Key
	l.list.Remove(ele)
	delete(l.keys, k)
	ele = l.list.PushBack(&node{
		Key:   key,
		Value: value,
	})
	l.keys[key] = ele
}
