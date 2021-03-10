package main

import (
	"math"
	"fmt"
)

type Node struct {
	Data  int
	Index int
	Left  *Node
	Right *Node
}

func main() {
	tree := createTree()
	n := 5
	w := make([][]int, n)
	for i := range w {
		w[i] = make([]int, i+1)
		for j := range w[i] {
			w[i][j] = math.MaxInt32
		}
	}

	w[0][0] = tree.Data
	travel(1, tree.Index, w, tree.Left)
	travel(1, tree.Index, w, tree.Right)

	minV := math.MaxInt32
	for _, v := range w[n-1] {
		if minV > v {
			minV = v
		}
	}
	fmt.Println("Shortest Path Value:", minV)
}

func travel(i int, lastIndex int, w [][]int, node *Node) {
	if node == nil {
		return
	}

	t := w[i-1][lastIndex] + node.Data
	if t >= w[i][node.Index] {
		return
	}

	w[i][node.Index] = t
	travel(i+1, node.Index, w, node.Left)
	travel(i+1, node.Index, w, node.Right)
}

/**
 * 杨辉三角的二叉树
 */
func createTree() *Node {
	tree := &Node{Data: 5, Index: 0}

	n21 := &Node{Data: 7, Index: 0}
	n22 := &Node{Data: 8, Index: 1}
	tree.Left = n21
	tree.Right = n22

	n31 := &Node{Data: 2, Index: 0}
	n32 := &Node{Data: 3, Index: 1}
	n33 := &Node{Data: 4, Index: 2}
	n21.Left = n31
	n21.Right = n32
	n22.Left = n32
	n22.Right = n33

	n41 := &Node{Data: 6, Index: 0}
	n42 := &Node{Data: 9, Index: 1}
	n43 := &Node{Data: 6, Index: 2}
	n44 := &Node{Data: 1, Index: 3}
	n31.Left = n41
	n31.Right = n42
	n32.Left = n42
	n32.Right = n43
	n33.Left = n43
	n33.Right = n44

	n51 := &Node{Data: 2, Index: 0}
	n52 := &Node{Data: 7, Index: 1}
	n53 := &Node{Data: 9, Index: 2}
	n54 := &Node{Data: 4, Index: 3}
	n55 := &Node{Data: 5, Index: 4}
	n41.Left = n51
	n41.Right = n52
	n42.Left = n52
	n42.Right = n53
	n43.Left = n53
	n43.Right = n54
	n44.Left = n54
	n44.Right = n55

	return tree
}
