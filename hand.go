package main

// a struct to hold a hand of cards
type hand struct {
	cards             []card // 5 cards
	strength          strength
	secondaryStrength int // for tie breakers in the same strength
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
