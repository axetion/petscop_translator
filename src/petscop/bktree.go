package petscop

import (
	"encoding/gob"
	"io"
)

type bknode struct {
	Key   [][]string
	Words []string

	Children []bknode
}

type BKTree struct {
	root *bknode
}

type Matches struct {
	Matches []Match
}

type Match struct {
	Word     string
	Distance int
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func (tree *BKTree) Add(key [][]string, word string) bool {
	if tree.root == nil {
		tree.root = &bknode{Key: key, Words: []string{word}}
		return true
	} else {
		return doAdd(key, word, tree.root)
	}
}

func doAdd(key [][]string, word string, current *bknode) bool {
	distance := Levenshtein(key, current.Key)

	if distance == 0 {
		current.Words = append(current.Words, word)
		return false
	}

	if len(current.Children) < distance {
		expandChildren := make([]bknode, distance)
		copy(expandChildren, current.Children)
		current.Children = expandChildren
	}

	next := &current.Children[distance-1]

	if next.Words == nil {
		current.Children[distance-1] = bknode{Key: key, Words: []string{word}}
		return true
	}

	return doAdd(key, word, next)
}

func (tree *BKTree) Find(key [][]string, threshold int) Matches {
	matches := Matches{}

	doFind(key, threshold, &matches, tree.root)

	return matches
}

func doFind(key [][]string, threshold int, matches *Matches, current *bknode) bool {
	distance := Levenshtein(key, current.Key)

	if distance <= threshold {
		for _, word := range current.Words {
			matches.Matches = append(matches.Matches, Match{Word: word, Distance: distance})
		}
	}

	if distance == 0 {
		return true
	}

	for i := max(distance-threshold-1, 0); i < min(distance+threshold, len(current.Children)); i++ {
		next := &current.Children[i]

		if next.Words != nil {
			if doFind(key, threshold, matches, next) {
				return true
			}
		}
	}

	return false
}

func (tree *BKTree) Serialize(writer io.Writer) error {
	encoder := gob.NewEncoder(writer)

	return encoder.Encode(tree.root)
}

func (tree *BKTree) Deserialize(reader io.Reader) error {
	tree.root = &bknode{}
	decoder := gob.NewDecoder(reader)

	return decoder.Decode(tree.root)
}
