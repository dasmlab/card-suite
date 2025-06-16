package cribbage_core

type Suit string
type Rank int

const (
    Clubs    Suit = "C"
    Diamonds Suit = "D"
    Hearts   Suit = "H"
    Spades   Suit = "S"
)

const (
    Ace   Rank = 1
    Two   Rank = 2
    Three Rank = 3
    Four  Rank = 4
    Five  Rank = 5
    Six   Rank = 6
    Seven Rank = 7
    Eight Rank = 8
    Nine  Rank = 9
    Ten   Rank = 10
    Jack  Rank = 11
    Queen Rank = 12
    King  Rank = 13
)

type Card struct {
    Suit Suit
    Rank Rank
}

type PlayerID string

