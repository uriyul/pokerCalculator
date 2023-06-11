package main

import (
	"fmt"
	"sort"
)

// a struct to hold a hand of cards
type hand struct {
	cards             []card
	secondaryStrength int
	strength          strength
}

type strength int

const (
	HighCard      strength = iota + 1 // 1
	Pair                              // 2
	TwoPair                           // 3
	ThreeOfAKind                      // 4
	Straight                          // 5
	Flush                             // 6
	FullHouse                         // 7
	FourOfAKind                       // 8
	StraightFlush                     // 9
	RoyalFlush                        // 10
)

var strength2String = map[strength]string{
	HighCard:      "High Card      ",
	Pair:          "Pair           ",
	TwoPair:       "Two Pair       ",
	ThreeOfAKind:  "Three of a Kind",
	Straight:      "Straight       ",
	Flush:         "Flush          ",
	FullHouse:     "Full House     ",
	FourOfAKind:   "Four of a Kind ",
	StraightFlush: "Straight Flush ",
	RoyalFlush:    "Royal Flush    ",
}

// String method for hand
func (h hand) String() string {
	return fmt.Sprintf("%v", h.cards)
}

func (h *hand) Strength() strength {

	if h.strength != 0 {
		return h.strength
	}

	h.sort()

	switch {
	case h.isRoyalFlush():
		h.strength = RoyalFlush
	case h.isStraightFlush():
		h.strength = StraightFlush
	case h.isFourOfAKind():
		h.strength = FourOfAKind
	case h.isFullHouse():
		h.strength = FullHouse
	case h.isFlush():
		h.strength = Flush
	case h.isStraight():
		h.strength = Straight
	case h.isThreeOfAKind():
		h.strength = ThreeOfAKind
	case h.isTwoPair():
		h.strength = TwoPair
	case h.isPair():
		h.strength = Pair
	default:
		h.HighCard()
		h.strength = HighCard
	}

	return h.strength
}

func (h hand) isRoyalFlush() bool {
	return h.isStraightFlush() && h.cards[4].value == 14
}

func (h *hand) isStraightFlush() bool {
	if h.isFlush() && h.isStraight() {
		h.secondaryStrength = int(h.cards[4].value)
		return true
	}
	return false
}

func (h *hand) isFourOfAKind() bool {
	if h.cards[0].value == h.cards[3].value || h.cards[1].value == h.cards[4].value {
		h.secondaryStrength = int(h.cards[2].value)
		return true
	}
	return false
}

func (h *hand) isFullHouse() bool {
	if h.cards[0].value == h.cards[2].value {
		if h.cards[3].value == h.cards[4].value {
			h.secondaryStrength = int(h.cards[0].value)
			return true
		}
	}
	if h.cards[0].value == h.cards[1].value {
		if h.cards[2].value == h.cards[4].value {
			h.secondaryStrength = int(h.cards[2].value)
			return true
		}
	}
	return false
}

func (h *hand) isFlush() bool {
	suit := h.cards[0].suit
	for _, card := range h.cards {
		if card.suit != suit {
			return false
		}
	}
	h.secondaryStrength = int(h.cards[4].value)
	return true
}

func (h *hand) isStraight() bool {
	for i := 0; i < len(h.cards)-1; i++ {
		if h.cards[i].value+1 != h.cards[i+1].value {
			return false
		}
		// special case for Ace low straight
		if i == 3 && h.cards[i+1].value == 14 && h.cards[0].value == 2 {
			h.secondaryStrength = int(h.cards[3].value)
			return true
		}
	}
	h.secondaryStrength = int(h.cards[4].value)
	return true
}

func (h *hand) isThreeOfAKind() bool {
	if h.cards[0].value == h.cards[2].value {
		h.secondaryStrength = int(h.cards[0].value)*100 + int(h.cards[4].value)
		return true
	}
	if h.cards[1].value == h.cards[3].value {
		h.secondaryStrength = int(h.cards[1].value)*100 + int(h.cards[4].value)
		return true
	}
	if h.cards[2].value == h.cards[4].value {
		h.secondaryStrength = int(h.cards[2].value)*100 + int(h.cards[1].value)
		return true
	}
	return false
}

func (h *hand) isTwoPair() bool {
	if h.cards[0].value == h.cards[1].value {
		if h.cards[2].value == h.cards[3].value {
			h.secondaryStrength = int(h.cards[3].value)*10000 + int(h.cards[1].value)*100 + int(h.cards[4].value)
			return true
		}
		if h.cards[3].value == h.cards[4].value {
			h.secondaryStrength = int(h.cards[3].value)*10000 + int(h.cards[1].value)*100 + int(h.cards[2].value)
			return true
		}
	}
	if h.cards[1].value == h.cards[2].value {
		if h.cards[3].value == h.cards[4].value {
			h.secondaryStrength = int(h.cards[3].value)*10000 + int(h.cards[1].value)*100 + int(h.cards[0].value)
			return true
		}
	}
	return false
}

func (h *hand) isPair() bool {
	strength := 0
	found := false
	for i := 0; i < len(h.cards)-1; i++ {
		if h.cards[i].value == h.cards[i+1].value {
			strength += int(h.cards[i].value) * (i + 1) * 100000
			found = true
		} else {
			strength += int(h.cards[i].value) * (i + 1) * 100
		}
	}

	if found {
		h.secondaryStrength = strength
		//fmt.Printf("$$$ Pair %v: %d\n", h, h.strength)
		return true
	}

	return false
}

func (h *hand) HighCard() {
	h.secondaryStrength = int(h.cards[4].value)
	for i := 3; i >= 0; i-- {
		h.secondaryStrength *= 100
		h.secondaryStrength += int(h.cards[i].value)
	}
}

func (h *hand) Compare(other hand) int {

	hs := h.Strength()
	os := other.Strength()

	if hs > os {
		return 1
	}
	if hs < os {
		return -1
	}

	if h.secondaryStrength == other.secondaryStrength {
		return 0
	}
	if h.secondaryStrength-other.secondaryStrength > 0 {
		return 1
	}
	return -1

}

func (h *hand) sort() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].value < h.cards[j].value
	})
}
