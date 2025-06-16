package ai

import (
	"cribbage/core"
	"math/rand"
)

// BotSkillType defines different classes/skills of AI bots.
type BotSkillType int

const (
	RandomBot BotSkillType = iota
	GreedyBot
	BeginnerBot
	ExpertBot
	AdaptiveBot
)

func (t BotSkillType) String() string {
	switch t {
	case RandomBot:
		return "RandomBot"
	case GreedyBot:
		return "GreedyBot"
	case BeginnerBot:
		return "BeginnerBot"
	case ExpertBot:
		return "ExpertBot"
	case AdaptiveBot:
		return "AdaptiveBot"
	default:
		return "Unknown"
	}
}

// Bot interface: all bots must implement PlayCard for pegging phase.
type Bot interface {
	Name() string
	PlayCard(game *core.Game, playerIdx int) (cardIdx int)
}

// BotFactory returns a bot of the requested skill type.
func BotFactory(skill BotSkillType, name string) Bot {
	switch skill {
	case RandomBot:
		return &RandomBotImpl{name}
	case GreedyBot:
		return &GreedyBotImpl{name}
	case BeginnerBot:
		return &BeginnerBotImpl{name}
	case ExpertBot:
		return &ExpertBotImpl{name}
	case AdaptiveBot:
		return &AdaptiveBotImpl{name}
	default:
		return &RandomBotImpl{name}
	}
}

// --- RandomBot: selects a random legal card ---
type RandomBotImpl struct{ botName string }

func (b *RandomBotImpl) Name() string { return b.botName }

func (b *RandomBotImpl) PlayCard(game *core.Game, playerIdx int) int {
	hand := game.Players[playerIdx].Hand
	legal := make([]int, 0)
	for i, card := range hand {
		if game.PlayTotal+core.CardValue(card) <= 31 {
			legal = append(legal, i)
		}
	}
	if len(legal) == 0 {
		return -1 // no legal play
	}
	return legal[rand.Intn(len(legal))]
}

// --- GreedyBot: play the card giving most immediate pegging points ---
type GreedyBotImpl struct{ botName string }

func (b *GreedyBotImpl) Name() string { return b.botName }

func (b *GreedyBotImpl) PlayCard(game *core.Game, playerIdx int) int {
	hand := game.Players[playerIdx].Hand
	bestIdx := -1
	maxPoints := -1
	for i, card := range hand {
		if game.PlayTotal+core.CardValue(card) > 31 {
			continue
		}
		// Simulate play, get points
		tempGame := *game // shallow copy is OK for simple pegging
		tempGame.PlayTable = append([]core.Card{}, game.PlayTable...)
		tempGame.PlayTotal = game.PlayTotal
		tempGame.PlayTable = append(tempGame.PlayTable, card)
		tempGame.PlayTotal += core.CardValue(card)
		points := tempGame.CalcPegPoints() // Use your core's pegging logic
		if points > maxPoints {
			bestIdx = i
			maxPoints = points
		}
	}
	return bestIdx
}

// --- BeginnerBot: like random, but avoids "worst" (obviously bad) plays ---
type BeginnerBotImpl struct{ botName string }

func (b *BeginnerBotImpl) Name() string { return b.botName }

func (b *BeginnerBotImpl) PlayCard(game *core.Game, playerIdx int) int {
	// For now: just use Random, can add heuristics later
	return BotFactory(RandomBot, b.botName).PlayCard(game, playerIdx)
}

// --- ExpertBot: placeholder for future advanced strategy ---
type ExpertBotImpl struct{ botName string }
func (b *ExpertBotImpl) Name() string { return b.botName }
func (b *ExpertBotImpl) PlayCard(game *core.Game, playerIdx int) int {
	return BotFactory(GreedyBot, b.botName).PlayCard(game, playerIdx) // stub: use Greedy
}

// --- AdaptiveBot: experimental stub ---
type AdaptiveBotImpl struct{ botName string }
func (b *AdaptiveBotImpl) Name() string { return b.botName }
func (b *AdaptiveBotImpl) PlayCard(game *core.Game, playerIdx int) int {
	// In the future: adjust skill level dynamically
	return BotFactory(BeginnerBot, b.botName).PlayCard(game, playerIdx)
}

