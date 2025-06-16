package ai

import (
    core "cribbage/core"
)

// SimpleBot always selects the first legal card to play
type SimpleBot struct{}

func (b *SimpleBot) ChooseCard(game *core.Game, hand []core.Card, playerIdx int) int {
    for i, card := range hand {
        if game.IsLegalPegPlay(playerIdx, card) {
            return i
        }
    }
    return -1 // "Go"
}

