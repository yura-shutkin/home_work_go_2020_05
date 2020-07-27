package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var EscapeCharList = []rune{
	'\\',
	'*',
	':',
}

func IsInList(myRune rune, listRunes []rune) bool {
	flag := false
	for _, val := range listRunes {
		if val == myRune {
			flag = true
			break
		}
	}
	return flag
}

func Unpack(msg string) (string, error) {
	var buf string
	var isPrevEsc bool
	resultString := ""

	for id, val := range msg {
		isPrevEsc = false
		// Check if first Char is Numeric
		if id == 0 {
			if unicode.IsDigit(val) {
				return "", ErrInvalidString
			}
		}
		// Check if letter
		if unicode.IsLetter(val) {
			if len(buf) != 0 {
				resultString += buf
				buf = string(val)
				continue
			}
			buf += string(val)
			continue
		}
		// Check if escape char
		if IsInList(val, EscapeCharList) {
			buf += string(val)
			continue
		}
		// Check if numeric
		if unicode.IsDigit(val) {
			if isPrevEsc {
				buf += string(val)
				continue
			}
			if len(buf) == 0 {
				return "", ErrInvalidString
			}
			resultString += strings.Repeat(buf, int(val-'0'))
			buf = ""
			continue
		}
		// Unknown char
		return "", ErrInvalidString
	}

	if len(buf) != 0 {
		resultString += buf
	}
	return resultString, nil
}
