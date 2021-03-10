package main

import "fmt"

var arr2 []int
var n2 int
var states2 []int
var j2 int

//我们有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4。
func init() {
	arr2 = []int{2, 9, 3, 6, 5, 1, 7}
	n2 = len(arr2)
	states2 = make([]int, n2 + 1)
	for i := 0; i <= n2; i++ {
		states2[i] = -1
	}
}

func main() {
	travel2(0, n2, 0)
	fmt.Println(states2)
	fmt.Println(j2)
}

func travel2(i int, lastIndex int, dist int) {
	if i == n2 {
		return
	}

	j2++

	//还未选择任何数 或者 当前决策的数满足递增
	if dist == 0 || arr2[i] > arr2[lastIndex] {
		if dist+1 > states2[i] {
			states2[i] = dist + 1
			travel2(i+1, i, dist+1)
		}
	}

	travel2(i+1, lastIndex, dist)
}
