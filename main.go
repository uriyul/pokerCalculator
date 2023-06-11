package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	Foo()

	fmt.Println("Time taken: ", time.Since(start))

}

func Foo() {
	h1 := []card{
		{suit: Clubs, value: 3},
		{suit: Spades, value: 4},
	}

	h2 := []card{
		{suit: Hearts, value: 9},
		{suit: Diamonds, value: 10},
	}

	board := []card{
		{suit: Hearts, value: 7},
		{suit: Spades, value: 8},
		{suit: Spades, value: 9},
		// {suit: Spades, value: 5},
	}

	handsData := HandsData{
		handsOdds: []HandOdds{{h1, map[strength]int{}}, {h2, map[strength]int{}}},
		board:     board,
	}

	fmt.Println("h1: ", h1, " board: ", board)

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
	}
}
