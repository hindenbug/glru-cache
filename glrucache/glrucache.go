package glrucache

import (
	"errors"
)

type Cache struct {
	capacity int
	Cache    map[interface{}]*Node
	CacheOperations
	Front *Node
	Rear  *Node
}

type Node struct {
	next, prev *Node
	Value      interface{}
	Key        interface{}
}

type CacheOperations interface {
	Get(key string) (interface{}, int)
	Set(key string, value interface{}) int
}

func NewCache(capacity int) (*Cache, error) {
	if capacity > 0 {
		return &Cache{
			capacity: capacity,
			Cache:    make(map[interface{}]*Node, capacity),
		}, nil
	}

	return nil, errors.New("Cache capacity cannot be zero")
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	if node, ok := c.Cache[key]; ok {
		c.removeNode(node)
		c.moveToFront(node)
		return node.Value, true
	}

	return nil, false
}

func (c *Cache) Set(key, value interface{}) bool {
	if node, exists := c.Cache[key]; exists {
		c.removeNode(node)
	}

	newNode := &Node{Key: key, Value: value, next: c.Front, prev: nil}

	c.moveToFront(newNode)

	// remove least recently used node i.e .rear
	if len(c.Cache) > c.capacity {
		c.removeNode(c.Rear)
	}

	return true
}

func (c *Cache) moveToFront(node *Node) {
	// push front node after given node
	if c.Front != nil {
		c.Front.prev = node
		node.prev = nil
		node.next = c.Front
	}

	// if give node is the only node
	if c.Rear == nil && c.Front == nil {
		c.Rear, c.Front = node, node
		node.prev = nil
		node.next = nil
	}

	c.Front = node
	c.Cache[node.Key] = node
}

func (c *Cache) removeNode(node *Node) {
	if node == nil {
		return
	}

	// if there is a node before given node
	if node.prev != nil {
		node.prev.next = node.next
	}

	// if there is a node after given node
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.Rear = node.prev
	}

	delete(c.Cache, node.Key)
}
