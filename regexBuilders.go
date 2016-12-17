package rake

import (
	"fmt"
	"regexp"
	"strings"
)

func RegexSplitWords() *regexp.Regexp {
	return regexp.MustCompile(`(\w+)`)
}

func RegexSplitSentences() *regexp.Regexp {
	return regexp.MustCompile(`[.,\/#!$%\^&\*;:{}=\-_~()]`)
}

//Build "stop-words" regex
func RegexStopWords(stopWordFilePath string) *regexp.Regexp {
	stopWordList := LoadStopWords(stopWordFilePath)
	stopWordRegexList := []string{}

	for _, word := range stopWordList {
		wordRegex := fmt.Sprintf(`(\b%s\b)`, word)
		stopWordRegexList = append(stopWordRegexList, wordRegex)
	}

	re := regexp.MustCompile(`(?i)` + strings.Join(stopWordRegexList, "|"))
	return re
}
