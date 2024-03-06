package main

import (
	"reflect"
	"testing"
)

func rotate(nums []int, k int) {
	if k = k % len(nums); k == 0 {
		return
	}
	tmp := make([]int, 0, len(nums))
	tmp = append(tmp, nums[len(nums)-k:]...)
	tmp = append(tmp, nums[:len(nums)-k]...)
	copy(nums, tmp)
}

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		args     args
		wantNums []int
	}{
		{
			name:     "a",
			args:     args{[]int{1, 2, 3, 4}, 3},
			wantNums: []int{2, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.wantNums) {
				t.Errorf("removeDuplicates() nums = %v, wantNums %v", tt.args.nums, tt.wantNums)
			}
		})
	}
}
