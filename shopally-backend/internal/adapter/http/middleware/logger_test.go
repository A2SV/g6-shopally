package middleware

import (
    "bytes"
    "log/slog"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
)

func newTestLogger() (*slog.Logger, *bytes.Buffer) {
    var buf bytes.Buffer
    h := slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
    return slog.New(h), &buf
}

func TestLogger_WritesEntryOnOK(t *testing.T) {
    gin.SetMode(gin.TestMode)
    log, buf := newTestLogger()

    r := gin.New()
    r.Use(Logger(log))
    r.GET("/ok", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/ok", nil)
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }

    out := buf.String()
    if !strings.Contains(out, "http_request") || !strings.Contains(out, "path=/ok") || !strings.Contains(out, "method=GET") || !strings.Contains(out, "status=200") {
        t.Fatalf("expected structured log entry, got: %q", out)
    }
}

func TestLogger_LogsOnAbort(t *testing.T) {
    gin.SetMode(gin.TestMode)
    log, buf := newTestLogger()

    r := gin.New()
    r.Use(Logger(log))
    r.GET("/abort", func(c *gin.Context) {
        c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
    })

    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/abort", nil)
    r.ServeHTTP(w, req)

    if w.Code != http.StatusTooManyRequests {
        t.Fatalf("expected 429, got %d", w.Code)
    }

    out := buf.String()
    if !strings.Contains(out, "http_request") || !strings.Contains(out, "path=/abort") || !strings.Contains(out, "status=429") {
        t.Fatalf("expected log for aborted request, got: %q", out)
    }
}
