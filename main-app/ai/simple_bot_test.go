package ai

import (
    "testing"
    core "cribbage/core"
)

func TestSimpleBot_FirstLegal(t *testing.T) {
    // Make a game where only the second card is legal for pegging
    hand := []core.Card{
        {core.Hearts, core.Ten},   // not legal: would bust 31
        {core.Spades, core.Five},  // legal
    }
    starter := core.Card{core.Diamonds, core.Four}
    // Each player gets a hand (use only one for this bot test)
    game := core.NewTestPegGame([][]core.Card{hand}, starter)
    game.PlayTotal = 26 // Only a 5 or less is legal
    bot := &SimpleBot{}
    idx := bot.ChooseCard(game, hand, 0)
    if idx != 1 {
        t.Errorf("SimpleBot should pick index 1, got %d", idx)
    }
}

func TestSimpleBot_NoLegal(t *testing.T) {
    hand := []core.Card{
        {core.Hearts, core.King},
        {core.Spades, core.Queen},
    }
    starter := core.Card{core.Diamonds, core.Four}
    game := core.NewTestPegGame([][]core.Card{hand}, starter)
    game.PlayTotal = 30 // Can't play either card
    bot := &SimpleBot{}
    idx := bot.ChooseCard(game, hand, 0)
    if idx != -1 {
        t.Errorf("SimpleBot should return -1 for 'Go', got %d", idx)
    }
}

