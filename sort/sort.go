package sort

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left int, right int) {
	// if the indexes are out of order, then there's nothing to do.
	if left >= right {
		return
	}

	// partition these elements around the pivot, then sort each side.
	pivot := nums[(left+right)/2]
	index := partition(nums, left, right, pivot)
	quickSort(nums, left, index-1)
	quickSort(nums, index, right)
}

func partition(nums []int, left int, right int, pivot int) int {
	for left <= right {
		// move the left pointer until element that is bigger than the pivot, because that should be on the right side.
		for nums[left] < pivot {
			left++
		}

		// move the left pointer until element that is smaller than the pivot, because that should be on the left side.
		for nums[right] > pivot {
			right--
		}

		// swap these values
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	// return the partition point which will be at the left point.
	return left
}
