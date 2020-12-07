package codility

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {
	// Map [input]expected
	results := make(map[int]int)

	results[1041] = 5
	results[32] = 0
	results[15] = 0
	results[9] = 2
	results[529] = 4

	for input, expected := range results {
		actual := BinaryGap(input)

		if actual != expected {
			t.Errorf("Expected %d for %d - but got %d", expected, input, actual)
		}
	}
}

func TestRotate(t *testing.T) {
	testitem1 := []int{3, 8, 9, 7, 6}

	testout1 := Rotate(testitem1, 3)

	expected1 := []int{9, 7, 6, 3, 8}

	if arrayequals(testout1, expected1) == false {
		t.Errorf("Expected %d - but got %d", expected1, testout1)
	}

	// Example test:    ([1, 2, 3, 4], 4)

	testitem2 := []int{1, 2, 3, 4}
	expected2 := []int{1, 2, 3, 4}

	testout2 := Rotate(testitem2, 4)

	if arrayequals(testout2, expected2) == false {
		t.Errorf("Expected %d - but got %d", expected2, testout1)
	}

	// Example test:    ([5, -1000], 1)

	testitem3 := []int{5, -1000}
	expected3 := []int{-1000, 5}

	testout3 := Rotate(testitem3, 1)

	if arrayequals(testout3, expected3) == false {
		t.Errorf("Expected %d - but got %d", expected3, testout3)
	}
}

func arrayequals(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
