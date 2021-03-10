package main

import (
	"strconv"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decbin(43454))
	fmt.Println(plus(3454, 40000))
}

func decbin(a int) string {
	var arr []int
	for ; a > 0; {
		arr = append(arr, a%2)
		a = a >> 1
	}
	fmt.Println(arr)
	return join(reverse(arr))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func sum(a, b int) int {
	if b == 0 {
		return a
	}
	return sum(a^b, a&b<<1)
}

func plus(a, b int) string {
	binA := decbin(a)
	fmt.Println(binA)
	binB := decbin(b)
	la := len(binA)
	lb := len(binB)
	maxLen := max(la, lb)

	var ia, ib, carry, sum, nextCarry int
	var arr []int
	for i := 0; i < maxLen; i++ {
		if i < la {
			ia, _ = strconv.Atoi(binA[la-1-i : la-i])
		} else {
			ia = 0
		}
		if i < lb {
			ib, _ = strconv.Atoi(binB[lb-1-i : lb-i])
		} else {
			ib = 0
		}

		sum = ia ^ ib ^ carry
		carry = ia & ib | ((ia ^ ib) & carry)
		arr = append(arr, sum)

	}

	if nextCarry > 0 {
		arr = append(arr, nextCarry)
	}

	return join(reverse(arr))
}

func multi(a, b int) {

}

func reverse(arr []int) []int {
	l := len(arr)
	newArr := make([]int, l)
	for i := 0; i < l; i++ {
		newArr[i] = arr[l-1-i]
	}
	return newArr
}

func join(arr []int) string {
	l := len(arr)
	var build strings.Builder
	for i := 0; i < l; i++ {
		build.WriteString(strconv.Itoa(arr[i]))
	}
	return build.String()
}
