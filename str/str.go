package str

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

func TrimStringSlice(raw []string) []string {
	if raw == nil {
		return []string{}
	}

	cnt := len(raw)
	arr := make([]string, 0, cnt)
	for i := 0; i < cnt; i++ {
		item := strings.TrimSpace(raw[i])
		if item == "" {
			continue
		}

		arr = append(arr, item)
	}

	return arr
}

func to(fn func(rune) rune, str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	var bf bytes.Buffer
	var err error

	for _, r := range str {
		_, err = bf.WriteRune(fn(r))
		if err != nil {
			break
		}
	}
	return bf.String(), err
}

// ToUpper convert a string to uppercase
func ToUpper(str string) string {
	s, _ := to(unicode.ToUpper, str)
	return s
}

// ToLower convert a string to lowercase
func ToLower(str string) string {
	s, _ := to(unicode.ToLower, str)
	return s
}

// ToTitle title string, initial char of the string to uppercase
func ToTitle(str string) string {
	if IsEmpty(str) {
		return str
	}
	runes := []rune(str)
	runes[0] = unicode.ToTitle(runes[0])
	return string(runes)
}

// Reverse reverse a string
func Reverse(str string) string {
	// special case
	if IsEmpty(str) {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	var bf bytes.Buffer

	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		bf.WriteRune(r)
		str = str[:len(str)-size]
	}
	return bf.String()
}

// Repeat returns a new string consisting of count copies of the string str
func Repeat(str string, count int) string {
	if IsEmpty(str) || count == 0 {
		return str
	}
	return strings.Repeat(str, count)
}

// IsEmpty returns a boolean of s equal "" or len(s) == 0
func IsEmpty(str string) bool {
	return "" == str || 0 == len(str)
}
