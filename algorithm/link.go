package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {
	tree := &Node{
		Data: 1,
	}
	tree.Next = &Node{
		Data: 2,
	}
	tree.Next.Next = &Node{
		Data: 3,
	}
	tree2 := reverse(tree)
	walk(tree2)
}

func reverse(tree *Node) *Node {
	if tree == nil {
		return nil
	}
	tree2 := &Node{}
	for {
		tree2.Data = tree.Data
		if tree.Next == nil {
			break
		}
		tree2 = &Node{Next: tree2}
		tree = tree.Next
	}
	return tree2
}

func walk(tree *Node){
	for {
		fmt.Println(tree.Data)
		if tree.Next == nil {
			break
		}
		tree = tree.Next
	}
}
