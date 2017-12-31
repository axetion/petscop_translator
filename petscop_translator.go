package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"petscop"
)

func main() {
	verboseFlag := len(os.Args) > 1 && os.Args[1] == "-v"

	fmt.Println("Loading dictionary...")

	dictionaryFh, err := os.Open("dictionary.gob")

	if err != nil {
		log.Fatal(err)
	}

	tree := petscop.BKTree{}
	if err := tree.Deserialize(dictionaryFh); err != nil {
		log.Fatal(err)
	}

	frequenciesFh, err := os.Open("frequencies.gob")

	if err != nil {
		log.Fatal(err)
	}

	frequencies := make(map[string]int)

	decoder := gob.NewDecoder(frequenciesFh)

	if err := decoder.Decode(&frequencies); err != nil {
		log.Fatal(err)
	}

	frequenciesFh.Close()

	fmt.Println("---")

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Ask? ")

	for input.Scan() {
		line := input.Text()

		if line == "quit" {
			break
		}

		phonemes, err := petscop.ToPhonemes(petscop.ParseButtons(line))

		if err != nil {
			fmt.Printf("%s\n\nAsk? ", err.Error())
		} else {
			if verboseFlag {
				fmt.Printf("Phonemes: %s\n", phonemes)
			}

			matches := tree.Find(phonemes, 1)

			if verboseFlag {
				fmt.Printf("All matches: %v\n", matches)
			}

			if len(matches.Matches) == 0 {
				fmt.Print("No matches.\n\nAsk? ")
			} else {
				match := petscop.PickMatch(matches.Matches, frequencies)

				if match.Distance == 0 {
					fmt.Printf("> %s (exact match)\n\nAsk? ", match.Word)
				} else {
					fmt.Printf("> %s (approximate match)\n\nAsk? ", match.Word)
				}
			}
		}
	}
}
