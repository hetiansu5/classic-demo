package main

import (
	"errors"
	"fmt"
)

func faibo(n int) int {
	if n <= 2 {
		return n
	}
	return faibo(n-1) + faibo(n-2)
}

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

type arrayStack struct {
	arr   []int
	size  int
	count int
}

//数组顺序栈
func NewArrayStack(size int) {
	if size < 1 {
		panic(fmt.Sprintf("input size(%d) is invalid", size))
	}
	a := new(arrayStack)
	a.size = size
	a.arr = make([]int, size)
	a.count = 0
}

func (a *arrayStack) Push(data int) bool {
	if a.count >= a.size {
		return false
	}

	a.arr[a.count] = data
	a.count++
	return true
}

func (a *arrayStack) Pop() (int, error) {
	if a.count <= 0 {
		return 0, errors.New("stack is empty")
	}

	data := a.arr[a.count]
	a.count--
	return data, nil
}

func main() {
	i := faibo(4)
	fmt.Println(i)

	j := factorial(5)
	fmt.Println(j)

	tree := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
			},
		},
	}
	fmt.Println(tree)
	newTree := reverseLinkedList(tree)
	for {
		if newTree == nil {
			break
		}
		fmt.Println(newTree.Data)
		newTree = newTree.Next
	}

}

type Node struct {
	Data int
	Next *Node
}

func reverseLinkedList(tree *Node) *Node {
	var newTree *Node
	next := tree
	for {
		if next == nil {
			break
		}
		newTree = &Node{Data:next.Data, Next: newTree}
		next = next.Next
	}
	return newTree
}
