package main

import (
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	validSymbol := func(i int) bool { return unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) }
	for i, j := 0, len(s)-1; j > i; {
		iOk := false
		for i < len(s) && !iOk {
			if iOk = validSymbol(i); !iOk {
				i++
			}
		}

		jOk := false
		for j > 0 && !jOk {
			if jOk = validSymbol(j); !jOk {
				j--
			}
		}

		if iOk && jOk && unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "a",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "b",
			s:    ",,,,,,,,,,,,acva",
			want: false,
		},
		{
			name: "c",
			s:    " ",
			want: true,
		},
		{
			name: "d",
			s:    "v' 5:UxU:5 v'",
			want: true,
		},
		{
			name: "e",
			s:    ".,",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
