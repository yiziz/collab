package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/reiver/go-porterstemmer"
	"github.com/yiziz/collab/path"
	"github.com/yiziz/collab/services/utils"
)

type StringArray []string

type WordScore struct {
	PerkID uint64
	Word   string
	Score  float64
}

type WordScoreArray []*WordScore

func loadStopwords() []string {
	var stopwords []string
	stopfile, err := os.Open("data/stopwords.txt")

	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil
	}

	defer stopfile.Close()

	reader := csv.NewReader(stopfile)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("ERROR: ", err)
			return nil
		}

		stopwords = append(stopwords, record[0])
	}

	return stopwords
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func (t StringArray) Len() int {
	return len(t)
}

func (t StringArray) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t StringArray) Swap(i, j int) {
	tmp := t[i]
	t[i] = t[j]
	t[j] = tmp
}

func (t WordScoreArray) Len() int {
	return len(t)
}

func (t WordScoreArray) Less(i, j int) bool {
	return t[i].Score < t[j].Score
}

func (t WordScoreArray) Swap(i, j int) {
	tmp := t[i]
	t[i] = t[j]
	t[j] = tmp
}

// Returns a map where the key is a word and the int is the number
// of times that word appears in the set of documents
//
// Providing a threshold (1.0 >= x > 0.0) will return only the words
// that appear in all the documents (x*100)% of the time
func termFrequency(recordArray [][]string, threshold float64) (m map[string]int, err error) {
	saveMap := make(map[string]map[string]int)
	for _, record := range recordArray {
		url := record[0]

		if _, ok := saveMap[url]; ok {
			continue
		}

		words := utils.LowercaseWords(strings.Fields(record[2]))

		for i := range words {
			w, err := utils.RemoveNonAlphaNumeric(words[i])
			if err != nil {
				continue
			} else {
				words[i] = w
			}
		}

		words, err = utils.RemoveStopwords(words)
		if err != nil {
			return nil, err
		}

		saveMap[url] = utils.WordFrequency(words)
	}

	documentFrequencyMap := make(map[string]int)

	for _, wordCountMap := range saveMap {
		for word := range wordCountMap {
			if _, ok := documentFrequencyMap[word]; ok {
				documentFrequencyMap[word]++
			} else {
				documentFrequencyMap[word] = 1
			}
		}
	}

	if threshold != 0.0 {
		for word, value := range documentFrequencyMap {
			if float64(value)/float64(len(saveMap)) < threshold {
				delete(documentFrequencyMap, word)
			}
		}
	}
	return documentFrequencyMap, nil
}

// Inverse Document Frequency
func inverseDocumentFrequency(recordArray [][]string) (m map[string]float64, err error) {
	d := float64(len(recordArray))

	wordCountMap := make(map[string]int)
	for _, record := range recordArray {
		words := utils.LowercaseWords(strings.Fields(record[2]))

		for i := range words {
			w, err := utils.RemoveNonAlphaNumeric(words[i])
			if err != nil {
				continue
			} else {
				words[i] = w
			}
		}

		words, err = utils.RemoveStopwords(words)
		if err != nil {
			return nil, err
		}

		words = utils.RemoveDuplicates(words)

		for _, word := range words {
			if _, ok := wordCountMap[word]; ok {
				wordCountMap[word]++
			} else {
				wordCountMap[word] = 1
			}
		}
	}

	idfMap := make(map[string]float64)
	for word, value := range wordCountMap {
		idfMap[word] = math.Log(d / float64(value))
	}
	return idfMap, nil
}

// Term Frequency-Inverse Document Frequency (TF-IDF)
func termFrequencyInverseDocumentFrequency(recordArray [][]string) (m map[string]float64, err error) {
	tfidfMap := make(map[string]float64)

	tf, err := termFrequency(recordArray, 0.0)
	if err != nil {
		return nil, err
	}

	idf, err := inverseDocumentFrequency(recordArray)
	if err != nil {
		return nil, err
	}

	for word, docFreq := range idf {
		tfidfMap[word] = float64(tf[word]) * docFreq
	}
	return tfidfMap, nil
}

func loadPerksData(stopwords []string) map[uint64]string {
	perks := make(map[uint64]string)
	// file, err := os.Open("data/perks_data.csv")
	file, err := os.Open(path.ProjectPath() + "/fixtures/perks_data.csv")

	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("ERROR: ", err)
			return nil
		}

		content := strings.Join(record[1:], " ")
		content = stripchars(content, "1234567890!@#$%^&*()`~'-=_+[]{}:;<,>./?\"\n\t")

		unfiltered := StringArray(strings.Split(content, " "))
		sort.Stable(unfiltered)
		var filtered_strings StringArray

		for _, word := range unfiltered {
			word = strings.ToLower(word)

			idx := sort.Search(len(stopwords), func(x int) bool {
				return stopwords[x] >= word
			})

			if idx < len(stopwords) && stopwords[idx] == word {
				// Word is a stopword
			} else if len(word) > 0 {
				stemmed_word := porterstemmer.StemString(word)
				filtered_strings = append(filtered_strings, stemmed_word)
			}
		}

		fmt.Println("cooL:", filtered_strings, "stuff:", len(filtered_strings))

		id, _ := strconv.ParseUint(record[0], 10, 64)
		perks[id] = strings.Join(filtered_strings, " ")
	}

	return perks
}

var TermScores map[string]float64

func PerkTermScores() map[uint64]map[string]float64 {
	perks_data := loadPerksData(nil)
	var entire_data [][]string

	for _, perk_data := range perks_data {
		entire_data = append(entire_data, strings.Split(perk_data, " "))
	}

	TermScores, _ := termFrequencyInverseDocumentFrequency(entire_data)
	print_term_scores(TermScores)

	perks_to_terms := make(map[uint64]map[string]float64)

	for perk_id, perk_terms := range perks_data {
		perks_to_terms[perk_id] = make(map[string]float64)
		terms_on_perk := perks_to_terms[perk_id]

		for _, term := range strings.Split(perk_terms, " ") {
			score := TermScores[term]

			if score > 0 {
				terms_on_perk[term] = score
			}
		}

		if len(terms_on_perk) == 0 {
			fmt.Println("ZERO:", perk_id)
		}
	}

	return perks_to_terms
}

func print_term_scores(term_scores map[string]float64) {
	var word_scores WordScoreArray

	for k, v := range term_scores {
		word_score := new(WordScore)
		word_score.Score = v
		word_score.Word = k
		word_score.PerkID = 42

		word_scores = append(word_scores, word_score)
	}

	sort.Stable(word_scores)

	for _, word_score := range word_scores {
		fmt.Println("word:", word_score.Word, "score:", word_score.Score)
	}
}
