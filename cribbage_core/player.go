package cribbage_core

type Player struct {
    ID     PlayerID
    Name   string
    Hand   []Card
    Score  int
    IsBot  bool
}

