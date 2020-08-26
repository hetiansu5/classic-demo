package main

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

//二分插入排序
func BinaryInsertSort(A []int) {
	length := len(A)
	for i := 1; i < length; i++ {
		index := BinarySearch(A, A[i])
		for j := index; j < i; j++ {
			A[i], A[j] = A[j], A[i]
		}
	}
}

//希尔排序：插入排序的改进算法
func ShellSort(A []int) []int {
	return nil
}

//二分查b找
func BinarySearch(A []int, value int) int {
	return binarySearchInternally(A, 0, len(A)-1, value)
}

func binarySearchInternally(A []int, start, end, value int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if A[mid] == value {
		return mid
	} else if A[mid] < value {
		return binarySearchInternally(A, start, mid-1, value)
	} else {
		return binarySearchInternally(A, mid+1, end, value)
	}
}


//归并排序：将数组拆分成子数组，再对子数组合并为排序好的大数组
func MergeSort(A []int) {
	mergeSortRange(A, 0, len(A)-1)
}

func mergeSortRange(A []int, p, q int) {
	if p >= q {
		return
	}
	var mid int
	mid = (p + q) / 2
	mergeSortRange(A, p, mid)
	mergeSortRange(A, mid+1, q)
	mergeArray(A[p:q+1], A[p:mid+1], A[mid+1:q+1])
}

func mergeArray(A []int, A1 []int, A2 []int) {
	i := 0
	j := 0
	l := len(A1) + len(A2)
	B := make([]int, l)
	for ; i+j < l; {
		if i >= len(A1) {
			B[i+j] = A2[j]
			j++
			continue
		} else if j >= len(A2) {
			B[i+j] = A1[i]
			i++
			continue
		}

		if A1[i] < A2[j] {
			B[i+j] = A1[i]
			i++
		} else {
			B[i+j] = A2[j]
			j++
		}
	}
	for k, v := range B {
		A[k] = v
	}
}

//快速排序
func QuickSort(A []int) {
	quickSortRange(A, 0, len(A)-1);
}

func quickSortRange(A []int, p int, q int) {
	if p >= q {
		return
	}

	pivot := q;
	j := p;
	for i := p; i <= q-1; i++ {
		if A[i] < A[pivot] {
			A[i], A[j] = A[j], A[i];
			j++
		}
	}
	A[j], A[pivot] = A[pivot], A[j]
	quickSortRange(A, p, j-1)
	quickSortRange(A, j+1, q)
}
