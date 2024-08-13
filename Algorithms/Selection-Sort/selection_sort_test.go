package selection_sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Already sorted list", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Unsorted list", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"List with duplicates", []int{3, 1, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
		{"List with single element", []int{1}, []int{1}},
		{"Empty list", []int{}, []int{}},
		{"List with negative numbers", []int{-2, -3, -1, -4, -5}, []int{-5, -4, -3, -2, -1}},
		{"List with all elements the same", []int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := selection_sort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("selection_sort(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
