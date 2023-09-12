package mahjong

import (
	"reflect"
	"testing"

	"server/service/room/fasterrun"
)

func TestStraightGroups(t *testing.T) {
	testCases := []struct {
		input    fasterrun.PokerCards
		t        fasterrun.PokerCardsType
		n        int
		expected []fasterrun.PokerCards
	}{
		// Straight
		{
			input: fasterrun.PokerCards{3, 4, 5, 6, 7, 8, 9, 9, 10, 10},
			t:     fasterrun.Straight,
			n:     5,
			expected: []fasterrun.PokerCards{
				{3, 4, 5, 6, 7},
				{4, 5, 6, 7, 8},
				{5, 6, 7, 8, 9},
				{6, 7, 8, 9, 10},
			},
		},

		// Straight
		{
			input: fasterrun.PokerCards{3, 4, 5, 6, 7, 8, 9, 9, 10, 10},
			t:     fasterrun.Straight,
			n:     5,
			expected: []fasterrun.PokerCards{
				{3, 4, 5, 6, 7},
				{4, 5, 6, 7, 8},
				{5, 6, 7, 8, 9},
				{6, 7, 8, 9, 10},
			},
		},

		//StraightPair
		{
			input: fasterrun.PokerCards{3, 3, 4, 4, 5, 5, 6, 6, 7, 8, 9, 9, 10, 10, 10},
			t:     fasterrun.StraightPair,
			n:     2,
			expected: []fasterrun.PokerCards{
				{3, 3, 4, 4},
				{4, 4, 5, 5},
				{5, 5, 6, 6},
				{9, 9, 10, 10},
			},
		},

		//Plane
		{
			input: fasterrun.PokerCards{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9, 9, 10, 10, 10},
			t:     fasterrun.Plane,
			n:     2,
			expected: []fasterrun.PokerCards{
				{3, 3, 3, 4, 4, 4},
				{4, 4, 4, 5, 5, 5},
				{9, 9, 9, 10, 10, 10},
			},
		},
	}

	for _, tc := range testCases {
		result := tc.input.StraightGroups(tc.t, tc.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Input: %v, n: %d\nExpected: %v\nActual: %v", tc.input, tc.n, tc.expected, result)
		}
	}
}
