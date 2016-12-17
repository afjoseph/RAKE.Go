package rake

import (
	"fmt"
	"regexp"
	"strings"
)

func Regex_SplitWords() *regexp.Regexp {
	return regexp.MustCompile(`(\w+)`)
}

func Regex_SplitSentences() *regexp.Regexp {
	return regexp.MustCompile(`[.,\/#!$%\^&\*;:{}=\-_~()]`)
}

//Build "stop-words" regex
func Regex_StopWords(stopWordFilePath string) *regexp.Regexp {
	stopWordList := LoadStopWords(stopWordFilePath)
	stopWordRegexList := []string{}

	for _, word := range stopWordList {
		wordRegex := fmt.Sprintf(`(\b%s\b)`, word)
		stopWordRegexList = append(stopWordRegexList, wordRegex)
	}

	re := regexp.MustCompile(`(?i)` + strings.Join(stopWordRegexList, "|"))
	return re
}
