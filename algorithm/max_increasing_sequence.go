package main

import "fmt"

var arr1 []int
var n1 int
var states1 [][]int
var j1 int

//我们有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4。
func init() {
	arr1 = []int{2, 9, 3, 6, 5, 1, 7}
	n1 = len(arr1)
	states1 = make([][]int, n1)
	for i := 0; i < n1; i++ {
		states1[i] = make([]int, n1 + 1)
		for j := 0; j <= n1; j++ {
			states1[i][j] = -1
		}
	}
	fmt.Println(j1)
}

func main() {
	travel1(0, n1, 0)
	fmt.Println(states1)
	fmt.Println(j1)
}

func travel1(i int, lastIndex int, dist int) {
	if i == n1 {
		return
	}

	j1++

	//还未选择任何数 或者 当前决策的数满足递增
	if dist == 0 || arr1[i] > arr1[lastIndex] {
		if dist+1 > states1[i][i] {
			states1[i][i] = dist + 1
			travel1(i+1, i, dist+1)
		}
	}

	if dist > states1[i][lastIndex] {
		states1[i][lastIndex] = dist
		travel1(i+1, lastIndex, dist)
	}
}
