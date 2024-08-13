package bisectional_search

import "testing"

func TestBisectionalSearch(t *testing.T) {
	tests := []struct {
		name     string
		L        []int
		e        int
		expected bool
	}{
		{"Element Found - Middle", []int{1, 2, 3, 4, 5}, 3, true},
		{"Element Found - Last", []int{1, 2, 3, 4, 5, 6}, 6, true},
		{"Element Found - Single Element", []int{1}, 1, true},
		{"Element Not Found - Less than Min", []int{1, 2, 3, 4, 5}, 0, false},
		{"Element Not Found - Greater than Max", []int{1, 2, 3, 4, 5}, 6, false},
		{"Element Not Found - Empty List", []int{}, 1, false},
		{"Element Found - First Element", []int{2}, 2, true},
		{"Element Not Found - Single Element", []int{2}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bisectional_search(tt.L, tt.e)
			if result != tt.expected {
				t.Errorf("bisectional_search(%v, %d) = %v; want %v", tt.L, tt.e, result, tt.expected)
			}
		})
	}
}
