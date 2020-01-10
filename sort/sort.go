package main

//插入排序：依次A[i]按照次序插入到已排序好的A[0]-A[i-1]数组中
func InsertSort(A []int) {
	length := len(A)
	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if A[i] < A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

//二分插入排序
func BinaryInsertSort(A []int) {
	length := len(A)
	for i := 1; i < length; i++ {
		index := binarySearch(A, A[i], 0, i-1)
		for j := index; j < i; j++ {
			A[i], A[j] = A[j], A[i]
		}
	}
}

//希尔排序：插入排序的改进算法
func ShellSort(A []int) []int {
	return nil
}

//选择排序：每次选择未排序数组中最小/最大的数字放到已排序好数组最边的位置
func SelectSort(A []int) {
	length := len(A)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

//冒泡排序：重复按顺序依次交换相邻位置的顺序，直至没有需要交换的时候
func BubbleSort(A []int) {
	length := len(A)
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
	}
}

//归并排序：将数组拆分成子数组，再对子数组合并为排序好的大数组
func MergeSort(A []int) []int {
	length := len(A)
	if length <= 1 {
		return A
	}

	middle := length / 2
	left := ShellSort(A[0:middle])
	len1 := len(left)
	right := ShellSort(A[middle:])
	len2 := len(right)
	arr := make([]int, length)
	for i,j,k := 0, 0, 0; k < length; k++ {
		if i == len1 {
			arr[k] = right[j]
			j++
		} else if j == len2 {
			arr[k] = left[i]
			i++
		} else {
			if left[i] <= right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
		}
	}
	return arr
}

//二分查找
func binarySearch(A []int, base, start, end int) int {
	if base < A[start] {
		return start
	} else if base > A[end] {
		return end + 1
	}
	return middleBinarySearch(A, base, start, end)
}

func middleBinarySearch(A []int, base, start, end int) int {
	if end == start {
		return start
	}

	m := (start + end) / 2
	if base < A[m] {
		if end == m {
			end--
		} else {
			end = m
		}
	} else if base > A[m] {
		if start == m {
			start++
		} else {
			start = m
		}
	} else {
		return m
	}
	return middleBinarySearch(A, base, start, end)
}
