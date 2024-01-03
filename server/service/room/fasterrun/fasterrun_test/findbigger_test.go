package mahjong

import (
	"fmt"
	"reflect"
	"testing"

	"server/service/room/fasterrun"
)

func TestFindBigger(t *testing.T) {
	testCases := []struct {
		name  string
		input fasterrun.PokerCards
		arg   fasterrun.CardsGroup
		want  []fasterrun.CardsGroup
	}{

		// Single
		{
			name:  "单张",
			input: fasterrun.PokerCards{fasterrun.Clubs_10, fasterrun.Clubs_4, fasterrun.Hearts_10, fasterrun.Spades_5},
			arg: fasterrun.CardsGroup{
				Type:      fasterrun.Single,
				Cards:     fasterrun.PokerCards{fasterrun.Clubs_5},
				SideCards: nil,
			},
			want: []fasterrun.CardsGroup{
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
			name:  "对子",
			input: fasterrun.PokerCards{fasterrun.Clubs_10, fasterrun.Clubs_4, fasterrun.Hearts_4, fasterrun.Hearts_10, fasterrun.Spades_6, fasterrun.Diamonds_6, fasterrun.Clubs_5},
			arg: fasterrun.CardsGroup{
				Type:      fasterrun.Pair,
				Cards:     fasterrun.PokerCards{fasterrun.Spades_5, fasterrun.Diamonds_5},
				SideCards: nil,
			},
			want: []fasterrun.CardsGroup{
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

		// Trips
		{
			name: "三张",
			input: fasterrun.PokerCards{
				fasterrun.Clubs_5,
				fasterrun.Spades_6,
				fasterrun.Diamonds_6,

				fasterrun.Clubs_9,
				fasterrun.Hearts_9,
				fasterrun.Spades_9,

				fasterrun.Clubs_10,
				fasterrun.Hearts_10,
				fasterrun.Spades_10,
			},
			arg: fasterrun.CardsGroup{
				Type:      fasterrun.Trips,
				Cards:     fasterrun.PokerCards{fasterrun.Spades_5, fasterrun.Diamonds_5, fasterrun.Hearts_5},
				SideCards: nil,
			},
			want: []fasterrun.CardsGroup{
				{
					Type: fasterrun.Trips,
					Cards: fasterrun.PokerCards{
						fasterrun.Clubs_9,
						fasterrun.Hearts_9,
						fasterrun.Spades_9,
					},
				},
				{
					Type: fasterrun.Trips,
					Cards: fasterrun.PokerCards{
						fasterrun.Clubs_10,
						fasterrun.Hearts_10,
						fasterrun.Spades_10,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.FindBigger(tc.arg)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Input: %v, n: %d\nExpected: %v\nActual: %v", tc.input, tc.arg, tc.want, result)
			}
		})
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

	cards = fasterrun.PokerCards{106, 206, 107, 307, 407, 108, 310, 411, 212, 412, 313}
	bigger = cards.FindBigger(fasterrun.CardsGroup{
		Type:  fasterrun.Pair,
		Cards: fasterrun.PokerCards{203, 303},
	})
	fmt.Println(bigger)
}

func TestFindBiggerDebug(t *testing.T) {
	cards := fasterrun.PokerCards{103, 203, 303, 204, 304, 404, 405, 406}
	dst := cards.AnalyzeCards(true)
	fmt.Println(dst)
}
func TestFindBiggerDebug2(t *testing.T) {
	cards := fasterrun.PokerCards{205, 305, 106, 206, 306, 408, 409, 213, 313}
	cardGroup := fasterrun.CardsGroup{
		Type:      fasterrun.PlaneWithTwo,
		Cards:     fasterrun.PokerCards{103, 203, 303, 204, 304, 404},
		SideCards: fasterrun.PokerCards{405, 406, 410, 411},
	}
	biggerCardGroups := cards.FindBigger(cardGroup)
	one := biggerCardGroups[0]
	for len(one.SideCards) > len(cardGroup.SideCards) {
		one.SideCards = one.SideCards[1:]
	}
	fmt.Println(biggerCardGroups)
}
