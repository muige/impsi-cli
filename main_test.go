package main

import "testing"

func TestIsAllowedRunes(t *testing.T) {
	tests := []struct {
		name    string
		runes   []rune
		current string
		want    bool
	}{
		{"valid integer", []rune("123"), "", true},
		{"valid decimal", []rune("12.34"), "", true},
		{"valid negative", []rune("-5"), "", true},
		{"valid negative decimal", []rune("-5.6"), "", true},
		{"single dot", []rune("."), "", true},
		{"zero", []rune("0"), "", true},
		{"invalid letters", []rune("abc"), "", false},
		{"invalid multiple dots", []rune("1.2.3"), "", false},
		{"invalid dash in middle", []rune("5-"), "", false},
		{"invalid dash not first", []rune("-"), "5", false},
		{"invalid mixed", []rune("12a34"), "", false},
		{"dot when dot exists", []rune("."), "12.3", false},
		{"valid digit after dot", []rune("5"), "12.", true},
		{"dash only at start", []rune("-"), "", true},
		{"invalid dash after digit", []rune("-5"), "1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAllowedRunes(tt.runes, tt.current)
			if got != tt.want {
				t.Errorf("isAllowedRunes(%q, %q) = %v; want %v", string(tt.runes), tt.current, got, tt.want)
			}
		})
	}
}
