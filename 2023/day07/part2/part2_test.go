package part2_test

import (
	"day07/part2"
	"fmt"
	"testing"
)

func Test_counts(t *testing.T) {
	var (
		counts           part2.Counter
		result, expected part2.Count
	)
	result = counts.First()
	if result.Value != 0 {
		t.Errorf("Expected nothing for empty counter, got %v", result)
	}

	counts.Add('A')
	result = counts.First()
	expected = part2.Count{'A', 1}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	counts.Add('A')
	result = counts.First()
	expected = part2.Count{'A', 2}
	if result != expected {
		fmt.Println(counts)
		t.Errorf("Expected %v, got %v", expected, result)
	}

	counts.Add('K')
	result = counts.First()
	expected = part2.Count{'A', 2}
	if result != expected {
		fmt.Println(counts)
		t.Errorf("Expected %v, got %v", expected, result)
	}
	if counts.Len() != 2 {
		t.Errorf("incorrectly counts.Len() = %d", counts.Len())
	}

	if counts.Pop('A') != 2 {
		t.Errorf("Invalid popped value")
	}
	if counts.Len() != 1 {
		t.Errorf("incorrectly counts.Len() = %d", counts.Len())
	}

	if counts.Pop('K') != 1 {
		t.Errorf("Invalid popped value")
	}
	if counts.Len() != 0 {
		t.Errorf("incorrectly counts.Len() = %d", counts.Len())
	}
}
