package api

import (

    // STD LIBS
    "net/http"
    "sync"
    "fmt"

    // 3PP
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"

    // Our Libs
    core "cribbage/core"
)

// -- GLOBAL STATE (for simplicity in 0.1.2) --
var (
    mu   sync.Mutex
    game *core.Game
)

// -- INIT LOGGER --
func init() {
    log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
    log.SetLevel(log.InfoLevel)
}

// -- API HANDLERS --

// Deal godoc
// @Summary Deals out a Cribbage Hand
// @Description Deals out and Starts a Cribbage hand
// @Tags cribbage
// @Produce json
// @Success 200 {object} api.DealResponse
// @Router /deal [get]
func DealHandler(c *gin.Context) {
    mu.Lock()
    defer mu.Unlock()

    if game == nil {
        // Default: 2 players, new RNG
        game = core.NewGame(
            core.Mode1v1,
            []string{"Alice", "Bob"},
            nil,
        )
    }
    err := game.Deal()
    if err != nil {
        log.Errorf("Deal error: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    log.Info("Dealt new hands")
    dealResp := DealResponse {
        Players: game.Players,
        Starter: game.Starter,
        State:   fmt.Sprintf("%v", game.State),
    }

    c.JSON(http.StatusOK, dealResp )
    
}

// Status godoc
// @Summary Gives the Status of a Cribbage Deal
// @Description Provides the current status of a cribbage game
// @Tags cribbage
// @Produce json
// @Success 200 {object} api.DealResponse
// @Router /status [get]
func StatusHandler(c *gin.Context) {
    mu.Lock()
    defer mu.Unlock()

    if game == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No active game"})
        return
    }
    dealResp := StatusResponse {
        Players: game.Players,
	Crib:	 game.Crib,
        Starter: game.Starter,
        State:   game.State,
        Turn:    game.CurrentTurn,
        Dealer:  game.CribOwnerIdx,
    }

    c.JSON(http.StatusOK, dealResp )
}

func ResetHandler(c *gin.Context) {
    mu.Lock()
    defer mu.Unlock()

    game = nil
    log.Warn("Game state reset")
    c.JSON(http.StatusOK, gin.H{"result": "Game reset"})
}

