package core

// Strategy defines the interface for any AI or remote player agent.
type Strategy interface {
    // ChooseCard is called when the player must select a card to play (e.g., in pegging or discard)
    ChooseCard(game *Game, playerIdx int) int
    // (You can add more hooks later: e.g., OnDeal, OnScore, etc.)
}

// HumanStrategy is a trivial implementation that always returns -1 (invalid).
type HumanStrategy struct{}

// ChooseCard for HumanStrategy is not used (returns invalid index)
func (h *HumanStrategy) ChooseCard(game *Game, playerIdx int) int {
    return -1
}

