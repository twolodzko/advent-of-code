package main_test

import (
	mod "day12"
	"testing"
)

// func Test_MatchesPattern(t *testing.T) {
// 	var testCases = []struct {
// 		pattern, example string
// 		matches          bool
// 		unallocated      int
// 	}{
// 		{"???.### 1,1,3", "#.#.###", true, 0},
// 		{"???.### 1,1,3", "???.###", true, 2},
// 		{"???.### 1,1,3", "#??.###", true, 1},
// 		{"???.### 1,1,3", "?#?.###", true, 1},
// 		{"???.### 1,1,3", "??#.###", true, 1},
// 		{"???.### 1,1,3", "#.#.###", true, 0},
// 		{"???.### 1,1,3", "#?#.###", true, 0},
// 		{"???.### 1,1,3", "#.#..##", false, 0},
// 		// {"???.### 1,1,3", "###.###", false, 0},
// 		{"???.### 1,1,3", "#######", false, 0},
// 		{".??..??...?##. 1,1,3", ".#...#....###.", true, 0},
// 		{"?###???????? 3,2,1", ".###.##....#", true, 0},
// 		{"?###???????? 3,2,1", "?###????????", true, 3},
// 		// {"?###???????? 3,2,1", "?####???????", false, 0},
// 		{"?###???????? 3,2,1", "?##???###???", false, 0},
// 	}
// 	for _, tt := range testCases {
// 		pattern := mod.ParseRow(tt.pattern)
// 		input := mod.ParseSprings(tt.example)

// 		matches, unallocated := pattern.MatchesPattern(input)
// 		if matches != tt.matches {
// 			t.Errorf(
// 				"for pattern '%s' and input '%s' expected matches == %v got %v",
// 				tt.pattern, tt.example, tt.matches, matches,
// 			)
// 		}
// 		if unallocated != tt.unallocated {
// 			t.Errorf(
// 				"for pattern '%s' and input '%s' expected unallocated == %v got %v",
// 				tt.pattern, tt.example, tt.unallocated, unallocated,
// 			)
// 		}
// 	}
// }

func Test_MatchesGroups(t *testing.T) {
	var testCases = []struct {
		pattern, example string
		expected         bool
	}{
		{"???.### 1,1,3", "#.#.###", true},
		{"???.### 1,1,3", "#?#.###", true},
		{"???.### 1,1,3", "##?.###", false},
		{"???.### 1,1,3", "?##.###", false},
		{"???.### 1,1,3", ".##.###", false},
		{"???.### 1,1,3", "..#.###", false},
		{"???.### 1,1,3", ".#..###", false},
		{"???.### 1,1,3", "#...###", false},
		{"???.### 1,1,3", "#.?.###", false},
	}
	for _, tt := range testCases {
		pattern := mod.ParseRow(tt.pattern)
		input, _ := mod.ParseSprings(tt.example)
		result := pattern.Matches(input)
		if result != tt.expected {
			t.Errorf("for pattern '%s' and input '%s' got %v expected %v", tt.pattern, tt.example, result, tt.expected)
		}
	}
}
