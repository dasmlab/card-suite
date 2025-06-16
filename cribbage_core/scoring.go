package cribbage_core

import (
    "sort"
)

// ------- FIFTEENS -------
func scoreFifteens(hand []Card, starter Card) int {
    // Score 2 points for every unique combination of 2 or more cards that sum to 15
    all := append(hand[:], starter)
    n := len(all)
    count := 0

    var dfs func(idx int, sum int, picked int)
    dfs = func(idx int, sum int, picked int) {
        if idx == n {
            if picked >= 2 && sum == 15 {
                count++
            }
            return
        }
        // Include card at idx
        dfs(idx+1, sum+cardValue(all[idx]), picked+1)
        // Exclude card at idx
        dfs(idx+1, sum, picked)
    }
    dfs(0, 0, 0)
    return count * 2
}

func cardValue(c Card) int {
    if c.Rank >= 10 {
        return 10
    }
    return int(c.Rank)
}

// ------- PAIRS -------
func scorePairs(hand []Card, starter Card) int {
    // Score 2 points for each pair of cards of same rank
    all := append(hand[:], starter)
    points := 0
    for i := 0; i < len(all); i++ {
        for j := i + 1; j < len(all); j++ {
            if all[i].Rank == all[j].Rank {
                points += 2
            }
        }
    }
    return points
}

// ------- RUNS -------
func scoreRuns(hand []Card, starter Card) int {
    // Score for all unique runs of 3, 4, or 5 cards (but only longest if overlaps)
    all := append(hand[:], starter)
    maxRunLen := 0
    runCount := 0
    for runLen := 5; runLen >= 3; runLen-- {
        combs := combinations(all, runLen)
        found := 0
        for _, comb := range combs {
            if isRun(comb) {
                found++
            }
        }
        if found > 0 {
            maxRunLen = runLen
            runCount = found
            break // Only score the longest runs present
        }
    }
    return runCount * maxRunLen
}

func combinations(cards []Card, k int) [][]Card {
    // All possible k-card combinations (no repeats)
    var res [][]Card
    var comb func(start int, path []Card)
    comb = func(start int, path []Card) {
        if len(path) == k {
            tmp := make([]Card, k)
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        for i := start; i < len(cards); i++ {
            comb(i+1, append(path, cards[i]))
        }
    }
    comb(0, []Card{})
    return res
}

func isRun(cards []Card) bool {
    ranks := make([]int, len(cards))
    for i, c := range cards {
        ranks[i] = int(c.Rank)
    }
    sort.Ints(ranks)
    for i := 1; i < len(ranks); i++ {
        if ranks[i] != ranks[i-1]+1 {
            return false
        }
    }
    // No duplicate ranks allowed (e.g., 5-5-6)
    for i := 1; i < len(ranks); i++ {
        if ranks[i] == ranks[i-1] {
            return false
        }
    }
    return true
}

// ------- FLUSH -------
func scoreFlush(hand []Card, starter Card, isCrib bool) int {
    // Score 4 points if all hand cards same suit, +1 if starter matches (except crib, must be all 5)
    suit := hand[0].Suit
    for _, c := range hand[1:] {
        if c.Suit != suit {
            return 0
        }
    }
    if starter.Suit == suit {
        return 5
    } else if !isCrib {
        return 4
    }
    return 0
}

// ------- NOBS -------
func scoreNobs(hand []Card, starter Card) int {
    // 1 point for a Jack in hand that matches the suit of the starter
    for _, c := range hand {
        if c.Rank == Jack && c.Suit == starter.Suit {
            return 1
        }
    }
    return 0
}

// ------- HAND SCORING ENTRY POINT -------
func ScoreHand(hand []Card, starter Card, isCrib bool) int {
    return scoreFifteens(hand, starter) +
        scorePairs(hand, starter) +
        scoreRuns(hand, starter) +
        scoreFlush(hand, starter, isCrib) +
        scoreNobs(hand, starter)
}

