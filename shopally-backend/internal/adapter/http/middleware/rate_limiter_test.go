package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

func TestRateLimiter_AllowsWithinLimit(t *testing.T) {
    gin.SetMode(gin.TestMode)
    rl := NewIPRateLimiter(rate.Limit(5), 2, time.Minute)

    r := gin.New()
    r.Use(rl.Middleware())
    r.GET("/ok", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

    for i := 0; i < 2; i++ { // burst=2 allowed instantly
        w := httptest.NewRecorder()
        req := httptest.NewRequest(http.MethodGet, "/ok", nil)
        r.ServeHTTP(w, req)
        if w.Code != http.StatusOK {
            t.Fatalf("expected 200, got %d on iter %d", w.Code, i)
        }
    }
}

func TestRateLimiter_ExceedsLimitGets429(t *testing.T) {
    gin.SetMode(gin.TestMode)
    rl := NewIPRateLimiter(rate.Limit(1), 1, time.Minute)

    r := gin.New()
    r.Use(rl.Middleware())
    r.GET("/ok", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

    // First allowed
    w1 := httptest.NewRecorder()
    req1 := httptest.NewRequest(http.MethodGet, "/ok", nil)
    r.ServeHTTP(w1, req1)
    if w1.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w1.Code)
    }

    // Immediate second should be limited
    w2 := httptest.NewRecorder()
    req2 := httptest.NewRequest(http.MethodGet, "/ok", nil)
    r.ServeHTTP(w2, req2)
    if w2.Code != http.StatusTooManyRequests {
        t.Fatalf("expected 429, got %d", w2.Code)
    }
}

func TestRateLimiter_UnknownIPHandled(t *testing.T) {
    gin.SetMode(gin.TestMode)
    rl := NewIPRateLimiter(rate.Limit(1000), 1, time.Minute)

    r := gin.New()
    r.Use(func(c *gin.Context) { // force empty IP
        c.Request.RemoteAddr = ""
        c.Next()
    })
    r.Use(rl.Middleware())
    r.GET("/ok", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/ok", nil)
    r.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }
}
