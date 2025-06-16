package core

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
    PlayTable   []Card     // Stack of played cards (since last reset)
    PlayHistory []PegAction // History of plays and "go"s
    PlayTotal   int        // Running total (<= 31)
    // New fields:
    Round        int
    GameOver     bool
    Winner       *Player
}


