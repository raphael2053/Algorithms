package array

import "strconv"

// Iterating through a slice using a for loop
func Traverse(arr []int) string {
	var res string
	for _, val := range arr {
		res += strconv.Itoa(val)
	}
	return res
}

// Iterating through a slice using recursion
var res string

func TraverseRecursive(arr []int, i int) string {
	if i == len(arr) {
		return res
	}
	// Preorder position
	res += strconv.Itoa(arr[i])

	TraverseRecursive(arr, i+1)
	// Postorder position

	return res
}
