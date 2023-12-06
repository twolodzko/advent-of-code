package main_test

import (
	mod "day06"
	"testing"
)

func Test_distance(t *testing.T) {
	var testCases = []struct {
		hold_time, limit, expected int
	}{
		{0, 7, 0},
		{1, 7, 6},
		{2, 7, 10},
		{3, 7, 12},
		{4, 7, 12},
		{5, 7, 10},
		{6, 7, 6},
		{7, 7, 0},
	}
	for _, tt := range testCases {
		result := mod.Distance(tt.hold_time, tt.limit)
		if result != tt.expected {
			t.Errorf(
				"For hold time %d and limit %v we expected the distance %d but got %v",
				tt.hold_time, tt.limit, tt.expected, result,
			)
		}
	}
}
