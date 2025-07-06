package goreld

import (
	"strings"
)

func isnumber(s []rune) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}
	return true
}

func Putflag(s []rune) []rune {
	flg := []rune{}
	lsfl := false

	for i := 0; i < len(s); i++ {
		flg = append(flg, (s[i]))
		if s[i] == ')' {

			if i != len(s)-1 && s[i+1] != '\n' && s[i+1] != ' ' {
				return s
			}

			lsfl = true
			break
		} else if i != len(s)-1 && s[i+1] == '(' {
			return s
		}
	}

	if !lsfl {
		return s
	}
	for i := 0; i < len(flg); i++ {
		if flg[i] == ',' {
			if !(string(flg[:i]) == "(low" || string(flg[:i]) == "(cap" || string(flg[:i]) == "(up") {
				return s
			}
			if i != len(flg)-2 && !isnumber(flg[i+2:len(flg)-1]) {
				return s
			}
		}
	}

	return flg
}

func Split(f string) [][]string {
	sn := false
	s := []rune((f))
	str := ""
	a := []string{}
	m := [][]string{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' && !sn {
			if string(Putflag(s[i:])) != string((s[i:])) {
				str = string(Putflag((s[i:])))
				i += len([]rune(str)) - 1
				a = append(a, str)
				str = ""
			} else {
				i--
				sn = true
			}
		} else if s[i] != ' ' && s[i] != '\n' {
			str += string(s[i])
			sn = true
		} else if sn && (s[i] == ' ') {
			a = append(a, str)
			str = ""
			sn = false

		} else if s[i] == '\n' {
			a = append(a, str)

			m = append(m, Clnslc(a))
			str = ""
			a = []string{}
			sn = false
		}
	}
	if len(str) > 0 {
		a = append(a, str)
		m = append(m, Clnslc(a))
	}
	return m
}

func Slctostr(b []string) string {
	m := ""
	b = Vowel((b))
	for i := 0; i < len(b); i++ {

		m += (Clean(b[i]))
		if i != len(b)-1 && b[i] != "\n" {
			m += " "
		}

	}

	return (m)
}

func Clnslc(b []string) []string {
	m := []string{}
	for i := 0; i < len(b); i++ {
		if b[i] != "" {
			m = append(m, b[i])
		}
	}
	return m
}

func Clean(f string) string {
	sn := false
	str := ""
	s := []rune(f)
	if len(f) == 1 {
		return f
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
			sn = true
		} else if s[i] == ' ' && sn {
			str += " "
			sn = false
		}
	}

	return strings.TrimSpace(str)
}
