package main

import (
	"testing"
)

func TestHand_Compare(t *testing.T) {

	th := testHands{}
	th.Init()

	tests := []struct {
		name     string
		hand1    hand
		hand2    hand
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hand1.Compare(tt.hand2); got != tt.expected {
				t.Errorf("Hand.Compare() = %v, want %v", got, tt.expected)
			}
		})
	}
}
