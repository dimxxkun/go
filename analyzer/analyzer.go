package analyzer

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Stats holds statistical information about processed text.
type Stats struct {
	CharacterCount int
	WordCount      int
	LineCount      int
	UniqueWordCount int
	LongestWord    string
}

var wordRegexp = regexp.MustCompile(`\b[\w'-]+\b`)

// Analyze takes raw input text and returns calculated statistics.
func Analyze(input string) Stats {
	if len(strings.TrimSpace(input)) == 0 {
		return Stats{}
	}

	charCount := utf8.RuneCountInString(input)
	lines := strings.Split(input, "\n")
	lineCount := len(lines)

	words := wordRegexp.FindAllString(strings.ToLower(input), -1)
	wordCount := len(words)

	uniqueWords := make(map[string]struct{})
	longest := ""

	for _, w := range words {
		uniqueWords[w] = struct{}{}
		if utf8.RuneCountInString(w) > utf8.RuneCountInString(longest) {
			longest = w
		}
	}

	return Stats{
		CharacterCount:  charCount,
		WordCount:       wordCount,
		LineCount:       lineCount,
		UniqueWordCount: len(uniqueWords),
		LongestWord:     longest,
	}
}

