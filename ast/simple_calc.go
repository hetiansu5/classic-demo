package main

import (
	"fmt"
)

const (
	TypeNum = iota
	TypeAdd
	TypeSub
	TypeMulti
	TypeDivide
)

type Node struct {
	Type  uint8
	Num   int
	Left  *Node
	Right *Node
}

type scanner struct {
	step func(*scanner, byte) int69
}

func main() {
	tree := &Node{
		Type: TypeAdd,
	}
	tree.Left = &Node{
		Type: TypeMulti,
	}
	tree.Left.Left = &Node{
		Num: 10,
	}
	tree.Left.Right = &Node{
		Type: TypeDivide,
	}
	tree.Left.Right.Left = &Node{
		Num: 12,
	}
	tree.Left.Right.Right = &Node{
		Num: 3,
	}
	tree.Right = &Node{
		Type: TypeNum,
		Num:  8,
	}

	fmt.Println(calc(tree))
}

func buildASTTree(data []byte) {

}

func calc(tree *Node) int {
	if tree == nil {
		return 0
	}
	switch tree.Type {
	case TypeAdd:
		return calc(tree.Left) + calc(tree.Right)
	case TypeSub:
		return calc(tree.Left) - calc(tree.Right)
	case TypeMulti:
		return calc(tree.Left) * calc(tree.Right)
	case TypeDivide:
		return calc(tree.Left) / calc(tree.Right)
	default:
		return tree.Num
	}
}
