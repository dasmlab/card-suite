package cribbage_core

import (
    "math/rand"
    "testing"

    log "github.com/sirupsen/logrus"
)

func TestDealAndDiscardToCrib(t *testing.T) {
    // Setup Logrus to always show info for this test
    log.SetLevel(log.InfoLevel)

    // Create deterministic RNG for repeatable tests
    rng := rand.New(rand.NewSource(42))
    playerNames := []string{"Alice", "Bob"}
    g := NewGame(Mode1v1, playerNames, rng)

    log.Info("Starting Deal phase test")
    if err := g.Deal(); err != nil {
        log.Errorf("Deal failed: %v", err)
        t.Fatalf("Deal failed: %v", err)
    }

    // Check that each player has the right number of cards after deal
    for _, p := range g.Players {
        log.Infof("Player %s hand after deal: %v", p.Name, p.Hand)
        if len(p.Hand) != 6 {
            t.Errorf("Player %s hand length = %d, want 6", p.Name, len(p.Hand))
        }
    }

    // Check that the starter card is set
    log.Infof("Starter card: %v", g.Starter)
    if g.Starter == (Card{}) {
        t.Error("Starter card not set")
    }

    // Simulate discards: Alice discards first two, Bob discards last two
    log.Info("Alice discards cards [0,1] to crib")
    if err := g.DiscardToCrib(0, []int{0, 1}); err != nil {
        log.Errorf("Alice discard failed: %v", err)
        t.Fatalf("Alice discard failed: %v", err)
    }
    log.Infof("Alice hand after discard: %v", g.Players[0].Hand)
    log.Infof("Crib after Alice: %v", g.Crib)

    log.Info("Bob discards cards [4,5] to crib")
    if err := g.DiscardToCrib(1, []int{4, 5}); err != nil {
        log.Errorf("Bob discard failed: %v", err)
        t.Fatalf("Bob discard failed: %v", err)
    }
    log.Infof("Bob hand after discard: %v", g.Players[1].Hand)
    log.Infof("Crib after Bob: %v", g.Crib)

    // After both players discard, game should move to Playing state
    if g.State != Playing {
        log.Errorf("Game state after discards = %v, want Playing", g.State)
        t.Errorf("Game state after discards = %v, want Playing", g.State)
    } else {
        log.Info("Game state is now Playing (as expected)")
    }

    // Crib should have 4 cards, each hand should have 4
    if len(g.Crib) != 4 {
        log.Errorf("Crib has %d cards, want 4", len(g.Crib))
        t.Errorf("Crib has %d cards, want 4", len(g.Crib))
    }
    for i, p := range g.Players {
        if len(p.Hand) != 4 {
            log.Errorf("Player %d (%s) has %d cards, want 4 after discards", i, p.Name, len(p.Hand))
            t.Errorf("Player %d (%s) has %d cards, want 4 after discards", i, p.Name, len(p.Hand))
        }
    }

    log.Info("Deal/DiscardToCrib integration test completed successfully")
}


func TestPeggingPlayPhase(t *testing.T) {
    log.SetLevel(log.InfoLevel)
    rng := rand.New(rand.NewSource(17))
    playerNames := []string{"Alice", "Bob"}
    g := NewGame(Mode1v1, playerNames, rng)

    // Force hands and starter for deterministic test
    // Alice: 7♣, 8♣, Bob: 7♥, 8♥, Starter: King♣
    g.Players[0].Hand = []Card{
        {Suit: Clubs, Rank: Seven},
        {Suit: Clubs, Rank: Eight},
    }
    g.Players[1].Hand = []Card{
        {Suit: Hearts, Rank: Seven},
        {Suit: Hearts, Rank: Eight},
    }
    g.Starter = Card{Suit: Clubs, Rank: King}
    g.State = Playing
    g.PlayTable = []Card{}
    g.PlayHistory = []PegAction{}
    g.PlayTotal = 0
    g.CurrentTurn = 0

    log.Info("Begin pegging phase - initial hands set")
    log.Infof("Alice hand: %v", g.Players[0].Hand)
    log.Infof("Bob hand: %v", g.Players[1].Hand)

    // Turn 1: Alice plays 7♣
    if err := g.PlayCard(0, 0); err != nil {
        t.Fatalf("Alice PlayCard failed: %v", err)
    }
    log.Infof("Peg: Alice plays 7♣, table: %v, total: %d, score: %d", g.PlayTable, g.PlayTotal, g.Players[0].Score)

    // Turn 2: Bob plays 7♥ (pair! should score 2)
    if err := g.PlayCard(1, 0); err != nil {
        t.Fatalf("Bob PlayCard failed: %v", err)
    }
    log.Infof("Peg: Bob plays 7♥, table: %v, total: %d, score: %d", g.PlayTable, g.PlayTotal, g.Players[1].Score)
    if g.Players[1].Score != 2 {
        t.Errorf("Bob did not get 2 points for pair: got %d", g.Players[1].Score)
    }

    // Turn 3: Alice plays 8♣ (total is 7+7+8 = 22)
    if err := g.PlayCard(0, 0); err != nil {
        t.Fatalf("Alice PlayCard failed: %v", err)
    }
    log.Infof("Peg: Alice plays 8♣, table: %v, total: %d, score: %d", g.PlayTable, g.PlayTotal, g.Players[0].Score)

    // Turn 4: Bob plays 8♥ (should be pair for 2, and run of 4 for 4 = 6 points)
    if err := g.PlayCard(1, 0); err != nil {
        t.Fatalf("Bob PlayCard failed: %v", err)
    }
    log.Infof("Peg: Bob plays 8♥, table: %v, total: %d, score: %d", g.PlayTable, g.PlayTotal, g.Players[1].Score)
    // 2 for pair + 4 for run of 4, total 6 for this play (+2 previous = 8)
    if g.Players[1].Score != 8 {
        t.Errorf("Bob should have 8 points total after run and pair, got %d", g.Players[1].Score)
    }

    // Both hands empty = pegging phase ends
    if len(g.Players[0].Hand) != 0 || len(g.Players[1].Hand) != 0 {
        t.Errorf("Players' hands not empty after pegging")
    }

    log.Info("Pegging phase test completed successfully")
}

