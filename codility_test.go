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
