package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"one element", []int{1}, []int{1}},
		{"two elements, sorted", []int{1, 2}, []int{1, 2}},
		{"two elements, unsorted", []int{2, 1}, []int{1, 2}},
		{"multiple elements, sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"multiple elements, unsorted", []int{4, 2, 3, 1}, []int{1, 2, 3, 4}},
		{"all elements the same", []int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
