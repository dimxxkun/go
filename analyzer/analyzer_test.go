package analyzer

import "testing"

func TestAnalyzeEmpty(t *testing.T) {
	stats := Analyze("")
	if stats.WordCount != 0 || stats.CharacterCount != 0 {
		t.Fatalf("expected empty stats, got %+v", stats)
	}
}

func TestAnalyzeSampleText(t *testing.T) {
	text := "Hello world!\nWelcome to Go programming.\nHello again."
	stats := Analyze(text)

	if stats.LineCount != 3 {
		t.Errorf("expected 3 lines, got %d", stats.LineCount)
	}

	if stats.WordCount != 8 {
		t.Errorf("expected 8 words, got %d", stats.WordCount)
	}

	if stats.UniqueWordCount != 7 { // "hello" is repeated
		t.Errorf("expected 7 unique words, got %d", stats.UniqueWordCount)
	}

	if stats.LongestWord != "programming" {
		t.Errorf("expected 'programming' as longest word, got '%s'", stats.LongestWord)
	}
}
