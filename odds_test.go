package main

import (
	"fmt"
	"testing"
)

func Test_odds(t *testing.T) {
	h1 := []card{
		{suit: Spades, value: 12},
		{suit: Spades, value: 13},
	}

	h2 := []card{
		{suit: Hearts, value: 14},
		{suit: Spades, value: 14},
	}

	board := []card{
		// {suit: Hearts, value: 7},
		// {suit: Spades, value: 8},
		// {suit: Spades, value: 9},
		// {suit: Spades, value: 5},
	}

	handsData := HandsData{
		handsOdds: []HandOdds{{h1, map[strength]int{}}, {h2, map[strength]int{}}},
		board:     board,
	}

	// The expected values are the number of hands that have a given strength
	expexted_h1 := []int{294394, 733544, 392909, 78023, 73340, 94935, 41716, 2420, 1021, 2}
	expexted_h2 := []int{0, 613226, 674008, 209450, 20626, 29255, 149956, 15620, 117, 46}

	o1, o2, tie := odds(&handsData)

	fmt.Println(o1, o2, tie)

	handsCount := 0
	for _, v := range handsData.handsOdds[0].mapOdds {
		handsCount += v
	}

	fmt.Println("		First hand     Second hand")
	fmt.Println("_______________________________________")
	for k := len(handsData.handsOdds[0].mapOdds); k > 0; k-- {
		fmt.Printf("%v: %.02f\t\t%.02f\n", strength2String[strength(k)], float64(handsData.handsOdds[0].mapOdds[strength(k)])*100/float64(handsCount), float64(handsData.handsOdds[1].mapOdds[strength(k)])*100/float64(handsCount))

		if handsData.handsOdds[0].mapOdds[strength(k)] != expexted_h1[k-1] {
			t.Errorf("Expected %v, got %v", expexted_h1[k-1], handsData.handsOdds[0].mapOdds[strength(k)])
		}
		if handsData.handsOdds[1].mapOdds[strength(k)] != expexted_h2[k-1] {
			t.Errorf("Expected %v, got %v", expexted_h2[k-1], handsData.handsOdds[1].mapOdds[strength(k)])
		}
	}

}
