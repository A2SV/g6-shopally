package util

import (
	"fmt"
	"log"
	"os"
	"sync"

	"container/heap"

	"github.com/joho/godotenv"
)

func GeminiAPIKeysList() []string {
	listOfTikenKeys := []string{}
	// Load env file
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			// Just log the error but don't fail - this is normal on Render
			log.Println("Note: .env file not found (this is expected in production)")
		}
	}

	// Load up to 17 API keys from environment variables
	next := true
	index := 1

	for next {
		key := fmt.Sprintf("GEMINI_API_KEY_%d", index)

		token, exist := os.LookupEnv(key)
		if exist && token != "" {
			listOfTikenKeys = append(listOfTikenKeys, token)
		}
		if !exist {
			next = false
		}
		index++
	}

	log.Println("GeminiAPIKeysList: Loaded", len(listOfTikenKeys), "API keys for Gemini")

	return listOfTikenKeys
}

// TokenEntry represents a Gemini API token with request count
type TokenEntry struct {
	Token string
	Count int // Number of requests made with this token
	index int // Index in the heap (required by heap.Interface)
}

type TokenHeap []*TokenEntry

// Implement heap.Interface methods - now based on Count (min-heap)
func (h TokenHeap) Len() int           { return len(h) }
func (h TokenHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }
func (h TokenHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }

func (h *TokenHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*TokenEntry)
	item.index = n
	*h = append(*h, item)
}

func (h *TokenHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// TokenManager manages the heap of tokens and provides thread-safe access
type TokenManager struct {
	heap   *TokenHeap
	tokens map[string]*TokenEntry
	mu     sync.RWMutex
}

// NewTokenManager creates a new token manager
func NewTokenManager() *TokenManager {
	h := make(TokenHeap, 0)
	heap.Init(&h)

	return &TokenManager{
		heap:   &h,
		tokens: make(map[string]*TokenEntry),
	}
}

// InitializeTokens loads tokens from environment and adds them to the manager
func (tm *TokenManager) InitializeTokens() {
	tokens := GeminiAPIKeysList()
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for _, token := range tokens {
		if _, exists := tm.tokens[token]; !exists {
			entry := &TokenEntry{
				Token: token,
				Count: 0,
			}
			tm.tokens[token] = entry
			heap.Push(tm.heap, entry)
		}
	}
}

// GetBestToken returns the token with the lowest request count
func (tm *TokenManager) GetBestToken() string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.heap.Len() == 0 {
		return ""
	}

	return (*tm.heap)[0].Token
}

// IncrementCount increments the request count for a token and rebalances the heap
func (tm *TokenManager) IncrementCount(token string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	entry, exists := tm.tokens[token]
	if !exists {
		return
	}

	entry.Count++

	// Rebalance the heap since count changed
	heap.Fix(tm.heap, entry.index)
}

// GetTokenStats returns statistics for all tokens
func (tm *TokenManager) GetTokenStats() map[string]map[string]interface{} {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	stats := make(map[string]map[string]interface{})
	for token, entry := range tm.tokens {
		stats[token] = map[string]interface{}{
			"count":      entry.Count,
			"heap_index": entry.index,
		}
	}

	return stats
}

// GetNextToken returns the best token (lowest count) and increments its count
func (tm *TokenManager) GetNextToken() string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.heap.Len() == 0 {
		return ""
	}

	// Get the best token (lowest count)
	bestToken := (*tm.heap)[0].Token

	// Increment its count and rebalance
	if entry, exists := tm.tokens[bestToken]; exists {
		entry.Count++
		heap.Fix(tm.heap, entry.index)
	}

	return bestToken
}

// ResetAllCounts resets all token counts to zero (useful for periodic reset)
func (tm *TokenManager) ResetAllCounts() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for _, entry := range tm.tokens {
		entry.Count = 0
	}

	// Rebuild the heap since all counts changed
	// (More efficient than fixing each entry individually)
	oldHeap := *tm.heap
	*tm.heap = make(TokenHeap, 0, len(oldHeap))
	for _, entry := range oldHeap {
		heap.Push(tm.heap, entry)
	}
}

// GetTokenCount returns the current count for a specific token
func (tm *TokenManager) GetTokenCount(token string) int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if entry, exists := tm.tokens[token]; exists {
		return entry.Count
	}
	return -1
}
