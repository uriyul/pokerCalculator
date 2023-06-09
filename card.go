package main

import "fmt"

// struct to hold a card
type card struct {
	suit  Suit
	value byte // 2 - 14 where 14 is an Ace
}

type Suit int

const (
	Spades   Suit = iota + 1 // 1
	Diamonds                 // 2
	Clubs                    // 3
	Hearts                   // 4
)

// String method for card
func (c card) String() string {

	switch c.value {
	case 10:
		return fmt.Sprintf("T%s", suitMap[c.suit])
	case 11:
		return fmt.Sprintf("J%s", suitMap[c.suit])
	case 12:
		return fmt.Sprintf("Q%s", suitMap[c.suit])
	case 13:
		return fmt.Sprintf("K%s", suitMap[c.suit])
	case 14:
		return fmt.Sprintf("A%s", suitMap[c.suit])
	default:
		return fmt.Sprintf("%d%s", c.value, suitMap[c.suit])

	}
}

var suitMap = map[Suit]string{
	Spades:   "♠",
	Diamonds: "♦",
	Clubs:    "♣",
	Hearts:   "♥",
}

//var deck []card

// create a new deck of cards
func newDeck() []card {
	deck := []card{}
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	for _, suit := range suits {
		for i := 2; i <= 14; i++ {
			deck = append(deck, card{suit, byte(i)})
		}
	}
	return deck
}
