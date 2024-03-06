package main

import (
	"strings"
	"testing"
)

var romanMap = map[byte][]struct {
	Symbol string
	Value  int
}{
	'I': {{"IV", 4}, {"IX", 9}},
	'X': {{"XL", 40}, {"XC", 90}},
	'C': {{"CD", 400}, {"CM", 900}},
}

var romanSym = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := 0
	i := 0

	read := func() {
		if s[i] == 'I' || s[i] == 'X' || s[i] == 'C' {
			for _, e := range romanMap[s[i]] {
				if strings.HasPrefix(s[i:], e.Symbol) {
					n += e.Value
					i += len(e.Symbol)
					return
				}
			}
		}
		n += romanSym[s[i]]
		i++
	}

	for i < len(s) {
		read()
	}
	return n
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "a",
			s:    "VII",
			want: 7,
		},
		{
			name: "b",
			s:    "MCMXCIV",
			want: 1994,
		},
		{
			name: "c",
			s:    "MMMXLV",
			want: 3045,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
