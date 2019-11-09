package rake

import (
	"regexp"
	"strconv"
	"strings"
)

//IsNumber returns true if the supplied string is a number
func IsNumber(str string) bool {
	if strings.Contains(str, ".") { //deal with float
		if _, err := strconv.ParseFloat(str, 32); err != nil {
			return false
		}
	} else { //deal with int
		if _, err := strconv.ParseInt(str, 10, 32); err != nil {
			return false
		}
	}

	return true
}

// SeperateWords returns a slice of all words that have a length greater than a specified number of characters.
func SeperateWords(text string) []string {
	words := []string{}

	splitWords := RegexSplitWords().FindAllString(text, -1)
	for _, singleword := range splitWords {
		currentword := strings.ToLower(strings.TrimSpace(singleword))
		if currentword != "" && !IsNumber(currentword) {
			words = append(words, currentword)
		}
	}

	return words
}

// SplitSentences returns a slice of sentences.
func SplitSentences(text string) []string {
	splitText := RegexSplitSentences().ReplaceAllString(strings.TrimSpace(text), "\n")
	return strings.Split(splitText, "\n")
}

//GenerateCandidateKeywords returns a slice of candidate keywords from a slice of sentences and a stop-words regex
func GenerateCandidateKeywords(sentenceList []string, stopWordPattern *regexp.Regexp) []string {
	phraseList := []string{}

	for _, sentence := range sentenceList {
		tmp := stopWordPattern.ReplaceAllString(strings.TrimSpace(sentence), " | ")
		for {
			abc := len(tmp)
			tmp = stopWordPattern.ReplaceAllString(strings.TrimSpace(tmp), " | ")
			if abc == len(tmp) {
				break
			}
		}

		multipleWhiteSpaceRe := regexp.MustCompile(`\s\s+`)
		tmp = multipleWhiteSpaceRe.ReplaceAllString(strings.TrimSpace(tmp), " ")

		phrases := strings.Split(tmp, "|")
		for _, phrase := range phrases {
			phrase = strings.ToLower(strings.TrimSpace(phrase))
			if phrase != "" {
				phraseList = append(phraseList, phrase)
			}
		}
	}

	return phraseList
}

// CalculateWordScores returns a map of (string,float64) that maps to a candidate word and its score in the text
func CalculateWordScores(phraseList []string) map[string]float64 {
	wordFrequency := map[string]int{}
	wordDegree := map[string]int{}

	for _, phrase := range phraseList {
		wordList := SeperateWords(phrase)
		wordListLength := len(wordList)
		wordListDegree := wordListLength - 1

		for _, word := range wordList {
			SetDefaultStringInt(wordFrequency, word, 0)
			wordFrequency[word]++

			SetDefaultStringInt(wordDegree, word, 0)
			wordDegree[word] += wordListDegree
		}
	}

	for key := range wordFrequency {
		wordDegree[key] = wordDegree[key] + wordFrequency[key]
	}

	wordScore := map[string]float64{}
	for key := range wordFrequency {
		SetDefaultStringFloat64(wordScore, key, 0)
		wordScore[key] = float64(wordDegree[key]) / float64(wordFrequency[key])
	}

	return wordScore
}

//GenerateCandidateKeywordScores returns a map of (string,float64) that contains keywords and their score in the text
func GenerateCandidateKeywordScores(phraseList []string, wordScore map[string]float64) map[string]float64 {
	keywordCandidates := map[string]float64{}

	for _, phrase := range phraseList {
		SetDefaultStringFloat64(keywordCandidates, phrase, 0)
		wordList := SeperateWords(phrase)
		candidateScore := float64(0.0)

		for _, word := range wordList {
			candidateScore = candidateScore + wordScore[word]
		}

		keywordCandidates[phrase] = candidateScore
	}

	return keywordCandidates
}

//SetDefaultStringInt is a util function that serves as a Go replacement for Python's `setDefault`: https://docs.python.org/2/library/stdtypes.html#dict.setdefault
//Basically, if key is in the dictionary, return its value. If not, insert key with a value of default and return default. default defaults to None.
func SetDefaultStringInt(h map[string]int, k string, v int) (set bool, r int) {
	if r, set = h[k]; !set {
		h[k] = v
		r = v
		set = true
	}
	return
}

//SetDefaultStringFloat64 is a util function that serves as a Go replacement for Python's `setDefault`: https://docs.python.org/2/library/stdtypes.html#dict.setdefault
//Basically, if key is in the dictionary, return its value. If not, insert key with a value of default and return default. default defaults to None.
func SetDefaultStringFloat64(h map[string]float64, k string, v float64) (set bool, r float64) {
	if r, set = h[k]; !set {
		h[k] = v
		r = v
		set = true
	}
	return
}

//RunRakeI18N returns a slice of key-value pairs (PairList) of a keyword and its score after running the RAKE algorithm on a given text
func RunRakeI18N(text string, stopWords []string) PairList {
	//Split sentences based on punctuation
	sentenceList := SplitSentences(text)

	//Build stop-word regex pattern
	words := StopWordsSlice
	if len(stopWords) > 0 {
		words = stopWords
	}
	stopWordPattern := RegexStopWords(words)

	//Build phrase list
	phraseList := GenerateCandidateKeywords(sentenceList, stopWordPattern)

	//Build word scores
	wordScores := CalculateWordScores(phraseList)

	//Build keyword candidates and sort it (see sort.go)
	keywordCandidates := GenerateCandidateKeywordScores(phraseList, wordScores)
	sorted := reverseSortByValue(keywordCandidates)
	return sorted
}

// RunRake wraps RunRakeI18N to respect API
func RunRake(text string) PairList {
	return RunRakeI18N(text, []string{})
}
