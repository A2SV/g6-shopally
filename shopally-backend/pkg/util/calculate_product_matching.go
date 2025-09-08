package util

import (
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
)

func CalculateProductMatching(keywords, productTitle string) float64 {
	if keywords == "" || productTitle == "" {
		return 0.0
	}

	// Normalize both strings
	normalizedKeywords := normalizeString(keywords)
	normalizedProduct := normalizeString(productTitle)

	if normalizedKeywords == "" || normalizedProduct == "" {
		return 0.0
	}

	// Split into words
	keywordList := strings.Fields(normalizedKeywords)
	productWords := strings.Fields(normalizedProduct)

	if len(keywordList) == 0 || len(productWords) == 0 {
		return 0.0
	}

	// Stem all keywords and product words using snowball stemmer
	stemmedKeywords := stemWords(keywordList)
	stemmedProductWords := stemWords(productWords)

	// Calculate score based on stemmed matches
	score := 0.0

	for _, stemmedKeyword := range stemmedKeywords {
		if stemmedKeyword == "" {
			continue
		}

		// Check if stemmed keyword exists in stemmed product words
		for _, stemmedProductWord := range stemmedProductWords {
			if stemmedProductWord == stemmedKeyword {
				score += 10.0 // Add 10 points for each stemmed match
				break         // Only count each keyword once
			}
		}
	}

	return score
}

// stemWords applies snowball stemming to a list of words
func stemWords(words []string) []string {
	stemmed := make([]string, len(words))
	for i, word := range words {
		stemmed[i] = stemWord(word)
	}
	return stemmed
}

// stemWord uses the snowball library for proper stemming
func stemWord(word string) string {
	if word == "" {
		return ""
	}

	// Use snowball stemmer for English language
	stemmed, err := snowball.Stem(word, "english", true)
	if err != nil {
		// If stemming fails, return the original word
		return word
	}
	return stemmed
}

func normalizeString(s string) string {
	// Convert to lowercase and remove extra spaces
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	// Remove punctuation and special characters
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			result.WriteRune(r)
		} else if unicode.IsPunct(r) {
			// Replace punctuation with space
			result.WriteRune(' ')
		}
	}

	// Collapse multiple spaces
	return strings.Join(strings.Fields(result.String()), " ")
}

// CalculateProductMatchingNormalized returns a normalized score between 0-100
func CalculateProductMatchingNormalized(keywords, productTitle string) float64 {
	if keywords == "" || productTitle == "" {
		return 0.0
	}

	// Normalize both strings
	normalizedKeywords := normalizeString(keywords)
	normalizedProduct := normalizeString(productTitle)

	if normalizedKeywords == "" || normalizedProduct == "" {
		return 0.0
	}

	// Split into words
	keywordList := strings.Fields(normalizedKeywords)
	productWords := strings.Fields(normalizedProduct)

	if len(keywordList) == 0 || len(productWords) == 0 {
		return 0.0
	}

	// Stem all keywords and product words using snowball stemmer
	stemmedKeywords := stemWords(keywordList)
	stemmedProductWords := stemWords(productWords)

	// Calculate score based on stemmed matches
	score := 0.0

	for _, stemmedKeyword := range stemmedKeywords {
		if stemmedKeyword == "" {
			continue
		}

		// Check if stemmed keyword exists in stemmed product words
		for _, stemmedProductWord := range stemmedProductWords {
			if stemmedProductWord == stemmedKeyword {
				score += 10.0 // Add 10 points for each stemmed match
				break         // Only count each keyword once
			}
		}
	}

	// Normalize to percentage (0-100)
	maxPossibleScore := float64(len(stemmedKeywords)) * 10.0
	if maxPossibleScore == 0 {
		return 0.0
	}

	normalizedScore := (score / maxPossibleScore) * 100.0

	// Ensure the score is within bounds
	if normalizedScore > 100.0 {
		return 100.0
	}
	return normalizedScore
}

// CalculateProductMatchingSimple returns a simple ratio between 0-1
func CalculateProductMatchingSimple(keywords, productTitle string) float64 {
	if keywords == "" || productTitle == "" {
		return 0.0
	}

	// Normalize both strings
	normalizedKeywords := normalizeString(keywords)
	normalizedProduct := normalizeString(productTitle)

	if normalizedKeywords == "" || normalizedProduct == "" {
		return 0.0
	}

	// Split into words
	keywordList := strings.Fields(normalizedKeywords)
	productWords := strings.Fields(normalizedProduct)

	if len(keywordList) == 0 || len(productWords) == 0 {
		return 0.0
	}

	// Stem all keywords and product words using snowball stemmer
	stemmedKeywords := stemWords(keywordList)
	stemmedProductWords := stemWords(productWords)

	// Count matches
	matches := 0
	for _, stemmedKeyword := range stemmedKeywords {
		if stemmedKeyword == "" {
			continue
		}

		for _, stemmedProductWord := range stemmedProductWords {
			if stemmedProductWord == stemmedKeyword {
				matches++
				break // Only count each keyword once
			}
		}
	}

	// Return simple ratio
	if len(stemmedKeywords) == 0 {
		return 0.0
	}
	return float64(matches) / float64(len(stemmedKeywords))
}
