package main

import (
	"reflect"
	"testing"
)

// merge nums2 slice into nums1 slice keeping the ascending order.
// len(nums1)==n+m, len(nums2)==n
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, 0; j < n; i++ {
		if nums1[i] > nums2[j] {
			m++
			copy(nums1[i+1:m], nums1[i:m])
			nums1[i], nums1[i+1] = nums2[j], nums1[i]
			j++
		} else if i >= m {
			nums1[i] = nums2[j]
			j++
		}
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name          string
		args          args
		expectedNums1 []int
	}{
		{
			name:          "a",
			args:          args{[]int{1, 0}, 1, []int{2}, 1},
			expectedNums1: []int{1, 2},
		},
		{
			name:          "b",
			args:          args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			expectedNums1: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:          "c",
			args:          args{[]int{1, 2, 2, 3}, 1, []int{}, 0},
			expectedNums1: []int{1, 2, 2, 3},
		},
		{
			name:          "d",
			args:          args{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3},
			expectedNums1: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:          "e",
			args:          args{[]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5},
			expectedNums1: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:          "f",
			args:          args{[]int{1, 2, 4, 5, 6, 0}, 5, []int{3}, 1},
			expectedNums1: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.expectedNums1) {
				t.Errorf("Expected nums1=%v, got=%v", tt.expectedNums1, tt.args.nums1)
			}
		})
	}
}
