package main

import (
	"os"
	"log"
	"fmt"
	"io"
	"bufio"
	"encoding/csv"
	"github.com/cloudflare/ahocorasick"
)

func createKeywordDictionaryFromCSVFile(f *os.File) map[string]string {
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	m := make(map[string]string)
	for {
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		// Display record.
		//fmt.Printf("%v	%v\n", record[0], record[1])
		m[record[0]] = record[1]
	}

	return m
}

func getKeys(m map[string]string) []string {
	i := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func main() {
	f, err := os.Open("./data/keywords.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	dict := createKeywordDictionaryFromCSVFile(f)
	keywords := getKeys(dict)
	m := ahocorasick.NewStringMatcher(keywords)
	hits := m.Match([]byte("http://www.adoreme.com/bras-and-panties.html"))
	fmt.Printf("# of hits: %d\n", len(hits))
	for _, v := range hits {
		fmt.Printf("%d %s\n", v, keywords[v])
	}
}