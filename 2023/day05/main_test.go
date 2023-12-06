package main_test

// import (
// 	mod "day05"
// 	"testing"
// )

// func Test_range_no_overlap_cross(t *testing.T) {
// 	this := mod.NewRange(0, 5)
// 	other := mod.NewRange(7, 10)
// 	before, intersection, after := this.Cross(other)

// 	if !before.Equal(this) {
// 		t.Errorf("Expected before: %v, got %v", before, this)
// 	}
// 	if intersection.IsValid() {
// 		t.Errorf("Expected no intersection, got %v", intersection)
// 	}
// 	if after.IsValid() {
// 		t.Errorf("Expected no after, got %v", after)
// 	}
// }

// func Test_range_full_overlap_cross(t *testing.T) {
// 	this := mod.NewRange(0, 5)
// 	other := mod.NewRange(0, 5)
// 	before, intersection, after := this.Cross(other)

// 	if before.IsValid() {
// 		t.Errorf("Expected no before, got %v", this)
// 	}
// 	if !intersection.Equal(this) {
// 		t.Errorf("Expected intersection: %v, got %v", this, intersection)
// 	}
// 	if after.IsValid() {
// 		t.Errorf("Expected no after, got %v", after)
// 	}
// }

// func Test_range_partial_overlap_cross(t *testing.T) {
// 	this := mod.NewRange(0, 5)
// 	other := mod.NewRange(2, 7)
// 	before, intersection, after := this.Cross(other)

// 	expected_before := mod.NewRange(0, 1)
// 	if !before.Equal(expected_before) {
// 		t.Errorf("Expected before: %v, got %v", expected_before, before)
// 	}
// 	expected_intersection := mod.NewRange(2, 5)
// 	if !intersection.Equal(expected_intersection) {
// 		t.Errorf("Expected intersection: %v, got %v", expected_intersection, intersection)
// 	}
// 	if after.IsValid() {
// 		t.Errorf("Expected no after, got %v", after)
// 	}
// }

// func Test_range_partial_overlap_behind_cross(t *testing.T) {
// 	this := mod.NewRange(0, 5)
// 	other := mod.NewRange(2, 7)
// 	before, intersection, after := this.Cross(other)

// 	expected_before := mod.NewRange(0, 1)
// 	if !before.Equal(expected_before) {
// 		t.Errorf("Expected before: %v, got %v", expected_before, before)
// 	}
// 	expected_intersection := mod.NewRange(2, 5)
// 	if !intersection.Equal(expected_intersection) {
// 		t.Errorf("Expected intersection: %v, got %v", expected_intersection, intersection)
// 	}
// 	if after.IsValid() {
// 		t.Errorf("Expected no after, got %v", after)
// 	}
// }

// func Test_range_partial_overlap_before_cross(t *testing.T) {
// 	this := mod.NewRange(5, 10)
// 	other := mod.NewRange(2, 7)
// 	before, intersection, after := this.Cross(other)

// 	if before.IsValid() {
// 		t.Errorf("Expected no before, got %v", before)
// 	}
// 	expected_intersection := mod.NewRange(5, 7)
// 	if !intersection.Equal(expected_intersection) {
// 		t.Errorf("Expected intersection: %v, got %v", expected_intersection, intersection)
// 	}
// 	expected_after := mod.NewRange(8, 10)
// 	if !after.Equal(expected_after) {
// 		t.Errorf("Expected after: %v, got %v", expected_after, after)
// 	}
// }

// func Test_range_partial_overlap_middle_cross(t *testing.T) {
// 	this := mod.NewRange(0, 10)
// 	other := mod.NewRange(3, 6)
// 	before, intersection, after := this.Cross(other)

// 	expected_before := mod.NewRange(0, 2)
// 	if !before.Equal(expected_before) {
// 		t.Errorf("Expected before: %v, got %v", expected_before, intersection)
// 	}
// 	expected_intersection := other
// 	if !intersection.Equal(expected_intersection) {
// 		t.Errorf("Expected intersection: %v, got %v", expected_intersection, intersection)
// 	}
// 	expected_after := mod.NewRange(7, 10)
// 	if !after.Equal(expected_after) {
// 		t.Errorf("Expected after: %v, got %v", expected_after, after)
// 	}
// }
