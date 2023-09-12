package mahjong

import (
	"reflect"
	"testing"

	"server/service/room/fasterrun"
)

func TestCombination(t *testing.T) {
	testCases := []struct {
		input    fasterrun.PokerCards
		n        int
		expected []fasterrun.PokerCards
	}{
		//
		{
			input: fasterrun.PokerCards{1, 2, 3, 4},
			n:     2,
			expected: []fasterrun.PokerCards{
				{4, 3}, {4, 2}, {4, 1}, {3, 2}, {3, 1}, {2, 1},
			},
		},

		//
		{
			input: fasterrun.PokerCards{1, 2, 3, 4},
			n:     3,
			expected: []fasterrun.PokerCards{
				{4, 3, 2},
				{4, 3, 1},
				{4, 2, 1},
				{3, 2, 1},
			},
		},
	}

	for _, tc := range testCases {
		result := tc.input.Combination(tc.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Input: %v, n: %d\nExpected: %v\nActual: %v", tc.input, tc.n, tc.expected, result)
		}
	}
}
