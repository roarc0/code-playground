package main

import (
	"reflect"
	"testing"
)

func removeDuplicates(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		want     int
		wantNums []int
	}{
		{
			name:     "a",
			nums:     []int{1, 1, 2},
			want:     2,
			wantNums: []int{1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.wantNums, tt.nums) {
				t.Errorf("removeDuplicates() nums = %v, wantNums %v", tt.nums, tt.wantNums)
			}
		})
	}
}
