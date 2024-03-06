package main

import "testing"

func isPalindromeNumber(x int) bool {
	if x%10 == 0 && x != 0 {
		return false
	}
	i, j := x, 0
	for i > j {
		i, j = i/10, j*10+i%10
	}
	return i == j || i == j/10
}

func TestIsPalindromeNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "a",
			n:    121,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeNumber(tt.n); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
