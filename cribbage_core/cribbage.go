package cribbage_core

import (
	"errors"
	"math/rand"
)

// Assumes GameMode, GameState, Player, Card, etc. are defined elsewhere

func NewGame(mode GameMode, playerNames []string, rng *rand.Rand) *Game {
	players := make([]*Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = &Player{
			ID:    PlayerID(name), // or use UUID for real IDs
			Name:  name,
			Score: 0,
		}
	}
	g := &Game{
		Mode:         mode,
		Players:      players,
		CribOwnerIdx: 0,
		Board:        make(map[PlayerID]int),
		Rng:          rng,
		State:        WaitingForPlayers,
	}
	return g
}

func (g *Game) Deal() error {
	g.Deck = NewDeck()
	Shuffle(g.Deck, g.Rng)
	numCards := 6 // Standard for 2 players
	if len(g.Players) > 2 {
		numCards = 5 // Standard for 3+ players
	}
	for i, p := range g.Players {
		p.Hand = make([]Card, numCards)
		for j := 0; j < numCards; j++ {
			p.Hand[j] = g.Deck[i*numCards+j]
		}
	}
	g.Starter = g.Deck[len(g.Players)*numCards] // Next card is starter
	g.State = DiscardToCrib
	return nil
}

func (g *Game) DiscardToCrib(playerIdx int, cardIdxs []int) error {
	if g.State != DiscardToCrib {
		return errors.New("not in discard phase")
	}
	player := g.Players[playerIdx]
	if len(cardIdxs) != 2 {
		return errors.New("must discard exactly 2 cards")
	}
	// Remove selected cards and add to crib
	// Always remove highest index first!
	for i := 1; i >= 0; i-- {
		idx := cardIdxs[i]
		g.Crib = append(g.Crib, player.Hand[idx])
		player.Hand = append(player.Hand[:idx], player.Hand[idx+1:]...)
	}
	// Optionally: check if all players have discarded
	// Move to Play phase if so
	return nil
}

func (g *Game) PlayCard(playerIdx int, cardIdx int) error {
	if g.State != Playing {
		return errors.New("not in play phase")
	}
	player := g.Players[playerIdx]
	if cardIdx < 0 || cardIdx >= len(player.Hand) {
		return errors.New("invalid card index")
	}
	// Play card (for now, just remove it from hand)
	card := player.Hand[cardIdx]
	player.Hand = append(player.Hand[:cardIdx], player.Hand[cardIdx+1:]...)
	// TODO: Add to play history for pegging
	_ = card // temp - supress compile error for now
	// Advance turn, update state as needed
	g.NextTurn()
	return nil
}

func (g *Game) ScoreRound() error {
	// Each player scores their hand + starter
	for _, p := range g.Players {
		score := ScoreHand(p.Hand, g.Starter, false)
		p.Score += score
	}
	// Crib owner scores the crib
	cribOwner := g.Players[g.CribOwnerIdx]
	cribScore := ScoreHand(g.Crib, g.Starter, true)
	cribOwner.Score += cribScore
	g.State = Finished
	return nil
}

func (g *Game) NextTurn() {
	// Naive round-robin
	g.CurrentTurn = (g.CurrentTurn + 1) % len(g.Players)
	// TODO: End round/game if all cards played
}

func (g *Game) IsGameOver() bool {
	// Simple: first to 121 points wins
	for _, p := range g.Players {
		if p.Score >= 121 {
			return true
		}
	}
	return false
}

