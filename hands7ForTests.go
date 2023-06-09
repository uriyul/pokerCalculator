package main

type testHands7 struct {
	handHighCard       hand7
	handPair           hand7
	handHighPair       hand7
	handTwoPair        hand7
	handThreeOfAKind   hand7
	handStraight       hand7
	handFlush          hand7
	handFullHouse      hand7
	handFourOfAKind    hand7
	handStraightFlush  hand7
	handRoyalFlush     hand7
	handStraightA5     hand7
	handTwoPairAAKKQQ5 hand7
	handTwoPairAAKK789 hand7
}

type fourOfAKindHands7Test struct {
	kickerAce    hand7
	kicker10     hand7
	fullHouse    hand7
	staightFlush hand7
}

func (foak *fourOfAKindHands7Test) Init() {
	foak.kickerAce = hand7{
		cards: []card{
			{value: 2, suit: Spades},
			{value: 2, suit: Hearts},
			{value: 2, suit: Diamonds},
			{value: 2, suit: Clubs},
			{value: 3, suit: Hearts},
			{value: 14, suit: Hearts},
			{value: 5, suit: Hearts},
		},
	}
	foak.kicker10 = hand7{
		cards: []card{
			{value: 2, suit: Spades},
			{value: 2, suit: Hearts},
			{value: 2, suit: Diamonds},
			{value: 2, suit: Clubs},
			{value: 3, suit: Hearts},
			{value: 10, suit: Hearts},
			{value: 5, suit: Hearts},
		},
	}
	foak.fullHouse = hand7{
		cards: []card{
			{value: 2, suit: Spades},
			{value: 2, suit: Hearts},
			{value: 2, suit: Diamonds},
			{value: 2, suit: Clubs},
			{value: 3, suit: Hearts},
			{value: 4, suit: Hearts},
			{value: 4, suit: Clubs},
		},
	}
	foak.staightFlush = hand7{
		cards: []card{
			{value: 2, suit: Spades},
			{value: 3, suit: Spades},
			{value: 4, suit: Spades},
			{value: 5, suit: Spades},
			{value: 6, suit: Spades},
			{value: 14, suit: Hearts},
			{value: 5, suit: Hearts},
		},
	}
}

func (th *testHands7) Init() {

	th.handHighCard =
		hand7{
			cards: []card{
				{value: 2, suit: Spades},
				{value: 3, suit: Hearts},
				{value: 6, suit: Hearts},
				{value: 8, suit: Hearts},
				{value: 12, suit: Hearts},
				{value: 13, suit: Diamonds},
				{value: 14, suit: Clubs},
			},
			hand: hand{},
		}
	th.handPair = hand7{
		cards: []card{
			{value: 7, suit: Spades},
			{value: 7, suit: Hearts},
			{value: 4, suit: Hearts},
			{value: 9, suit: Hearts},
			{value: 12, suit: Hearts},
			{value: 13, suit: Diamonds},
			{value: 14, suit: Clubs},
		},
	}
	th.handHighPair = hand7{
		cards: []card{
			{value: 10, suit: Spades},
			{value: 10, suit: Hearts},
			{value: 4, suit: Hearts},
			{value: 9, suit: Hearts},
			{value: 6, suit: Hearts},
			{value: 6, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handTwoPair = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 5, suit: Hearts},
			{value: 5, suit: Hearts},
			{value: 6, suit: Diamonds},
			{value: 9, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handThreeOfAKind = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 3, suit: Diamonds},
			{value: 5, suit: Hearts},
			{value: 6, suit: Diamonds},
			{value: 9, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handStraight = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Hearts},
			{value: 5, suit: Diamonds},
			{value: 6, suit: Hearts},
			{value: 7, suit: Diamonds},
			{value: 3, suit: Clubs},
			{value: 4, suit: Clubs},
		},
	}
	th.handFlush = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Spades},
			{value: 5, suit: Spades},
			{value: 12, suit: Spades},
			{value: 7, suit: Spades},
			{value: 6, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handFullHouse = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 3, suit: Diamonds},
			{value: 5, suit: Hearts},
			{value: 5, suit: Diamonds},
			{value: 6, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handFourOfAKind = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 3, suit: Hearts},
			{value: 3, suit: Diamonds},
			{value: 3, suit: Clubs},
			{value: 5, suit: Diamonds},
			{value: 6, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handStraightFlush = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Spades},
			{value: 5, suit: Spades},
			{value: 6, suit: Spades},
			{value: 7, suit: Spades},
			{value: 6, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handRoyalFlush = hand7{
		cards: []card{
			{value: 10, suit: Spades},
			{value: 11, suit: Spades},
			{value: 12, suit: Spades},
			{value: 13, suit: Spades},
			{value: 14, suit: Spades},
			{value: 6, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
	th.handStraightA5 = hand7{
		cards: []card{
			{value: 3, suit: Spades},
			{value: 4, suit: Hearts},
			{value: 5, suit: Diamonds},
			{value: 2, suit: Hearts},
			{value: 14, suit: Diamonds},
			{value: 3, suit: Clubs},
			{value: 4, suit: Clubs},
		},
	}
	th.handTwoPairAAKKQQ5 = hand7{
		cards: []card{
			{value: 14, suit: Spades},
			{value: 14, suit: Hearts},
			{value: 13, suit: Diamonds},
			{value: 13, suit: Hearts},
			{value: 12, suit: Diamonds},
			{value: 12, suit: Clubs},
			{value: 4, suit: Clubs},
		},
	}
	th.handTwoPairAAKK789 = hand7{
		cards: []card{
			{value: 14, suit: Spades},
			{value: 14, suit: Hearts},
			{value: 13, suit: Diamonds},
			{value: 13, suit: Hearts},
			{value: 6, suit: Diamonds},
			{value: 7, suit: Clubs},
			{value: 8, suit: Clubs},
		},
	}
}
