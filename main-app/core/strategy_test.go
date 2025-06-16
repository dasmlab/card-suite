package core

import (
    "testing"
)

func TestHumanStrategyChooseCard(t *testing.T) {
    g := &Game{}
    h := &HumanStrategy{}
    idx := h.ChooseCard(g, 0)
    if idx != -1 {
        t.Errorf("HumanStrategy.ChooseCard() = %d, want -1", idx)
    }
}

