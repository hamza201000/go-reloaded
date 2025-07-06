package goreld

import (
	"strings"
	"unicode"
)

func Mark(f string) string {
	s := []rune(Clean(f))
	b := false
	m := []rune{}
	for i := 0; i < len(s); i++ {
		if i != 0 && i != len(s)-1 && s[i] == '\'' && unicode.IsLetter(s[i+1]) && unicode.IsLetter(s[i-1]) {
			m = append(m, s[i])
			continue
		}
		if i != len(s)-1 && s[i] == '\'' && !b {
			b = true
			if i != 0 {
				m = append(m, ' ')
			}

			m = append(m, s[i])
			if s[i+1] == ' ' {
				i++
			}

		} else if s[i] == ' ' && s[i+1] == '\'' && b {
			continue
		} else if s[i] == '\'' && b {
			b = false
			m = append(m, s[i])
			m = append(m, ' ')
		} else {
			m = append(m, s[i])
		}
	}
	return string(m)
}

func Punct(f string) string {
	s := Clean(f)
	str := []rune(s)
	str2 := []rune{}
	m := 0
	for i := 0; i < len(str)-1; i++ {
		if i != 0 && (str[i] == '.' || str[i] == ',' || str[i] == '?' || str[i] == '!' || str[i] == ':' || str[i] == ';') {
			if str[i-1] == ' ' && !(str[i-1] == '.' || str[i-1] == ',' || str[i-1] == '?' || str[i-1] == '!' || str[i-1] == ':' || str[i-1] == ';') {
				str2 = append(str2, str[m:i-1]...)

				m = i

			}
			if str[i+1] != ' ' && !(str[i+1] == '.' || str[i+1] == ',' || str[i+1] == '?' || str[i+1] == '!' || str[i+1] == ':' || str[i+1] == ';') {
				str2 = append(str2, str[m:i+1]...)
				str2 = append(str2, ' ')

				m = i + 1
			}
		}
		if str[i] == '.' || str[i] == ',' || str[i] == '?' || str[i] == '!' || str[i] == ':' || str[i] == ';' {
			if i == 0 && str[i+1] != ' ' && !(str[i+1] == '.' || str[i+1] == ',' || str[i+1] == '?' || str[i+1] == '!' || str[i+1] == ':' || str[i+1] == ';') {
				str2 = append(str2, str[m:i+1]...)
				str2 = append(str2, ' ')

				m = i + 1
			}
		}
	}
	str2 = append(str2, str[m:]...)
	if len(str2) > 1 && (str2[len(str2)-2] == ' ') && (str2[len(str2)-1] == '.' || str2[len(str2)-1] == ',' || str2[len(str2)-1] == '?' || str2[len(str2)-1] == '!' || str2[len(str2)-1] == ':' || str2[len(str2)-1] == ';') {
		str2[len(str2)-2] = str2[len(str2)-1]
		str2[len(str2)-1] = ' '
	}
	return Clean(string(str2))
}

func Vowel(s []string) []string {
	str2 := []string{}
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && (s[i] == "a" || s[i] == "A") && strings.Contains("aeiouahAEIOUAH", string(s[i+1][0])) {
			str2 = append(str2, s[i]+"n")
		} else {
			str2 = append(str2, s[i])
		}
	}
	return (str2)
}
