package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"strings"
)

func MassReplace(str, search, replace string) string {
	for {
		if strings.Contains(str, search) {
			str = strings.ReplaceAll(str, search, replace)
		} else {
			return str
		}
	}
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
