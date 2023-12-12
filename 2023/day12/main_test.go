package main_test

import (
	mod "day12"
	"testing"
)

func Test_Count(t *testing.T) {
	var testCases = []struct {
		pattern  string
		expected int
	}{
		{"#.#.### 1,1,3", 1},
		{"???.### 1,1,3", 2},
		// 	{"???.### 1,1,3", 0},
		// 	{".??..??...?##. 1,1,3", 0},
		// 	{"?###???????? 3,2,1", 0},
	}
	for _, tt := range testCases {
		pattern := mod.ParseRow(tt.pattern)
		result := pattern.CountArrangements()

		if result != tt.expected {
			t.Errorf("for pattern '%s' expected %v got %v", tt.pattern, tt.expected, result)
		}
	}
}
