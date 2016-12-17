package rake

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

/* LoadStopWords is a Utility function to load stop words from a file and return a list of words
   @param stop_word_file Path and file name of a file containing stop words.
   @return list A list of stop words. */
func LoadStopWords(filePath string) []string {
	stopWords := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if string(line[0]) != "#" {
			substrings := strings.Split(line, ` `)
			for _, word := range substrings {
				stopWords = append(stopWords, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stopWords
}

/* Utility function to return a list of all words that are have a length greater than a specified number of characters.
   @param text The text that must be split in to words.
   @param min_word_return_size The minimum no of characters a word must have to be included.*/
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

/* Utility function to return a list of sentences.
   @param text The text that must be split in to sentences.
*/
func SplitSentences(text string) []string {
	splitText := RegexSplitSentences().ReplaceAllString(strings.TrimSpace(text), "\n")
	return strings.Split(splitText, "\n")
}

func GenerateCandidateKeywords(sentenceList []string, stopWordPattern *regexp.Regexp) []string {
	phraseList := []string{}

	for _, sentence := range sentenceList {
		tmp := stopWordPattern.ReplaceAllString(strings.TrimSpace(sentence), "|")

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

//Util function as a Go replacement for Python's `setDefault`: https://docs.python.org/2/library/stdtypes.html#dict.setdefault
//Basically, if key is in the dictionary, return its value. If not, insert key with a value of default and return default. default defaults to None.
func SetDefaultStringInt(h map[string]int, k string, v int) (set bool, r int) {
	if r, set = h[k]; !set {
		h[k] = v
		r = v
		set = true
	}
	return
}

//Util function as a Go replacement for Python's `setDefault`: https://docs.python.org/2/library/stdtypes.html#dict.setdefault
//Basically, if key is in the dictionary, return its value. If not, insert key with a value of default and return default. default defaults to None.
func SetDefaultStringFloat64(h map[string]float64, k string, v float64) (set bool, r float64) {
	if r, set = h[k]; !set {
		h[k] = v
		r = v
		set = true
	}
	return
}

func RunRake(text string) PairList {
	//Split sentences based on punctuation
	sentenceList := SplitSentences(text)

	//Build stop-word regex pattern
	stopPath := "/Users/ab/SmartStoplist.txt"
	stopWordPattern := RegexStopWords(stopPath)

	//Build phrase list
	phraseList := GenerateCandidateKeywords(sentenceList, stopWordPattern)

	//Build word scores
	wordScores := CalculateWordScores(phraseList)

	//Build keyword candidates and sort it (see sort.go)
	keywordCandidates := GenerateCandidateKeywordScores(phraseList, wordScores)
	sorted := reverseSortByValue(keywordCandidates)
	return sorted
}
