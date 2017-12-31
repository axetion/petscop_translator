package petscop

import "fmt"

var PHONEMES = map[[2]string][]string{
	{"R2", "X"}:        []string{"UW"},
	{"R2", "SQUARE"}:   []string{"P"},
	{"R2", "CIRCLE"}:   []string{"T"},
	{"R2", "TRIANGLE"}: []string{"B"},

	{"R2", "LEFT"}:  []string{"TH"},
	{"R2", "UP"}:    []string{"F"},
	{"R2", "DOWN"}:  []string{"V"},
	{"R2", "START"}: []string{"D"},

	{"R1", "X"}:        []string{"EY"},
	{"R1", "CIRCLE"}:   []string{"IY", "IH"},
	{"R1", "TRIANGLE"}: []string{"IH"},

	{"R1", "UP"}: []string{"OW"},

	{"L2", "X"}:        []string{"M"},
	{"L2", "SQUARE"}:   []string{"N"},
	{"L2", "CIRCLE"}:   []string{"R"},
	{"L2", "TRIANGLE"}: []string{"L"},

	{"L2", "LEFT"}:  []string{"UH", "AH"},
	{"L2", "DOWN"}:  []string{"HH"},
	{"L2", "START"}: []string{"W"},

	{"L1", "X"}:        []string{"S", "Z"},
	{"L1", "SQUARE"}:   []string{"Z"},
	{"L1", "TRIANGLE"}: []string{"SH"},

	{"L1", "DOWN"}:  []string{"K"},
	{"L1", "START"}: []string{"CH"},

	{"X", ""}:        []string{"AA"},
	{"SQUARE", ""}:   []string{"AE"},
	{"TRIANGLE", ""}: []string{"AO"},

	{"LEFT", ""}:  []string{"EH"},
	{"UP", ""}:    []string{"AY"},
	{"START", ""}: []string{"AH"},
}

func ToPhonemes(buttons [][2]string) ([][]string, error) {
	output := make([][]string, len(buttons))

	for i, button := range buttons {
		if possibilities, ok := PHONEMES[button]; ok {
			output[i] = possibilities
		} else {
			return nil, fmt.Errorf("Unknown button combination %v.", button)
		}
	}

	return output, nil
}
