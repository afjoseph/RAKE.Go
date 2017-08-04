package rake

import (
	"fmt"
	"regexp"
	"strings"
)

// RegexSplitWords returns a regexp object that split words
func RegexSplitWords() *regexp.Regexp {
	return regexp.MustCompile("[\\p{L}\\d_]+")
}

// RegexSplitSentences returns a regexp object that detects punctuation marks
func RegexSplitSentences() *regexp.Regexp {
	return regexp.MustCompile(`[.,\/#!$%\^&\*;:{}=\-_~()]`)
}

//RegexStopWords builds "stop-words" regex based on a slice of "stop-words"
func RegexStopWords(stopWordsSlice []string) *regexp.Regexp {
	stopWordRegexList := []string{}

	for _, word := range stopWordsSlice {
		wordRegex := fmt.Sprintf(`(?:\A|\z|\s)%s(?:\A|\z|\s)`, word)
		stopWordRegexList = append(stopWordRegexList, wordRegex)
	}

	re := regexp.MustCompile(`(?i)` + strings.Join(stopWordRegexList, "|"))
	return re
}
