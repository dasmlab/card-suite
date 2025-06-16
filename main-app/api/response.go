package api

import "cribbage/core"

type DealResponse struct {
    Players []*core.Player `json:"players"`
    Starter core.Card     `json:"starter"`
    State   string        `json:"state"`
}

type StatusResponse struct {
    Players     []*core.Player `json:"players"`
    Crib        []core.Card   `json:"crib"`
    Starter     core.Card     `json:"starter"`
    State       core.GameState `json:"state"`
    Turn        int           `json:"turn"`
    Dealer      int           `json:"dealer"`
}

