package checker

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var incorrect_words = []string{}
var doc_suggestions [][]string
var doc_contexts [][]string

func CheckSpelling(dict []string, doc *bufio.Scanner) []string {
	incorrect_details := []string{}
	line_num := 0

	for doc.Scan() {
		line_num += 1
		line := strings.TrimSpace(doc.Text())
		doc_line := strings.Split(line, " ")
		for i, word := range doc_line {
			word_comparison := CompareDict(dict, strings.TrimSpace(word))
			if !word_comparison {
				incorrect := word + " (" + strconv.Itoa(line_num) + ":" + strconv.Itoa(getColumnNumber(doc_line, i)) + ")"
				incorrect_words = append(incorrect_words, word)
				incorrect_details = append(incorrect_details, incorrect)
				doc_contexts = append(doc_contexts, BuildContextString(doc_line, word, i))
			}
		}
	}

	return incorrect_details
}

func CompareDict(dict []string, word string) bool {
	for _, dict_word := range dict {
		if strings.Compare(dict_word, word) == 0 {
			return true
		}

		if word[0] >= 65 && word[0] <= 90 {
			return true
		}
	}
	return false
}

func getColumnNumber(line []string, word_index int) int {
	col_number := word_index
	i := 0

	for i < word_index {
		col_number = +col_number + len(line[i])
		i++
	}

	return col_number
}

func BuildContextString(line []string, word string, word_index int) []string {
	line[word_index] = " [" + word + "] "

	return line
}

func GetSuggestions(dict []string) [][]string {
	suggestions := []string{}

	for _, word := range incorrect_words {
		for _, dict_word := range dict {
			distance := levenshtein.DistanceForStrings([]rune(word), []rune(dict_word), levenshtein.DefaultOptions)
			if distance <= 3 {
				suggestions = append(suggestions, dict_word)
			}
		}
		doc_suggestions = append(doc_suggestions, suggestions)
		suggestions = []string{}
	}

	return doc_suggestions
}

func GetDocContexts() [][]string {
	return doc_contexts
}
