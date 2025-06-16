package cribbage_core

import (
    "errors"
    "math/rand"
    "time"
)


// -- Create a new game and player structs --
func NewGame(mode GameMode, playerNames []string, rng *rand.Rand) *Game {
    players := make([]*Player, len(playerNames))
    for i, name := range playerNames {
        players[i] = &Player{
            ID:    PlayerID(name), // Use UUID for real IDs in prod
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

// -- Deal hands to all players, select starter card --
func (g *Game) Deal() error {
    if len(g.Players) < 2 {
        return errors.New("not enough players")
    }
    g.Deck = NewDeck()
    if g.Rng == nil {
        g.Rng = rand.New(rand.NewSource(time.Now().UnixNano()))
    }
    Shuffle(g.Deck, g.Rng)
    numCards := 6
    if len(g.Players) > 2 {
        numCards = 5
    }

    // Deal hands
    idx := 0
    for _, p := range g.Players {
        p.Hand = make([]Card, numCards)
        for j := 0; j < numCards; j++ {
            p.Hand[j] = g.Deck[idx]
            idx++
        }
    }

    // Select starter
    g.Starter = g.Deck[idx]
    idx++

    // Update deck to remove dealt/starter cards (optional, can ignore)
    g.Deck = g.Deck[idx:]

    // Clear crib, update state
    g.Crib = []Card{}
    g.State = DiscardToCrib
    g.CurrentTurn = g.CribOwnerIdx
    return nil
}

// -- Each player discards cards to the crib (must discard 2 cards) --
func (g *Game) DiscardToCrib(playerIdx int, cardIdxs []int) error {
    if g.State != DiscardToCrib {
        return errors.New("not in discard phase")
    }
    player := g.Players[playerIdx]
    if len(cardIdxs) != 2 {
        return errors.New("must discard exactly 2 cards")
    }
    // Remove by highest index first to avoid reindexing problems
    if cardIdxs[0] < cardIdxs[1] {
        cardIdxs[0], cardIdxs[1] = cardIdxs[1], cardIdxs[0]
    }
    for _, idx := range cardIdxs {
        if idx < 0 || idx >= len(player.Hand) {
            return errors.New("invalid card index")
        }
        g.Crib = append(g.Crib, player.Hand[idx])
        player.Hand = append(player.Hand[:idx], player.Hand[idx+1:]...)
    }
    // Optionally: Transition to play phase when all players have discarded
    expectedCribCards := len(g.Players) * 2
    if len(g.Crib) == expectedCribCards {
        g.State = Playing
    }
    return nil
}


// -- Score hands and crib at end of round --
func (g *Game) ScoreRound() error {
    for _, p := range g.Players {
        score := ScoreHand(p.Hand, g.Starter, false)
        p.Score += score
    }
    cribOwner := g.Players[g.CribOwnerIdx]
    cribScore := ScoreHand(g.Crib, g.Starter, true)
    cribOwner.Score += cribScore
    g.State = Finished
    return nil
}

// -- Advance to the next player's turn (round robin) --
func (g *Game) NextTurn() {
    g.CurrentTurn = (g.CurrentTurn + 1) % len(g.Players)
}

// -- Is the game over? --
func (g *Game) IsGameOver() bool {
    for _, p := range g.Players {
        if p.Score >= 121 {
            return true
        }
    }
    return false
}

// PegAction represents a single play or "go"
type PegAction struct {
    PlayerIdx int
    Card      *Card // nil means "go"
    TableIdx  int   // index in play stack (for possible scoring)
}

// PlayCard plays a card during pegging. Returns error if invalid move.
func (g *Game) PlayCard(playerIdx int, cardIdx int) error {
    if g.State != Playing {
        return errors.New("not in play phase")
    }
    player := g.Players[playerIdx]
    if playerIdx != g.CurrentTurn {
        return errors.New("not this player's turn")
    }
    if cardIdx < 0 || cardIdx >= len(player.Hand) {
        return errors.New("invalid card index")
    }

    card := player.Hand[cardIdx]
    // Check if card can be played (total must not exceed 31)
    if g.PlayTotal+cardValue(card) > 31 {
        return errors.New("cannot play card: would exceed 31")
    }

    // Play card: add to table, history, remove from hand
    g.PlayTable = append(g.PlayTable, card)
    g.PlayHistory = append(g.PlayHistory, PegAction{PlayerIdx: playerIdx, Card: &card, TableIdx: len(g.PlayTable)-1})
    g.PlayTotal += cardValue(card)
    player.Hand = append(player.Hand[:cardIdx], player.Hand[cardIdx+1:]...)

    // Pegging points!
    points := g.calcPegPoints()
    player.Score += points

    // If exactly 31, award 2 points and reset table
    if g.PlayTotal == 31 {
        player.Score += 2
        g.resetPegTable(playerIdx)
    } else if g.allPlayersCannotPlay() {
        // "Go" - award 1 point to last player to play, reset table
        player.Score += 1
        g.resetPegTable(playerIdx)
    } else {
        g.NextTurn()
    }

    return nil
}

// GoAction: Player says "go" when they can't play
func (g *Game) Go(playerIdx int) error {
    if g.State != Playing {
        return errors.New("not in play phase")
    }
    if playerIdx != g.CurrentTurn {
        return errors.New("not this player's turn")
    }
    g.PlayHistory = append(g.PlayHistory, PegAction{PlayerIdx: playerIdx, Card: nil, TableIdx: len(g.PlayTable)})
    g.NextTurn()
    return nil
}

// Helper: Calculate pegging points for most recent play
func (g *Game) calcPegPoints() int {
    n := len(g.PlayTable)
    if n == 0 {
        return 0
    }
    points := 0

    // Check for 15 (2 points)
    sum := 0
    for i := n - 1; i >= 0; i-- {
        sum += cardValue(g.PlayTable[i])
        if sum >= 15 {
            break
        }
    }
    if sum == 15 {
        points += 2
    }

    // Check for pairs/trips/quads (from end)
    pairCount := 1
    for i := n - 2; i >= 0 && g.PlayTable[i].Rank == g.PlayTable[n-1].Rank; i-- {
        pairCount++
    }
    switch pairCount {
    case 2:
        points += 2
    case 3:
        points += 6
    case 4:
        points += 12
    }

    // Check for runs of 3, 4, 5 (from end)
    for l := 5; l >= 3; l-- {
        if n < l {
            continue
        }
        if isRun(g.PlayTable[n-l:]) {
            points += l
            break
        }
    }
    return points
}

// Helper: Can all players not play (must "go")?
func (g *Game) allPlayersCannotPlay() bool {
    for _, p := range g.Players {
        if len(p.Hand) == 0 {
            continue
        }
        // If any card can be played, return false
        for _, c := range p.Hand {
            if g.PlayTotal+cardValue(c) <= 31 {
                return false
            }
        }
    }
    return true
}

// Helper: Reset the pegging table for next run
func (g *Game) resetPegTable(lastPlayerIdx int) {
    g.PlayTable = []Card{}
    g.PlayTotal = 0
    g.CurrentTurn = lastPlayerIdx
    // Optionally: mark hands empty as round finished, etc.
}

