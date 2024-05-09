package array

func BinarySearchTemplate1(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n - 1

    for left <= right {
        mid := left + (right-left)/2

		// It is not correct to compute the midpoint as (left+right)/2, because if the array is quite large (>2G)
		// then the result of `left+right` may overflow and become negative.
		// The proper way of choosing the midpoint pivot is `left + (right-left) / 2` which is mathematically equivalent but immune to overflow. 


        if nums[mid] == target {
            return mid // Target found, return its index
        } else if nums[mid] < target {
            left = mid + 1 // Move the left boundary to mid + 1
        } else {
            right = mid - 1 // Move the right boundary to mid - 1
        }
    }

    return -1 // Target not found
}