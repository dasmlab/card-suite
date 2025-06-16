package cribbage_core

import "math/rand"

type GameMode int

const (
    Mode1v1 GameMode = iota
    Mode3Way
    Mode2v2
    Mode3Teams
)

type GameState int

const (
    WaitingForPlayers GameState = iota
    Dealing
    DiscardToCrib
    Playing
    Scoring
    Finished
)

type Game struct {
    Mode         GameMode
    Players      []*Player
    CribOwnerIdx int
    Deck         []Card
    Crib         []Card
    Board        map[PlayerID]int // Peg positions
    State        GameState
    CurrentTurn  int
    Rng          *rand.Rand
    Starter      Card
}

