package array

import (
	"testing"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"5 elements, sorted", []int{1, 3, 5, 7, 9}, "13579"},
		{"5 elements, sorted", []int{9, 8, 7, 6, 5}, "98765"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Traverse(tt.nums)
			if res != tt.want {
				t.Errorf("nums = %v, Traverse() = %v, want %v", tt.nums, res, tt.want)
			}
		})
	}
}
func TestTraverseRecursive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"5 elements, sorted", []int{1, 3, 5, 7, 9}, "13579"},
		{"5 elements, sorted", []int{9, 8, 7, 6, 5}, "98765"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res = ""
			res := TraverseRecursive(tt.nums, 0)
			if res != tt.want {
				t.Errorf("nums = %v, Traverse() = %v, want %v", tt.nums, res, tt.want)
			}
		})
	}
}
