package utils

import "container/list"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	PurgeLast()
}

type lruCache struct {
	capacity int
	queue    *list.List
	items    map[Key]*list.Element
}

type Item struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		queue:    list.New(),
		items:    make(map[Key]*list.Element),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if element, ok := l.items[key]; ok {
		l.queue.MoveToFront(element)
		return true
	}
	if l.queue.Len() == l.capacity {
		l.PurgeLast()
	}
	i := &Item{
		key:   key,
		value: value,
	}
	element := l.queue.PushFront(l)
	l.items[i.key] = element

	return true
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := l.items[key]; ok {
		// тут просто возьмем  кусок из Set
		l.queue.MoveToFront(item)
		return item.Value.(*Item).value, true
	}
	return nil, false
}

func (l *lruCache) PurgeLast() {
	if element := l.queue.Back(); element != nil {
		item := l.queue.Remove(element).(*Item)
		delete(l.items, item.key)
	}
}
