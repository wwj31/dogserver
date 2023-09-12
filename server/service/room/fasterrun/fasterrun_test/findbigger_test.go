package mahjong

import (
	"fmt"
	"reflect"
	"testing"

	"server/service/room/fasterrun"
)

func TestFindBigger(t *testing.T) {
	testCases := []struct {
		input    fasterrun.PokerCards
		arg      fasterrun.CardsGroup
		expected []fasterrun.CardsGroup
	}{

		// Single
		{
			input: fasterrun.PokerCards{fasterrun.Clubs_10, fasterrun.Clubs_4, fasterrun.Hearts_10, fasterrun.Spades_5},
			arg: fasterrun.CardsGroup{
				Type:      fasterrun.Single,
				Cards:     fasterrun.PokerCards{fasterrun.Clubs_5},
				SideCards: nil,
			},
			expected: []fasterrun.CardsGroup{
				{
					Type:      fasterrun.Single,
					Cards:     fasterrun.PokerCards{fasterrun.Clubs_10},
					SideCards: nil,
				},
				{
					Type:      fasterrun.Single,
					Cards:     fasterrun.PokerCards{fasterrun.Hearts_10},
					SideCards: nil,
				},
			},
		},

		// Pair
		{
			input: fasterrun.PokerCards{fasterrun.Clubs_10, fasterrun.Clubs_4, fasterrun.Hearts_4, fasterrun.Hearts_10, fasterrun.Spades_6, fasterrun.Diamonds_6, fasterrun.Clubs_5},
			arg: fasterrun.CardsGroup{
				Type:      fasterrun.Pair,
				Cards:     fasterrun.PokerCards{fasterrun.Spades_5, fasterrun.Diamonds_5},
				SideCards: nil,
			},
			expected: []fasterrun.CardsGroup{
				{
					Type:      fasterrun.Pair,
					Cards:     fasterrun.PokerCards{fasterrun.Clubs_10, fasterrun.Hearts_10},
					SideCards: nil,
				},
				{
					Type:      fasterrun.Pair,
					Cards:     fasterrun.PokerCards{fasterrun.Spades_6, fasterrun.Diamonds_6},
					SideCards: nil,
				},
			},
		},
	}

	for _, tc := range testCases {
		result := tc.input.FindBigger(tc.arg)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Input: %v, n: %d\nExpected: %v\nActual: %v", tc.input, tc.arg, tc.expected, result)
		}
	}
}

func TestFindBiggerPlaneWithTwo(t *testing.T) {
	cards := fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 12, 12, 13, 14}
	bigger := cards.FindBigger(fasterrun.CardsGroup{
		Type:      fasterrun.PlaneWithTwo,
		Cards:     fasterrun.PokerCards{3, 3, 3, 4, 4, 4, 5, 5, 5},
		SideCards: fasterrun.PokerCards{7, 9, 10, 11, 14, 14},
	})
	fmt.Println(bigger)

	cards = fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 11, 12, 12, 13, 14}
	bigger = cards.FindBigger(fasterrun.CardsGroup{
		Type:      fasterrun.PlaneWithTwo,
		Cards:     fasterrun.PokerCards{4, 4, 4, 5, 5, 5, 6, 6, 6},
		SideCards: fasterrun.PokerCards{7, 9, 10, 11, 14, 14},
	})
	fmt.Println(bigger)
}

func TestFindBiggerPlane(t *testing.T) {
	cards := fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 12, 12, 13, 14}
	bigger := cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.Plane,
		Cards: fasterrun.PokerCards{3, 3, 3, 4, 4, 4, 5, 5, 5},
	})
	fmt.Println(bigger)

	cards = fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 11, 12, 12, 13, 14}
	bigger = cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.Plane,
		Cards: fasterrun.PokerCards{4, 4, 4, 5, 5, 5, 6, 6, 6},
	})
	fmt.Println(bigger)
}

func TestFindBiggerStraightPair(t *testing.T) {
	cards := fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 12, 12, 13, 14}
	bigger := cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.StraightPair,
		Cards: fasterrun.PokerCards{3, 3, 4, 4, 5, 5},
	})
	fmt.Println(bigger)

	cards = fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 11, 11, 12, 12, 13, 13}
	bigger = cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.StraightPair,
		Cards: fasterrun.PokerCards{4, 4, 5, 5, 6, 6},
	})
	fmt.Println(bigger)

	cards = fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 11, 11, 12, 12, 13, 13}
	bigger = cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.StraightPair,
		Cards: fasterrun.PokerCards{4, 4, 5, 5},
	})
	fmt.Println(bigger)
}

func TestFindBiggerStraight(t *testing.T) {
	cards := fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 10, 11, 12, 12, 13, 14}
	bigger := cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.Straight,
		Cards: fasterrun.PokerCards{4, 5, 6, 7, 8, 9, 10},
	})
	fmt.Println(bigger)

	cards = fasterrun.PokerCards{6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 11, 11, 12, 12, 13, 13}
	bigger = cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.Straight,
		Cards: fasterrun.PokerCards{6, 7, 8, 9, 10},
	})
	fmt.Println(bigger)
}
