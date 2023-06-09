package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	h1 := []card{
		{suit: Clubs, value: 3},
		{suit: Spades, value: 4},
	}

	h2 := []card{
		{suit: Hearts, value: 9},
		{suit: Diamonds, value: 10},
	}

	board := []card{
		//{suit: Hearts, value: 7},
		//{suit: Spades, value: 8},
		//{suit: Spades, value: 9},
		//{suit: Spades, value: 5},
	}

	o1, o2, tie := odds(h1, h2, board)

	fmt.Println(o1, o2, tie)

	fmt.Println("Time taken: ", time.Since(start))

}
