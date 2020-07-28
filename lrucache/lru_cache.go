package main

import "fmt"

type node struct {
	Data     string
	prev     *node
	next     *node
	hashNext *node
}

type lruCache struct {
	head      *node
	hashTable []*node
	capacity  int
}

func New() *lruCache {
	l := new(lruCache)
	l.capacity = 1 << 4;
	l.hashTable = make([]*node, l.capacity)
	return l
}

func (l *lruCache) Add(url string) {
	hashValue := l.hash(url)
	if l.hashTable[hashValue] == nil {
		newNode := newNode(url)
		if l.head == nil {
			l.head = newNode
		} else {
			l.hashTable[hashValue] = newNode
			l.appendToHead(newNode)
		}
		return
	}

	node := l.hashTable[hashValue]

	for {
		if node.Data == url {
			if l.head == node {
				break
			}
			l.cutNode(node)
			l.appendToHead(node)
			break
		}

		if node.hashNext == nil {
			newNode := newNode(url)
			node.hashNext = newNode
			l.appendToHead(newNode)
			break
		}

		node = node.hashNext
	}
}

func (l *lruCache) hash(url string) int {
	arr := []rune(url)
	a := 0
	for _, v := range arr {
		a = (a + int(v)) % l.capacity
	}
	return a
}

func (l *lruCache) traversalLinkedList() {
	slot := l.head
	for slot != nil {
		fmt.Println(slot)
		slot = slot.next
	}
}

func (l *lruCache) cutNode(node *node) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

func (l *lruCache) appendToHead(node *node) {
	node.prev = nil
	l.head.prev = node
	l.head, node.next = node, l.head
}

func newNode(url string) *node {
	node := new(node)
	node.Data = url
	return node
}

func main() {
	lru := New()
	lru.Add("http://a.com/10")
	lru.Add("http://a.com/14")
	lru.Add("http://a.com/21")
	lru.Add("http://a.com/25")
	lru.Add("http://a.com/14")
	lru.Add("http://a.com/28")
	lru.Add("http://a.com/25")
	lru.Add("http://a.com/105")
	lru.Add("http://a.com/106")
	lru.Add("http://a.com/107")
	lru.traversalLinkedList()
}
