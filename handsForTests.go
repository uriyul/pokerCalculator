package main

type testHands struct {
	handHighCard      hand
	handPair          hand
	handHighPair      hand
	handTwoPair       hand
	handThreeOfAKind  hand
	handStraight      hand
	handFlush         hand
	handFullHouse     hand
	handFourOfAKind   hand
	handStraightFlush hand
	handRoyalFlush    hand
}

func (th *testHands) Init() {

	th.handHighCard = hand{
		cards: []card{
			{value: 2, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 6, suit: Hearts},
			{value: 8, suit: Hearts},
			{value: 12, suit: Hearts},
		},
	}
	th.handPair = hand{
		cards: []card{
			{value: 7, suit: Spades},
			{value: 7, suit: Hearts},
			{value: 4, suit: Hearts},
			{value: 9, suit: Hearts},
			{value: 12, suit: Hearts},
		},
	}
	th.handHighPair = hand{
		cards: []card{
			{value: 10, suit: Spades},
			{value: 10, suit: Hearts},
			{value: 4, suit: Hearts},
			{value: 9, suit: Hearts},
			{value: 6, suit: Hearts},
		},
	}
	th.handTwoPair = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 5, suit: Hearts},
			{value: 5, suit: Hearts},
			{value: 6, suit: Diamonds},
		},
	}
	th.handThreeOfAKind = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 3, suit: Diamonds},
			{value: 5, suit: Hearts},
			{value: 6, suit: Diamonds},
		},
	}
	th.handStraight = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Hearts},
			{value: 5, suit: Diamonds},
			{value: 6, suit: Hearts},
			{value: 7, suit: Diamonds},
		},
	}
	th.handFlush = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Spades},
			{value: 5, suit: Spades},
			{value: 12, suit: Spades},
			{value: 7, suit: Spades},
		},
	}
	th.handFullHouse = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 3, suit: Diamonds},
			{value: 5, suit: Hearts},
			{value: 5, suit: Diamonds},
		},
	}
	th.handFourOfAKind = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 3, suit: Diamonds},
			{value: 3, suit: Clubs},
			{value: 5, suit: Diamonds},
		},
	}
	th.handStraightFlush = hand{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Spades},
			{value: 5, suit: Spades},
			{value: 6, suit: Spades},
			{value: 7, suit: Spades},
		},
	}
	th.handRoyalFlush = hand{
		cards: []card{
			{value: 10, suit: Spades},
			{value: 11, suit: Spades},
			{value: 12, suit: Spades},
			{value: 13, suit: Spades},
			{value: 14, suit: Spades},
		},
	}
}
