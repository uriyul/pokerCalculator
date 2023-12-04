package main

import (
	"sort"
	"sync"
)

// a struct to hold a hand of cards
type hand7 struct {
	cards []card // 7 cards
	hand  hand   // 5 chosen cards out of 7
}

func (h *hand7) Strength(handOdds *HandOdds) strength {
	if h.hand.strength != 0 {
		return h.hand.strength
	}

	h.sort()

	h.hand.cards = make([]card, 5)

	switch {
	case h.isStraightFlush():
		if h.hand.cards[4].value == 14 && h.hand.cards[0].value == 10 {
			h.hand.strength = RoyalFlush
		} else {
			h.hand.strength = StraightFlush
		}
	case h.isFourOfAKind():
		h.hand.strength = FourOfAKind
	case h.isFullHouse():
		h.hand.strength = FullHouse
	case h.isFlush():
		h.hand.strength = Flush
	case h.isStraight():
		h.hand.strength = Straight
	case h.isThreeOfAKind():
		h.hand.strength = ThreeOfAKind
	case h.isTwoPair():
		h.hand.strength = TwoPair
	case h.isPair():
		h.hand.strength = Pair
	default:
		h.HighCard()
		h.hand.strength = HighCard
	}

	safeMapIncrement(handOdds.mapOdds, h.hand.strength)

	return h.hand.strength
}

var mu sync.Mutex = sync.Mutex{}

func safeMapIncrement(m map[strength]int, s strength) {
	mu.Lock()
	defer mu.Unlock()
	m[s]++
}

func (h *hand7) isStraightFlush() bool {
	suited := h.getSuited()
	if len(suited) < 5 {
		return false
	}

	sort.Slice(suited, func(i, j int) bool {
		return suited[i].value < suited[j].value
	})

	for i := len(suited) - 5; i >= 0; i-- {
		if suited[i].value+4 == suited[i+4].value {
			h.hand.cards = suited[i : i+5]
			h.hand.secondaryStrength = int(suited[i+4].value)
			return true
		}
	}

	// treat case of A2345 as a special case
	if suited[0].value == 2 && suited[3].value == 5 && suited[len(suited)-1].value == 14 {
		h.hand.cards = suited[0:4]
		h.hand.cards = append(h.hand.cards, suited[len(suited)-1])
		h.hand.secondaryStrength = 5
		return true
	}

	return false
}

func (h *hand7) isFourOfAKind() bool {
	for i := 0; i < 4; i++ {
		if h.cards[i].value == h.cards[i+3].value {
			h.hand.cards = h.cards[i : i+4]
			if i <= 2 {
				h.hand.cards = append(h.hand.cards, h.cards[6])
				h.hand.secondaryStrength = int(h.cards[6].value)
			} else {
				h.hand.cards = append(h.cards, h.cards[2])
				h.hand.secondaryStrength = int(h.cards[2].value)
			}
			return true
		}
	}
	return false
}

// return cards of the same suit
func (h hand7) getSuited() []card {
	var suited []card
	h.sortSuit()
	defer h.sort()

	for _, card := range h.cards {
		if card.suit == h.cards[3].suit {
			suited = append(suited, card)
		}
	}
	return suited
}

func (h *hand7) isFlush() bool {
	h.sortSuit()
	defer h.sort()

	switch {
	case h.cards[2].suit == h.cards[6].suit:
		setFlush(h.cards[2:], &h.hand)
		return true
	case h.cards[1].suit == h.cards[5].suit:
		setFlush(h.cards[1:6], &h.hand)
		return true
	case h.cards[0].suit == h.cards[4].suit:
		setFlush(h.cards[:5], &h.hand)
		return true
	default:
		return false
	}
}

func setFlush(cards []card, h *hand) {
	copy(h.cards, cards)
	h.secondaryStrength = int(cards[4].value)*100000000 + int(cards[3].value)*1000000 + int(cards[2].value)*10000 + int(cards[1].value)*100 + int(cards[0].value)
}

func (h *hand7) isStraight() bool {
	consecutive := 0
	var i int
	var cards []card

	for i = len(h.cards) - 2; i >= 0 && consecutive < 5; i-- {
		if h.cards[i].value+1 == h.cards[i+1].value {
			cards = append(cards, h.cards[i+1])
			consecutive++
			if consecutive == 4 {
				cards = append(cards, h.cards[i])
				h.hand.cards = cards
				h.hand.secondaryStrength = int(h.cards[i+4].value)
				return true
				// treat case of A2345 as a special case
			} else if consecutive == 3 && h.cards[i].value == 2 && cards[0].value == 5 && h.cards[6].value == 14 {
				cards = append(cards, h.cards[i])
				cards = append(cards, h.cards[6])
				h.hand.cards = cards
				h.hand.secondaryStrength = 5
				return true
			}
		} else if h.cards[i].value == h.cards[i+1].value {
			// Do nothing
		} else {
			cards = []card{}
			consecutive = 0
		}
	}

	return false
}

func (h *hand7) isFullHouse() bool {
	cardsMap := make(map[byte][]card)
	for _, card := range h.cards {
		cardsMap[card.value] = append(cardsMap[card.value], card)
	}

	var threeValue byte
	var twoValue byte

	for _, cards := range cardsMap {
		if len(cards) == 3 {
			if threeValue > 0 {
				if cards[0].value > threeValue {
					twoValue = threeValue
					threeValue = cards[0].value
				} else if cards[0].value > twoValue {
					twoValue = cards[0].value
				}
			} else {
				threeValue = cards[0].value
			}
		}
		if len(cards) == 2 {
			if cards[0].value > twoValue {
				twoValue = cards[0].value
			}
		}
	}

	if threeValue > 0 && twoValue > 0 {
		h.hand.cards = append(cardsMap[threeValue], cardsMap[twoValue]...)
		h.hand.secondaryStrength = int(threeValue)*100 + int(twoValue)
		return true
	}

	return false
}

func (h *hand7) isThreeOfAKind() bool {
	cardsMap := make(map[byte][]card)
	for _, card := range h.cards {
		cardsMap[card.value] = append(cardsMap[card.value], card)
	}

	var threeValue byte
	var kicker1 byte
	var kicker2 byte

	for _, cards := range cardsMap {
		if len(cards) == 3 {
			if cards[0].value > threeValue {
				threeValue = cards[0].value
			}
		} else {
			if cards[0].value > kicker1 {
				kicker2 = kicker1
				kicker1 = cards[0].value
			} else if cards[0].value > kicker2 {
				kicker2 = cards[0].value
			}
		}
	}

	if threeValue > 0 {
		h.hand.cards = append(cardsMap[threeValue], cardsMap[kicker1][0], cardsMap[kicker2][0])
		h.hand.secondaryStrength = int(threeValue)*10000 + int(kicker1)*100 + int(kicker2)
		return true
	}

	return false
}

func (h *hand7) isTwoPair() bool {
	cardsMap := make(map[byte][]card)
	for _, card := range h.cards {
		cardsMap[card.value] = append(cardsMap[card.value], card)
	}

	var pair1 byte
	var pair2 byte
	var kicker byte

	for _, cards := range cardsMap {
		if len(cards) == 2 {
			if cards[0].value > pair1 {
				if pair2 > kicker {
					kicker = pair2
				}
				pair2 = pair1
				pair1 = cards[0].value
			} else if cards[0].value > pair2 {
				if pair2 > kicker {
					kicker = pair2
				}
				pair2 = cards[0].value
			} else if cards[0].value > kicker {
				kicker = cards[0].value
			}
		} else {
			if cards[0].value > kicker {
				kicker = cards[0].value
			}
		}
	}

	if pair1 > 0 && pair2 > 0 {
		h.hand.cards = append(cardsMap[pair1], cardsMap[pair2]...)
		h.hand.cards = append(h.hand.cards, cardsMap[kicker][0])
		h.hand.secondaryStrength = int(pair1)*10000 + int(pair2)*100 + int(kicker)
		return true
	}

	return false

}

func (h *hand7) isPair() bool {
	cardsMap := make(map[byte][]card)
	for _, card := range h.cards {
		cardsMap[card.value] = append(cardsMap[card.value], card)
	}

	var pair byte
	var kicker1 byte
	var kicker2 byte
	var kicker3 byte

	for _, cards := range cardsMap {
		if len(cards) == 2 {
			if cards[0].value > pair {
				pair = cards[0].value
			}
		} else {
			if cards[0].value > kicker1 {
				kicker3 = kicker2
				kicker2 = kicker1
				kicker1 = cards[0].value
			} else if cards[0].value > kicker2 {
				kicker3 = kicker2
				kicker2 = cards[0].value
			} else if cards[0].value > kicker3 {
				kicker3 = cards[0].value
			}
		}
	}

	if pair > 0 {
		h.hand.cards = cardsMap[pair]
		h.hand.cards = append(h.hand.cards, cardsMap[kicker1][0])
		h.hand.cards = append(h.hand.cards, cardsMap[kicker2][0])
		h.hand.cards = append(h.hand.cards, cardsMap[kicker3][0])
		h.hand.secondaryStrength = int(pair)*1000000 + int(kicker1)*10000 + int(kicker2)*100 + int(kicker3)
		return true
	}

	return false
}

func (h *hand7) HighCard() {
	h.hand.cards = h.cards[2:]
	h.hand.secondaryStrength = int(h.hand.cards[4].value)*100000000 +
		int(h.hand.cards[3].value)*1000000 +
		int(h.hand.cards[2].value)*10000 +
		int(h.hand.cards[2].value)*100 +
		int(h.hand.cards[0].value)
}

func Compare(handsData *HandsData, c []card) int {
	hand1 := hand7{cards: append(handsData.handsOdds[0].pockets, c...)}
	hand1.cards = append(hand1.cards, handsData.board...)

	hand2 := hand7{cards: append(handsData.handsOdds[1].pockets, c...)}
	hand2.cards = append(hand2.cards, handsData.board...)

	hand1.Strength(&handsData.handsOdds[0])
	hand2.Strength(&handsData.handsOdds[1])

	if hand1.hand.strength > hand2.hand.strength {
		return 1
	} else if hand1.hand.strength < hand2.hand.strength {
		return -1
	} else if hand1.hand.secondaryStrength > hand2.hand.secondaryStrength {
		return 1
	} else if hand1.hand.secondaryStrength < hand2.hand.secondaryStrength {
		return -1
	} else {
		return 0
	}

}

func (h *hand7) sort() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].value < h.cards[j].value
	})
}

func (h *hand7) sortSuit() {
	sort.Slice(h.cards, func(i, j int) bool {
		return h.cards[i].suit < h.cards[j].suit
	})
}
