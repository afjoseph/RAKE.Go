package rake

import (
	"fmt"
	"regexp"
	"strings"
)

// RegexSplitWords returns a regexp object that split words
func RegexSplitWords() *regexp.Regexp {
	return regexp.MustCompile(`(\w+)`)
}

// RegexSplitSentences returns a regexp object that detects punctuation marks
func RegexSplitSentences() *regexp.Regexp {
	return regexp.MustCompile(`[.,\/#!$%\^&\*;:{}=\-_~()]`)
}

//RegexStopWords builds "stop-words" regex based on a slice of "stop-words"
func RegexStopWords() *regexp.Regexp {
	stopWordList := LoadStopWords()
	stopWordRegexList := []string{}

	for _, word := range stopWordList {
		wordRegex := fmt.Sprintf(`(\b%s\b)`, word)
		stopWordRegexList = append(stopWordRegexList, wordRegex)
	}

	re := regexp.MustCompile(`(?i)` + strings.Join(stopWordRegexList, "|"))
	return re
}
