package cribbage_core

import ( 
	"testing"
	log "github.com/sirupsen/logrus"
)

func TestPerfectHand(t *testing.T) {
    // Define the hand (a perfet hand in crib is 29)
    hand := []Card{
        {Hearts, Five},
        {Clubs, Five},
        {Spades, Five},
        {Diamonds, Jack},
    }

    // Starter card (the fourth five)
    starter := Card{Diamonds, Five}

    log.Infof("test: TestPerfectHand - Testing the perfect cribbage hand in scoring")
    log.Infof("tess: TestPerfectHand - Hand: %v, Starter: %v", hand, starter)


    got := ScoreHand(hand, starter, false)
    want := 29

    if got != want {
        log.Errorf("❌ ScoreHand = %d, want %d", got, want)
        t.Errorf("ScoreHand = %d, want %d", got, want)
    } else {
        log.Info("✅ Scoring logic PASSED for perfect hand")
    }
}

