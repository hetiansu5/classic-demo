package main

import (
	"fmt"
)

func threeNumbers(nums []int) {
	size := len(nums)
	quickSort(nums, 0, size-1)
	zeroPos := binarySearch(nums, 0, size-1, 0)
	zeroLeft, zeroRight := zeroPos, zeroPos

	for ; zeroLeft > 0; zeroLeft-- {
		if nums[zeroLeft-1] != 0 {
			break
		}
	}

	for ; zeroRight < size-1; zeroRight++ {
		if nums[zeroRight+1] != 0 {
			break
		}
	}

	zeroCount := zeroRight - zeroLeft + 0


	
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	pivot := right
	pos := left
	for i := left; i <= right-1; i++ {
		if nums[i] < nums[pivot] {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
	nums[pos], nums[pivot] = nums[pivot], nums[pos]

	quickSort(nums, left, pos-1)
	quickSort(nums, pos+1, right)
}

func binarySearch(nums []int, left int, right int, num int) int {
	if left >= right {
		return -1
	}

	mid := left + (right-left)/2
	if num == nums[mid] {
		return mid
	} else if num < nums[mid] {
		return binarySearch(nums, left, mid, num)
	}
	return binarySearch(nums, mid+1, right, num)
}

func main() {
	nums := []int{1, 3, 4, 10, 27, 2, -1, 28, 0, 11}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
