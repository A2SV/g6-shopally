package middleware

import (
    "bytes"
    "log/slog"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

func TestRecovery_Returns500OnPanic(t *testing.T) {
    gin.SetMode(gin.TestMode)
    var buf bytes.Buffer
    logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo}))

    r := gin.New()
    r.Use(Recovery(logger))
    r.GET("/panic", func(c *gin.Context) { panic("boom") })

    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/panic", nil)
    r.ServeHTTP(w, req)
    if w.Code != http.StatusInternalServerError {
        t.Fatalf("expected 500, got %d", w.Code)
    }
    if !strings.Contains(w.Body.String(), "internal server error") {
        t.Fatalf("expected error body, got %q", w.Body.String())
    }
    if !strings.Contains(buf.String(), "panic_recovered") {
        t.Fatalf("expected panic log entry, got %q", buf.String())
    }
}

func TestRecovery_WithOtherMiddlewares(t *testing.T) {
    gin.SetMode(gin.TestMode)
    var buf bytes.Buffer
    logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo}))

    r := gin.New()
    // ensure interaction with logger and rate limiter doesn't break recovery
    rl := NewIPRateLimiter(rate.Limit(1000), 1, time.Minute)
    r.Use(Logger(logger))
    r.Use(rl.Middleware())
    r.Use(Recovery(logger))
    r.GET("/ok", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/ok", nil)
    r.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }
}
