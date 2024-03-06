package main

import (
	"sort"
	"testing"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	hIndex := 0
	for i, c := range citations {
		if i+1 <= c {
			hIndex = i + 1
		} else {
			break
		}
	}
	return hIndex
}

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{
			name:      "a",
			citations: []int{3, 0, 6, 1, 5},
			want:      3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
