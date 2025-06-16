# Cribbage AI Bot Framework

This document outlines the pluggable "bot" (AI) framework for the Dasmlab Card Suite.

---

## Goals

- Allow any `Player` to use a bot/AI instead of human input.
- Support multiple bot “skill levels” and strategies (Random, Greedy, Beginner, Expert, Adaptive).
- Enable API/game creation to specify a bot type/skill for any player.
- Easily extend to more sophisticated AI or remote/networked agents in future milestones.

---

## Strategy Interface

All bots implement the `Strategy` interface:
```go
type Strategy interface {
    // Given a game and player index, return the index of the chosen card to play.
    ChooseCard(game *Game, playerIdx int) int
}
```

This interface may later be extended with more decision hooks (discard to crib, pegging, etc).

### Built-in Bot Types

| Type        | Behavior                                         |
|-------------|--------------------------------------------------|
| RandomBot   | Selects a random legal card.                     |
| GreedyBot   | Selects the play with maximum immediate score.   |
| BeginnerBot | Similar to Random, but avoids "worst" plays.     |
| ExpertBot   | Uses advanced cribbage heuristics (future).      |
| AdaptiveBot | Learns and adjusts skill dynamically.            |


You can assign any of these to a Player via the Strategy field.

### Assigning Bots
To add a bot to a player slot:

```go
p := &Player{
    Name: "Bot1",
    Strategy: NewBotBySkill("greedy"), // or "random", "expert", etc.
}
```
Use the helper:

```go
func NewBotBySkill(level string) Strategy
```
Example Usage in Game: 

```go
players[1].Strategy = NewBotBySkill("random")
players[2].Strategy = NewBotBySkill("adaptive")
If Strategy is nil or HumanStrategy, the API expects human input.
```

### Extending

To add new bots: implement the Strategy interface.

Adaptive/learning bots can track performance or game history.

You can later plug in remote/networked bots using the same interface.

### API Integration
You can expose /bot/skills to list available bots, or accept a bot_skill field in your API for game creation.

Dasmlab Card Suite, 2024 – Modular AI Bots for Card Games

