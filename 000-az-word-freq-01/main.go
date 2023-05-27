package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// Open a file
	f, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
clea
	// The frequency of words in the file
	words, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main: %s", err)
	}

	// Display the word frequencies
	// for word, frq := range words {
	// 	fmt.Printf("%s \t\t %d\n", word, frq)
	// }

	// Sort the word frequencies
	pairs := sortWordFrequency(words)

	// Print the sorted results
	for _, pair := range pairs {
		fmt.Printf("%s \t\t %d\n", pair.Key, pair.Value)
	}
}

func freq(r io.Reader) (map[string]int, error) {

	// Create a map to store word frequencies
	wordFreq := make(map[string]int)

	// Create a scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	// Read each word and update the word frequencies
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil
}

// For sorting frequency of words
type Pair struct {
	Key   string
	Value int
}

// implement the Len, Less, and Swap methods on PairList
// to satisfy the sort.Interface interface.

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value // Sort in descending order
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFrequency(m map[string]int) PairList {
	// Convert the map to a pair list
	pairs := make(PairList, len(m))
	i := 0
	for key, value := range m {
		pairs[i] = Pair{key, value}
		i++
	}

	// Sort the pair list
	sort.Sort(pairs)

	return pairs
}

// word frequency in go
// https://chat.openai.com/share/6cb2b004-4cfa-4aad-a2d8-47b0eacd36dd

// sort maps in go
// turn them into structs
// get a slice of those structs
// implement the sort.Interface interface
// https://chat.openai.com/share/03a44e91-fc0d-4cdb-884a-c8acd8f440d8