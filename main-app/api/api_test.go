package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
)

func TestDealAndStatusEndpoints(t *testing.T) {
    router := gin.Default()
    router.GET("/deal", DealHandler)
    router.GET("/status", StatusHandler)
    router.POST("/reset", ResetHandler)

    // Hit /deal
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/deal", nil)
    router.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)

    // Hit /status
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/status", nil)
    router.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
}

