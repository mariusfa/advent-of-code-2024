package utils

import "testing"

func TestReadData(t *testing.T) {
	testFile := "datatest.txt"
	contentString, err := readData(testFile)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}
	println(contentString)
	left, right, err := parseData(contentString)
	if err != nil {
		t.Errorf("Error parsing data: %s", err)
	}

	printList(left)

	leftExpected := []int{49744, 20738, 20319}
	rightExpected := []int{57964, 85861, 65072}

	for i := 0; i < len(left); i++ {
		if left[i] != leftExpected[i] {
			t.Errorf("Expected %d, got %d", leftExpected[i], left[i])
		}
		if right[i] != rightExpected[i] {
			t.Errorf("Expected %d, got %d", rightExpected[i], right[i])
		}
	}
}
