package core

import "math/rand"

func NewDeck() []Card {
    deck := make([]Card, 0, 52)
    for _, suit := range []Suit{Clubs, Diamonds, Hearts, Spades} {
        for r := Ace; r <= King; r++ {
            deck = append(deck, Card{Suit: suit, Rank: r})
        }
    }
    return deck
}

func Shuffle(deck []Card, rng *rand.Rand) {
    n := len(deck)
    for i := n - 1; i > 0; i-- {
        j := rng.Intn(i + 1)
        deck[i], deck[j] = deck[j], deck[i]
    }
}

