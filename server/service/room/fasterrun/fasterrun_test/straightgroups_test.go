package mahjong

import (
	"reflect"
	"testing"

	"server/service/room/fasterrun"
)

// 测试筛查顺子，连对，飞机
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

		// Straight2
		{
			input: fasterrun.PokerCards{3, 3, 3, 3, 4, 5, 6, 7, 8, 9, 9, 10, 10, 11},
			t:     fasterrun.Straight,
			n:     5,
			expected: []fasterrun.PokerCards{
				{3, 4, 5, 6, 7},
				{4, 5, 6, 7, 8},
				{5, 6, 7, 8, 9},
				{6, 7, 8, 9, 10},
				{7, 8, 9, 10, 11},
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

		//StraightPair2
		{
			input: fasterrun.PokerCards{3, 3, 4, 5, 5, 6, 6, 7, 8, 9, 9, 10, 10, 10},
			t:     fasterrun.StraightPair,
			n:     2,
			expected: []fasterrun.PokerCards{
				{5, 5, 6, 6},
				{9, 9, 10, 10},
			},
		},

		//StraightPair3
		{
			input: fasterrun.PokerCards{3, 3, 4, 5, 5, 6, 6, 6, 7, 7, 8, 9, 9, 9, 9, 10, 10, 10, 10},
			t:     fasterrun.StraightPair,
			n:     2,
			expected: []fasterrun.PokerCards{
				{5, 5, 6, 6},
				{6, 6, 7, 7},
				{9, 9, 10, 10},
			},
		},

		//StraightPair4
		{
			input: fasterrun.PokerCards{3, 3, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 9, 9, 10, 10, 10, 10},
			t:     fasterrun.StraightPair,
			n:     3,
			expected: []fasterrun.PokerCards{
				{5, 5, 6, 6, 7, 7},
				{6, 6, 7, 7, 8, 8},
				{7, 7, 8, 8, 9, 9},
				{8, 8, 9, 9, 10, 10},
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

		//Plane2
		{
			input: fasterrun.PokerCards{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10},
			t:     fasterrun.Plane,
			n:     3,
			expected: []fasterrun.PokerCards{
				{3, 3, 3, 4, 4, 4, 5, 5, 5},
				{4, 4, 4, 5, 5, 5, 6, 6, 6},
				{8, 8, 8, 9, 9, 9, 10, 10, 10},
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
