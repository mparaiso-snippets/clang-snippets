package utils

type Alphabet struct {
	str []rune
}

func NewAlphabet(s string) *Alphabet {
	alphabet := &Alphabet{}
	index := 0
	runeMap := map[rune]int{}
	for _, Rune := range s {
		if _, ok := runeMap[Rune]; !ok {
			runeMap[Rune] = index
			index++
		}
	}
	alphabet.str = make([]rune, len(runeMap))
	for Rune, index := range runeMap {
		alphabet.str[index] = Rune
	}
	return alphabet
}

func (alphabet *Alphabet) ToChar(index int) rune {
	return alphabet.str[index]
}
func (alphabet *Alphabet) ToIndex(char rune) int {
	for index, Rune := range alphabet.str {
		if Rune == char {
			return index
		}
	}
	return -1
}
func (alphabet *Alphabet) Contains(char rune) bool { return alphabet.ToIndex(char) >= 0 }
func (alphabet *Alphabet) R() int                  { return len(alphabet.str) }

// func (alphabet *Alphabet) LgR() int                     { return 0 }
// func (alphabet *Alphabet) ToIndices(s string) []int     { return []int{} }
// func (alphabet *Alphabet) ToChars(indices []int) string { return "" }
