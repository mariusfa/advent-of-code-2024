package day1

import "testing"

func TestDay1(t *testing.T) {
	leftList := []int{3, 4, 2, 1, 3, 3}
	rightList := []int{4, 3, 5, 3, 9, 3}

	distance := CalcDistance(leftList, rightList)
	if distance != 11 {
		t.Errorf("Expected 11, got %d", distance)
	}
}
