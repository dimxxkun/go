package main

import (
	"fmt"
	"strings"

	"github.com/dimxxkun/go-quickstart/analyzer"
)

func main() {
	sample := `Go is an open source programming language supported by Google.
Easy to learn and great for building scalable services.
Concurrency is easy with goroutines and channels.`

	fmt.Println("=========================================")
	fmt.Println("   Go Quickstart - Text Analyzer CLI     ")
	fmt.Println("=========================================")
	fmt.Println("\nAnalyzing input text:")
	fmt.Println(strings.Repeat("-", 45))
	fmt.Println(sample)
	fmt.Println(strings.Repeat("-", 45))

	stats := analyzer.Analyze(sample)

	fmt.Printf("\nStatistics:\n")
	fmt.Printf("  • Total Characters:   %d\n", stats.CharacterCount)
	fmt.Printf("  • Total Words:        %d\n", stats.WordCount)
	fmt.Printf("  • Total Lines:        %d\n", stats.LineCount)
	fmt.Printf("  • Unique Words:       %d\n", stats.UniqueWordCount)
	fmt.Printf("  • Longest Word:       %s (%d chars)\n", stats.LongestWord, len(stats.LongestWord))
	fmt.Println("\nExecution completed successfully!")
}
