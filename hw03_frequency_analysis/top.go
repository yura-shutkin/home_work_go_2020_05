package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"math/rand"
	"strings"
	"time"
)

// I had a situation when a single run did not remove every occurrence of "search"
// That's why I run infinite loop that guaranties every "search" will be replaced
// with "replace"
func MassReplace(str, search, replace string) string {
	for {
		if strings.Contains(str, search) {
			str = strings.ReplaceAll(str, search, replace)
		} else {
			return str
		}
	}
}

func GenerateText(words map[string]int) string {
	randSrc := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(randSrc)
	var text string

	wordListLen := 0
	for _, val := range words {
		wordListLen += val
	}

	wordList := make([]string, wordListLen)

	position := 0
	for word, count := range words {
		for i := position; i < position+count; i++ {
			wordList[i] = word
		}
		position += count
	}

	randomizer.Shuffle(wordListLen, func(i, j int) {
		wordList[i], wordList[j] = wordList[j], wordList[i]
	})

	text = strings.Join(wordList, " ")

	return text
}

func Top10(msg string) []string {
	if len(msg) == 0 {
		return nil
	}
	ratings := make(map[string]int)
	msg = MassReplace(msg, "\t", " ")
	msg = MassReplace(msg, "\n", " ")
	msg = MassReplace(msg, "  ", " ")

	raw := strings.Split(msg, " ")

	for _, val := range raw {
		_, ok := ratings[val]
		if !ok {
			ratings[val] = 1
		} else {
			ratings[val] += 1
		}
	}

	var top []string
	for lim := 10; lim > 0; lim -= 1 {
		max := 0
		for _, val := range ratings {
			if val > max {
				max = val
			}
		}
		for key, val := range ratings {
			if val == max {
				top = append(top, key)
				delete(ratings, key)
				break
			}
		}
	}

	return top
}
