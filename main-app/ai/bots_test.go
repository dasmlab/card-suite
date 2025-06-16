package ai

import (
	"testing"
//	"math/rand"
	core "cribbage/core"
)

func makeTestGame() *core.Game {
	players := []*core.Player{
		{Name: "Bot1"}, {Name: "Bot2"},
	}
	g := &core.Game{
		Players: players,
		PlayTable: []core.Card{},
		PlayTotal: 0,
	}
	// Give each player two cards for pegging
	players[0].Hand = []core.Card{
		{Suit: core.Clubs, Rank: core.Seven},
		{Suit: core.Hearts, Rank: core.Eight},
	}
	players[1].Hand = []core.Card{
		{Suit: core.Spades, Rank: core.Seven},
		{Suit: core.Diamonds, Rank: core.Eight},
	}
	return g
}

func TestAllBotTypes_PlayCard(t *testing.T) {
	game := makeTestGame()
	botTypes := []BotSkillType{RandomBot, GreedyBot, BeginnerBot, ExpertBot, AdaptiveBot}
	for _, typ := range botTypes {
		bot := BotFactory(typ, typ.String())
		cardIdx := bot.PlayCard(game, 0)
		if cardIdx < -1 || cardIdx > 1 {
			t.Errorf("%s returned invalid card index: %d", typ.String(), cardIdx)
		}
	}
}

func TestRandomBot(t *testing.T) {
	game := makeTestGame()
	bot := BotFactory(RandomBot, "RandBot")
	cardIdx := bot.PlayCard(game, 0)
	if cardIdx != 0 && cardIdx != 1 {
		t.Errorf("RandomBot gave invalid index: %d", cardIdx)
	}
}

func TestGreedyBot(t *testing.T) {
	game := makeTestGame()
	bot := BotFactory(GreedyBot, "Greedy")
	cardIdx := bot.PlayCard(game, 0)
	if cardIdx < 0 || cardIdx > 1 {
		t.Errorf("GreedyBot gave invalid index: %d", cardIdx)
	}
}

func TestBeginnerBot(t *testing.T) {
	game := makeTestGame()
	bot := BotFactory(BeginnerBot, "Beginner")
	cardIdx := bot.PlayCard(game, 0)
	if cardIdx < 0 || cardIdx > 1 {
		t.Errorf("BeginnerBot gave invalid index: %d", cardIdx)
	}
}

func TestExpertBot(t *testing.T) {
	game := makeTestGame()
	bot := BotFactory(ExpertBot, "Expert")
	cardIdx := bot.PlayCard(game, 0)
	if cardIdx < 0 || cardIdx > 1 {
		t.Errorf("ExpertBot gave invalid index: %d", cardIdx)
	}
}

func TestAdaptiveBot(t *testing.T) {
	game := makeTestGame()
	bot := BotFactory(AdaptiveBot, "Adaptive")
	cardIdx := bot.PlayCard(game, 0)
	if cardIdx < 0 || cardIdx > 1 {
		t.Errorf("AdaptiveBot gave invalid index: %d", cardIdx)
	}
}

func TestNoLegalPlays(t *testing.T) {
	game := makeTestGame()
	game.PlayTotal = 31 // No card can be played
	bot := BotFactory(RandomBot, "RandBot")
	cardIdx := bot.PlayCard(game, 0)
	if cardIdx != -1 {
		t.Errorf("Bot should return -1 when no legal plays, got %d", cardIdx)
	}
}

