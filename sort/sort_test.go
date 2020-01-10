package main

import (
	"testing"
	"reflect"
	"fmt"
)

var origin []int
var target []int

func init() {
	origin = []int{203, 1, 20, 50, 4, 28, 20, 405435, -2}
	target = []int{-2, 1, 4, 20, 20, 28, 50, 203, 405435}
}

func TestInsertSort(t *testing.T) {
	arr := copyArray(origin)
	InsertSort(arr)
	fmt.Println(origin, arr)
	if !reflect.DeepEqual(arr, target) {
		t.Error("insert sort error")
	}
}

func TestBinaryInsertSort(t *testing.T) {
	arr := copyArray(origin)
	BinaryInsertSort(arr)
	fmt.Println(origin, arr)
	if !reflect.DeepEqual(arr, target) {
		t.Error("insert sort error")
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 6, 20, 40}
	r := binarySearch(arr, 5, 0, 4)
	if r != 2 {
		t.Error("binary search error")
	}

	r = binarySearch(arr, 41, 0, 4)
	if r != 5 {
		t.Error("binary search error")
	}

	r = binarySearch(arr, -1, 0, 4)
	if r != 0 {
		t.Error("binary search error")
	}
}

//func TestShellSort(t *testing.T) {
//	arr := copyArray(origin)
//	arr1 := ShellSort(arr)
//	fmt.Println(origin, arr1)
//	if !reflect.DeepEqual(arr1, target) {
//		t.Error("shell sort error")
//	}
//}

func TestSelectSort(t *testing.T) {
	arr := copyArray(origin)
	SelectSort(arr)
	fmt.Println(origin, arr)
	if !reflect.DeepEqual(arr, target) {
		t.Error("select sort error")
	}
}

func TestBubbleSort(t *testing.T) {
	arr := copyArray(origin)
	BubbleSort(arr)
	fmt.Println(origin, arr)
	if !reflect.DeepEqual(arr, target) {
		t.Error("bubble sort error")
	}
}

func copyArray(arr []int) []int {
	newArr := make([]int, len(arr))
	for k, v := range arr {
		newArr[k] = v
	}
	return newArr
}
