package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func Top10(text string) []string {

	words := strings.Fields(text)

	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	var wordCounts []WordCount
	for word, count := range wordsMap {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count != wordCounts[j].Count {
			return wordCounts[i].Count > wordCounts[j].Count
		}
		return wordCounts[i].Word < wordCounts[j].Word
	})

	if len(wordCounts) > 10 {
		result := make([]string, 10)
		for i := 0; i < 10; i++ {
			result[i] = wordCounts[i].Word
		}
		return result
	}

	result := make([]string, len(wordCounts))
	for i, w := range wordCounts {
		result[i] = w.Word
	}
	return result
}
