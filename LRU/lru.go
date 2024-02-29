package LRU

import (
	"container/list"
	"sync"
	"time"
)

// Node represents a key-value pair with expiration time
type Node struct {
	key       interface{}
	value     interface{}
	expiresAt time.Time
}

// LRUCache implements a least recently used cache with expiration
type LRUCache struct {
	m       sync.Mutex
	capacity int
	cache    map[interface{}]*list.Element
	queue    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*list.Element),
		queue:    list.New(),
	}
}


func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	if element, ok := c.cache[key]; ok {
		if c.isExpired(element) {
			c.removeElement(element)
			delete(c.cache, key)
			return nil, false
		}
		c.moveToFront(element)
		return element.Value.(*Node).value, true
	}
	return nil, false
}


func (c *LRUCache) Put(key, value interface{}, expiration time.Duration) {
	c.m.Lock()
	defer c.m.Unlock()

	if element, ok := c.cache[key]; ok {
		element.Value.(*Node).value = value
		element.Value.(*Node).expiresAt = time.Now().Add(expiration)
		c.moveToFront(element)
		return
	}

	if c.queue.Len() == c.capacity {
		c.removeOldest()
	}

	newNode := &Node{
		key:       key,
		value:     value,
		expiresAt: time.Now().Add(expiration),
	}
	newElement := c.queue.PushBack(newNode)
	c.cache[key] = newElement
}

func (c *LRUCache) removeOldest() {
	element := c.queue.Front()
	if element == nil {
		return
	}
	c.removeElement(element)
	delete(c.cache, element.Value.(*Node).key)
}


func (c *LRUCache) removeElement(element *list.Element) {
	c.queue.Remove(element)
}

func (c *LRUCache) moveToFront(element *list.Element) {
	if element.Next() == nil {
		return
	}
	c.queue.Remove(element)
	c.queue.PushFront(element.Value)
}


func (c *LRUCache) isExpired(element *list.Element) bool {
	return element.Value.(*Node).expiresAt.Before(time.Now())
}