package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var wg sync.WaitGroup

type HandsData struct {
	handsOdds []HandOdds
	board     []card
}

type HandOdds struct {
	pockets []card
	mapOdds map[strength]int
}

func odds(handsData *HandsData) (float64, float64, float64) {
	deck := newDeck()
	deck = substruct(deck, handsData.handsOdds[0].pockets)
	deck = substruct(deck, handsData.handsOdds[1].pockets)
	deck = substruct(deck, handsData.board)
	m = sync.Mutex{}
	wg = sync.WaitGroup{}

	first := 0
	second := 0
	tie := 0

	for k := range handsData.handsOdds {
		m := map[strength]int{StraightFlush: 0, FourOfAKind: 0, FullHouse: 0, Flush: 0, Straight: 0, ThreeOfAKind: 0, TwoPair: 0, Pair: 0, HighCard: 0}
		handsData.handsOdds[k].mapOdds = m
	}

	first, second, tie = calcOdds(handsData, deck)

	size := first + second + tie

	o1 := float64(first) * 100 / float64(size)
	o2 := float64(second) * 100 / float64(size)
	otie := float64(tie) * 100 / float64(size)
	fmt.Printf("first hand: %.2f%% (%v), second hand: %.2f%% (%v), tie: %.2f%% (%v), total cases: %v\n", o1, first, o2, second, otie, tie, size)

	return o1, o2, otie
}

func calcOdds(handsData *HandsData, deck []card) (int, int, int) {
	first := 0
	second := 0
	tie := 0
	cards := []card{}
	nestedLoop(cards, deck, len(deck)-43, handsData, &first, &second, &tie)
	wg.Wait()

	return first, second, tie
}

// / nestedLoop is a recursive function that loops through all possible combinations of cards
func nestedLoop(cards []card, deck []card, i int, handsData *HandsData, first *int, second *int, tie *int) {
	// The stop condition - no more cards to draw (river card).
	if i == 0 {
		wg.Add(1)
		go func() {
			compareHands(handsData, cards, first, second, tie)
			wg.Done()
		}()

		return
	}

	// Loop through all cards in the deck.
	for j := len(deck) - 1; j >= 0; j-- {
		c := deck[j]
		deck = deck[:j]
		nestedLoop(append(cards, c), deck, i-1, handsData, first, second, tie)
	}
}

func compareHands(handsData *HandsData, c []card, first *int, second *int, tie *int) {
	result := Compare(handsData, c)

	if result == 1 {
		safeIncrement(first)
	} else if result == -1 {
		safeIncrement(second)
	} else {
		safeIncrement(tie)
	}
}

func safeIncrement(i *int) {
	m.Lock()
	defer m.Unlock()
	*i++
}

func substruct(deck []card, cards []card) []card {
	var result []card
	for _, c := range deck {
		if !contains(cards, c) {
			result = append(result, c)
		}
	}
	return result
}

func contains(cards []card, c card) bool {
	for _, card := range cards {
		if card == c {
			return true
		}
	}
	return false
}
