package main

import (
	"strconv"
)

var arr1 []string
var arr2 []string

func init() {
	size1 := 100
	size2 := 1
	arr1 = make([]string, size1)
	arr2 = make([]string, size2)
	for i := 0; i < size1; i++ {
		arr1[i] = "" + strconv.Itoa(i)
	}

	for i := 0; i < size2; i++ {
		arr2[i] = "1" + strconv.Itoa(i)
	}

}

func main() {
	copySlice()
}

//在arr1数组大，arr2数组小的时候（100， 1） copy的性能更优
//BenchmarkCopySlice-4    	 9149899	       390 ns/op
//BenchmarkMergeSlice-4   	 5519101	       645 ns/op
func copySlice() []string {
	arr := make([]string, len(arr1)+len(arr2))
	copy(arr, arr1)
	copy(arr[len(arr1):], arr2)
	return arr
}

func mergeSlice() []string {
	arr := append(arr1, arr2...)
	return arr
}
