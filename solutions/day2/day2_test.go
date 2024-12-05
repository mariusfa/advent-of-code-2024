package day2

import "testing"

// testdata
// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9

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
