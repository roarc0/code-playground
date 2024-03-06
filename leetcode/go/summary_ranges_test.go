package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func summaryRanges(nums []int) (summary []string) {
	if len(nums) == 0 {
		return nil
	}

	i, j := nums[0], 0
	jOk := false
	render := func() string {
		if !jOk {
			return strconv.Itoa(i)
		}
		return fmt.Sprintf("%d->%d", i, j)
	}

	for _, n := range nums[1:] {
		if !jOk && n == i+1 {
			jOk = true
			j = n
		} else if jOk && n == j+1 {
			j = n
		} else {
			summary = append(summary, render())
			jOk = false
			i = n
		}
	}
	summary = append(summary, render())
	return
}

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		wantSummary []string
	}{
		{
			name:        "a",
			nums:        []int{0, 1, 2, 4, 5, 7},
			wantSummary: []string{"0->2", "4->5", "7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSummary := summaryRanges(tt.nums); !reflect.DeepEqual(gotSummary, tt.wantSummary) {
				t.Errorf("summaryRanges() = %v, want %v", gotSummary, tt.wantSummary)
			}
		})
	}
}
