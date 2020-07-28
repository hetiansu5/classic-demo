package main

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type TreeRoot *Node

func newNode(d int) *Node {
	return &Node{
		Data: d,
	}
}

//递归前序遍历
func preOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	print(node.Data, " ")
	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)
}

//递归中序遍历
func inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left)
	print(node.Data, " ")
	inOrderTraversal(node.Right)
}

//递归后序遍历
func postOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	print(node.Data, " ")
}

//层级遍历
//原理：通过队列的先入先出原则，取出根节点，将左右子节点入队，从队列依次取出左右节点，打印并将左右子节点继续入队
func levelOrderTraversal(node *Node) {
	S := make([]*Node, 1)
	S[0] = node
	var T *Node
	for len(S) > 0 {
		T = S[0]
		S = S[1:]
		if T != nil {
			print(T.Data, " ")
			S = append(S, T.Left)
			S = append(S, T.Right)
		}
	}
}

//堆栈前序遍历
func preOrderTraversalByStack(node *Node) {
	S := make([]*Node, 0)
	T := node
	for T != nil || len(S) > 0 {
		for T != nil {
			print(T.Data, " ")
			S = append(S, T)
			T = T.Left
		}
		if len(S) > 0 {
			T = S[len(S)-1]
			S = S[0 : len(S)-1]
			T = T.Right
		}
	}
}

//堆栈中序遍历
func InOrderTraversalByStack(node *Node) {
	S := make([]*Node, 0)
	T := node
	for T != nil || len(S) > 0 {
		for T != nil {
			S = append(S, T)
			T = T.Left
		}
		if len(S) > 0 {
			T = S[len(S)-1]
			S = S[0 : len(S)-1]
			print(T.Data, " ")
			T = T.Right
		}
	}
}

//堆栈后序遍历
func postOrderTraversalByStack(node *Node) {
	S1 := make([]*Node, 0)
	S2 := make([]*Node, 0)
	T := node
	for T != nil || len(S1) > 0 {
		for T != nil {
			S2 = append(S2, T)
			S1 = append(S1, T)
			T = T.Right
		}
		if len(S1) > 0 {
			T = S1[len(S1)-1]
			S1 = S1[0 : len(S1)-1]
			T = T.Left
		}
	}
	for len(S2) > 0 {
		T = S2[len(S2)-1]
		S2 = S2[0 : len(S2)-1]
		print(T.Data, " ")
	}
}

func main() {
	treeHead := newDemoTree()

	println("递归前序遍历")
	preOrderTraversal(treeHead)
	println("")

	println("堆栈前序遍历")
	preOrderTraversalByStack(treeHead)
	println("")

	println("递归中序遍历")
	inOrderTraversal(treeHead)
	println("")

	println("层级遍历")
	levelOrderTraversal(treeHead)
	println("")

	println("堆栈中序遍历")
	InOrderTraversalByStack(treeHead)
	println("")

	println("递归后序遍历")
	postOrderTraversal(treeHead)
	println("")

	println("堆栈后序遍历")
	postOrderTraversalByStack(treeHead)
	println("")
}

/**
                   1
         2                   3
      4        5       6         7
	8   9   10  11   12  13   14   15
                             16 17

 */
func newDemoTree() TreeRoot {
	n1 := newNode(1)

	n2 := newNode(2)
	n1.Left = n2

	n3 := newNode(3)
	n1.Right = n3

	n4 := newNode(4)
	n2.Left = n4

	n5 := newNode(5)
	n2.Right = n5

	n6 := newNode(6)
	n3.Left = n6

	n7 := newNode(7)
	n3.Right = n7

	n8 := newNode(8)
	n4.Left = n8

	n9 := newNode(9)
	n4.Right = n9

	n10 := newNode(10)
	n5.Left = n10

	n11 := newNode(11)
	n5.Right = n11

	n12 := newNode(12)
	n6.Left = n12

	n13 := newNode(13)
	n6.Right = n13

	n14 := newNode(14)
	n7.Left = n14

	n15 := newNode(15)
	n7.Right = n15

	n16 := newNode(16)
	n14.Left = n16

	n17 := newNode(17)
	n14.Right = n17

	return n1
}
