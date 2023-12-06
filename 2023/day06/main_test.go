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

func Test_solve(t *testing.T) {
	var testCases = []struct {
		limit, record, lower, upper int
	}{
		{7, 9, 2, 5},
	}
	for _, tt := range testCases {
		lower, upper := mod.Solve(tt.limit, tt.record)
		if lower != tt.lower {
			t.Errorf(
				"For limit=%d and record=%d expected %d, got %d",
				tt.limit, tt.record, tt.lower, lower,
			)
		}
		if upper != tt.upper {
			t.Errorf(
				"For limit=%d and record=%d expected %d, got %d",
				tt.limit, tt.record, tt.upper, upper,
			)
		}
	}
}

func Test_solutions_count(t *testing.T) {
	var testCases = []struct {
		limit, record, count int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
		{71530, 940200, 71503},
	}
	for _, tt := range testCases {
		result := mod.SolutionsCount(tt.limit, tt.record)
		if result != tt.count {
			t.Errorf(
				"For limit=%d and record=%d expected %d, got %d",
				tt.limit, tt.record, tt.count, result,
			)
		}
	}
}

func Test_compare_solutions(t *testing.T) {
	var testCases = []struct {
		limit, record int
	}{
		{7, 9},
		{15, 40},
		{30, 200},
		{71530, 940200},
	}
	for _, tt := range testCases {
		result := mod.SolutionsCount(tt.limit, tt.record)
		explored_result := mod.ExploreSolutions(tt.limit, tt.record)
		if result != explored_result {
			t.Errorf(
				"For limit=%d and record=%d result of exploration is %d, but the analytical result is %d",
				tt.limit, tt.record, explored_result, result,
			)
		}
	}
}
