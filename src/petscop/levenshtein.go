package petscop

func min3(x int, y int, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}

func Levenshtein(s [][]string, t [][]string) int {
	lengthS := len(s)
	lengthT := len(t)

	previous := make([]int, lengthT+1)
	current := make([]int, lengthT+1)

	for i := 0; i < lengthT+1; i++ {
		previous[i] = i
	}

	for i := 0; i < lengthS; i++ {
		current[0] = i + 1

		for j := 0; j < lengthT; j++ {
			insertions := current[j] + 1
			deletions := previous[j+1] + 1
			substitutions := previous[j]

			commonality := false
			for _, phoneme1 := range s[i] {
				for _, phoneme2 := range t[j] {
					if phoneme1 == phoneme2 {
						commonality = true
						break
					}
				}
			}

			if !commonality {
				substitutions++
			}

			current[j+1] = min3(insertions, deletions, substitutions)
		}

		copy(previous, current)
	}

	return previous[lengthT]
}
