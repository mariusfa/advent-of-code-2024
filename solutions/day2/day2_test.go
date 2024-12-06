package day2

import "testing"

func TestSafeReports(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	safeReports := getSafeReports(reports)

	expectedSafeReports := 2
	if safeReports != expectedSafeReports {
		t.Errorf("Expected safe reports to be %d, got %d", expectedSafeReports, safeReports)
	}
}

func TestIsTwoValuesWithinValidRange(t *testing.T) {
	if !isTwoValuesWithinValidRange(2, 3) {
		t.Errorf("Expected 2 and 3 to be within valid range")
	}
	if !isTwoValuesWithinValidRange(2, 5) {
		t.Errorf("Expected 2 and 5 to be within valid range")
	}

	if isTwoValuesWithinValidRange(2, 0) {
		t.Errorf("Expected 2 and 1 to be outside valid range")
	}
	if isTwoValuesWithinValidRange(2, 6) {
		t.Errorf("Expected 2 and 6 to be outside valid range")
	}
	if isTwoValuesWithinValidRange(2, 2) {
		t.Errorf("Expected 2 and 2 to be outside valid range")
	}
}
