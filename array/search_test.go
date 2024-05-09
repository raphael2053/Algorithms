package array

import (
	"reflect"
	"testing"
)

func TestBinarySearchTemplate1(t *testing.T) {
	tests := []struct{
		name string
		nums []int
		target int
		want int
	}{
		{"one element", []int{1}, 1, 0},
		{"two elements, sorted", []int{1,2}, 2, 1},
		{"multiple elements, sorted", []int{1,3,5,7,9}, 7, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := BinarySearchTemplate1(tt.nums, tt.target)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("nums = %v, target = %v, BinarySearchTemplate1() = %v, want %v", tt.nums, tt.target, res, tt.want)
			}
		})
	}
}