package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var wg sync.WaitGroup

func odds(h1 []card, h2 []card, board []card) (float64, float64, float64) {
	deck := newDeck()
	deck = substruct(deck, h1)
	deck = substruct(deck, h2)
	deck = substruct(deck, board)
	m = sync.Mutex{}
	wg = sync.WaitGroup{}

	first := 0
	second := 0
	tie := 0
	h1 = append(h1, board...)
	h2 = append(h2, board...)

	new := true

	if new {
		first, second, tie = odds_option2(h1, h2, deck)
	} else {
		first, second, tie = odds_option1(h1, h2, deck)
	}

	size := first + second + tie

	o1 := float64(first) * 100 / float64(size)
	o2 := float64(second) * 100 / float64(size)
	otie := float64(tie) * 100 / float64(size)
	fmt.Printf("first hand: %.2f%% (%v), second hand: %.2f%% (%v), tie: %.2f%% (%v), total cases: %v\n", o1, first, o2, second, otie, tie, size)

	return o1, o2, otie
}

func odds_option2(h1 []card, h2 []card, deck []card) (int, int, int) {
	fmt.Println("odds_option2")
	first := 0
	second := 0
	tie := 0
	cards := []card{}
	nestedLoop(cards, deck, len(deck)-43, h1, h2, &first, &second, &tie)
	wg.Wait()

	return first, second, tie
}

func nestedLoop(cards []card, deck []card, i int, h1 []card, h2 []card, first *int, second *int, tie *int) {
	if i == 0 {
		wg.Add(1)
		go func() {
			compareHands(h1, h2, cards, first, second, tie)
			wg.Done()
		}()

		return
	}

	for j := len(deck) - 1; j >= 0; j-- {
		c := deck[j]
		deck = deck[:j]
		nestedLoop(append(cards, c), deck, i-1, h1, h2, first, second, tie)
	}
}

func odds_option1(h1 []card, h2 []card, deck []card) (int, int, int) {
	fmt.Println("odds_option1")
	missingCardsCount := 7 - len(h1)
	combinations := getCombinations(deck, missingCardsCount)
	wg.Add(len(combinations))

	size := len(combinations)
	first := 0
	second := 0
	tie := 0

	fmt.Printf("number of combinations: %v\n", size)
	for _, c := range combinations {
		go func(c []card) {
			compareHands(h1, h2, c, &first, &second, &tie)
			wg.Done()
		}(c)
	}
	wg.Wait()

	return first, second, tie
}

func compareHands(h1 []card, h2 []card, c []card, first *int, second *int, tie *int) {
	hand1 := hand7{cards: append(h1, c...)}
	hand2 := hand7{cards: append(h2, c...)}
	result := hand1.Compare(&hand2)

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

func getCombinations(deck []card, missingCardsCount int) [][]card {

	combinations := [][]card{}

	for j := 0; j < len(deck); j++ {
		combinations = append(combinations, []card{deck[j]})
	}

	for i := 1; i < missingCardsCount; i++ {
		combinations = combine(combinations, deck)
	}

	return combinations
}

func combine(combinations [][]card, cards []card) [][]card {
	for _, c := range combinations {
		for _, c2 := range cards {
			if !contains(c, c2) {
				a := append(c, c2)
				combinations = append(combinations, a)
			}
		}
		combinations = combinations[1:]
	}
	return combinations
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
