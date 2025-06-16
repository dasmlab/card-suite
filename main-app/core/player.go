package core

type Player struct {
    ID     PlayerID
    Name   string
    Hand   []Card
    Score  int
    IsBot  bool

    // 0.1.4 - Added for AIBot vs. Human Player
    Strategy Strategy // new:  'nil' means "human" unless assigned
}

