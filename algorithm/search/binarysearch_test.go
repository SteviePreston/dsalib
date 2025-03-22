package search

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11, 13}

	tests := []struct {
		target int
		want   int
	}{
		{7, 3},
		{1, 0},
		{13, 6},
		{2, -1},
	}

	for _, tt := range tests {
		got := BinarySearch(nums, tt.target)
		if got != tt.want {
			t.Errorf("BinarySearch(%d) = %d; want %d", tt.target, got, tt.want)
		}
	}
}
