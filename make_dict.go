package main

import (
	"bufio"
	"encoding/gob"
	stemmer "go-porterstemmer"
	"log"
	"os"
	"petscop"
	"strings"
)

func stripLexicalStress(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' || r == '(' || r == ')' {
			return -1
		}

		return r
	}, s)
}

func writeFrequencies(input string, output string) {
	frequenciesFh, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer frequenciesFh.Close()

	frequenciesReader := bufio.NewScanner(frequenciesFh)
	frequencies := make(map[string]int)

	for frequenciesReader.Scan() {
		word := stemmer.StemString(frequenciesReader.Text())

		if _, ok := frequencies[word]; !ok {
			frequencies[word] = len(frequencies)
		}
	}

	if err := frequenciesReader.Err(); err != nil {
		log.Fatal(err)
	}

	treeFh, err := os.Create(output)

	if err != nil {
		log.Fatal(err)
	}

	defer treeFh.Close()

	encoder := gob.NewEncoder(treeFh)

	if err := encoder.Encode(frequencies); err != nil {
		log.Fatal(err)
	}
}

func writeDictionary(inputPath string, dictionaryPath string) {
	dictionaryFh, err := os.Open(inputPath)

	if err != nil {
		log.Fatal(err)
	}

	defer dictionaryFh.Close()

	dictionary := bufio.NewScanner(dictionaryFh)

	tree := petscop.BKTree{}

	for dictionary.Scan() {
		entry := dictionary.Text()

		if entry[0] == ';' {
			continue
		}

		parts := strings.Fields(stripLexicalStress(entry))

		word := parts[0]
		phonemes := make([][]string, len(parts)-1)

		for i := 1; i < len(parts); i++ {
			phonemes[i-1] = []string{parts[i]}
		}

		tree.Add(phonemes, word)
	}

	if err := dictionary.Err(); err != nil {
		log.Fatal(err)
	}

	treeFh, err := os.Create(dictionaryPath)

	if err != nil {
		log.Fatal(err)
	}

	defer treeFh.Close()

	if err := tree.Serialize(treeFh); err != nil {
		log.Fatal(err)
	}
}

func main() {
	writeDictionary("cmudict-0.7b", "data/dictionary.gob")
	writeFrequencies("frequencies", "data/frequencies.gob")
}
