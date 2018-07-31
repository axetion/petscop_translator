package petscop

import (
	stemmer "go-porterstemmer"
	"math"
	"regexp"
	"strings"
)

/* put here any irregular verbs the stemmer doesn't handle */
var IRREGULAR_VERBS = map[string]string{
	"fell": "fall",
	"am":   "be",
	"is":   "be",
	"are":  "be",
}

var buttonRegexp = regexp.MustCompile("[^[:punct:]]+")

func ParseButtons(line string) [][2]string {
	parts := strings.Fields(strings.ToUpper(line))

	buttons := [][2]string{}

	for i, part := range parts {
		buttons = append(buttons, [2]string{"", ""})

		button := buttonRegexp.FindAllString(part, 2)
		copy(buttons[i][:], button[:])
	}

	return buttons
}

func PickMatch(matches []Match, frequencies map[string]int) Match {
	match := Match{Distance: math.MaxInt32}

	for _, possible := range matches {
		if match.Distance > possible.Distance {
			match = possible
		} else if match.Distance == possible.Distance {
			var currentRarity, newRarity int
			var ok bool

			stem1 := stemmer.StemString(match.Word)

			if participle, ok := IRREGULAR_VERBS[stem1]; ok {
				stem1 = participle
			}

			if currentRarity, ok = frequencies[stem1]; !ok {
				currentRarity = len(frequencies) + 1
			}

			stem2 := stemmer.StemString(possible.Word)

			if participle, ok := IRREGULAR_VERBS[stem2]; ok {
				stem2 = participle
			}

			if newRarity, ok = frequencies[stem2]; !ok {
				newRarity = len(frequencies) + 1
			}

			if currentRarity > newRarity || currentRarity == newRarity && match.Word > possible.Word {
				match = possible
			}
		}
	}

	return match
}
