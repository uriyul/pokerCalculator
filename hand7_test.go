package main

import (
	"testing"
)

func TestHand7_Compare(t *testing.T) {

	th := testHands7{}
	th.Init()
	foak := fourOfAKindHands7Test{}
	foak.Init()

	tests := []struct {
		name     string
		hand1    hand7
		hand2    hand7
		expected int
	}{
		{"Same hand", th.handPair, th.handPair, 0},
		{"Pair takes high card", th.handPair, th.handHighCard, 1},
		{"Pair loses to two pair", th.handPair, th.handTwoPair, -1},
		{"Two pair takes pair", th.handTwoPair, th.handPair, 1},
		{"Three of a kind takes two pair", th.handThreeOfAKind, th.handTwoPair, 1},
		{"High pair takes pair", th.handHighPair, th.handPair, 1},
		{"Royal flush takes staright flush", th.handRoyalFlush, th.handStraightFlush, 1},
		{"Straight flush takes four of a kind", th.handStraightFlush, th.handFourOfAKind, 1},
		{"Four of a kind takes full house", th.handFourOfAKind, th.handFullHouse, 1},
		{"Full house takes flush", th.handFullHouse, th.handFlush, 1},
		{"Flush takes straight", th.handFlush, th.handStraight, 1},
		{"Straight takes three of a kind", th.handStraight, th.handThreeOfAKind, 1},
		{"Straight A-5 takes three of a kind", th.handStraightA5, th.handThreeOfAKind, 1},
		{"Straight A-5 beaten by a higher straight", th.handStraightA5, th.handStraight, -1},
		{"TwoPair with high kicker part op pair, takes TwoPair with lower kicker ", th.handTwoPairAAKK789, th.handTwoPairAAKKQQ5, -1},

		{"FOAK kicker Ace takes kicker 10", foak.kickerAce, foak.kicker10, 1},
		{"FOAK kicker 10 takes FOAK with low full", foak.kickerAce, foak.fullHouse, 1},
		{"Straight flush takes FOAK Ace", foak.kickerAce, foak.staightFlush, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handsData := HandsData{
				handsOdds: []HandOdds{{tt.hand1.cards, map[strength]int{}}, {tt.hand2.cards, map[strength]int{}}},
				board:     []card{},
			}
			if got := Compare(&handsData, []card{}); got != tt.expected {
				t.Errorf("Hand7.Compare() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestHand7_Strength(t *testing.T) {
	foak := fourOfAKindHands7Test{}
	foak.Init()

	tests := []struct {
		name     string
		hand     hand7
		expected strength
	}{
		{"Four of a kind with kicker Ace", foak.kickerAce, FourOfAKind},
		{"Four of a kind with kicker 10", foak.kicker10, FourOfAKind},
		{"Full house", foak.fullHouse, FourOfAKind},
		{"Straight flush", foak.staightFlush, StraightFlush},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hand.Strength(&HandOdds{mapOdds: map[strength]int{}}); got != tt.expected {
				t.Errorf("Hand7.Strength() = %v, want %v", got, tt.expected)
			}
		})
	}

}
