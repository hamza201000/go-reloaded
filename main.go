package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkpun2(b []string) string {
	m := ""
	for i := 0; i < len(b); i++ {

		m += (b[i])
		if i != len(b)-1 {
			m += " "
		}
	}

	return m
}

func split(f string) []string {
	sn := false
	s := []rune(f)
	str := ""
	a := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
			sn = true
		} else if sn && s[i] == ' ' {
			a = append(a, str)
			str = ""
			sn = false
		}
	}
	if sn {
		a = append(a, str)
	}
	return a
}

func clean(f string) string {
	sn := false
	str := ""
	s := []rune(f)
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


func check(s string) int {
	f := ""

	if s[0] != '(' || s[len(s)-1] != ')' {
		return 0
	} else {
		f = (s[1 : len(s)-1])
		// return string(f)
	}

	// wor := ""
	for i := 0; i < len(s); i++ {
		if f[i] == ',' {
			// wor := f[0:i]
			m, err := strconv.Atoi(f[i+1:])
			if err != nil {
				return 0
			}
			return m
		}
	}
	return 0
}

func cln(b []string) []string {
	m := []string{}
	for i := 0; i < len(b); i++ {
		if b[i] != "" {
			m = append(m, b[i])
		}
	}
	return m
}

func cptlez(s string) string {
	slice := []rune(s)
	flag := false

	for i := 0; i < len(slice); i++ {
		if unicode.IsLetter(slice[i]) && !flag {
			slice[i] = unicode.ToUpper(slice[i])
			flag = true
		} else if flag {
			slice[i] = unicode.ToLower(slice[i])
		}
	}
	return string(slice)
}

func check2(f string) string {
	s := clean(f)
	str := []rune(s)
	str2 := []rune{}
	m := 0
	for i := 0; i < len(str)-1; i++ {
		if i != 0 && (str[i] == '.' || str[i] == ',' || str[i] == '?' || str[i] == '!' || str[i] == ':' || str[i] == ';') {
			if str[i-1] == ' ' && !(str[i-1] == '.' || str[i-1] == ',' || str[i-1] == '?' || str[i-1] == '!' || str[i-1] == ':' || str[i-1] == ';') {
				str2 = append(str2, str[m:i-1]...)
				// str2 = append(str2, str[i])
				// str2 = append(str2, ' ')
				// fmt.Println(string(str2))
				m = i

			}
			if str[i+1] != ' ' && !(str[i+1] == '.' || str[i+1] == ',' || str[i+1] == '?' || str[i+1] == '!' || str[i+1] == ':' || str[i+1] == ';') {
				str2 = append(str2, str[m:i+1]...)
				str2 = append(str2, ' ')
				// fmt.Println(string(str2))
				m = i + 1
			}
		}
		if str[i] == '.' || str[i] == ',' || str[i] == '?' || str[i] == '!' || str[i] == ':' || str[i] == ';' {
			if i == 0 && str[i+1] != ' ' && !(str[i+1] == '.' || str[i+1] == ',' || str[i+1] == '?' || str[i+1] == '!' || str[i+1] == ':' || str[i+1] == ';') {
				str2 = append(str2, str[m:i+1]...)
				str2 = append(str2, ' ')
				// fmt.Println(string(str2))
				m = i + 1
			}
		}
	}
	str2 = append(str2, str[m:]...)
	if len(str2) > 1 && (str2[len(str2)-2] == ' ') && (str2[len(str2)-1] == '.' || str2[len(str2)-1] == ',' || str2[len(str2)-1] == '?' || str2[len(str2)-1] == '!' || str2[len(str2)-1] == ':' || str2[len(str2)-1] == ';') {
		str2[len(str2)-2] = str2[len(str2)-1]
		str2[len(str2)-1] = ' '
	}
	return clean(string(str2))
}

func isvalisflag(s, m string) bool {
	f := clean(s + m)
	if f[0] != '(' || f[len(f)-1] != ')' {
		return false
	}
	return true
}

func mark(f string) string {
	s := []rune(clean(f))
	b := false
	m := []rune{}
	for i := 0; i < len(s); i++ {
		if s[i] == '\'' && !b && i != len(s)-1 {
			b = true
			m = append(m, ' ')
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

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error1")
	}

	// a := "(up, 2) (low,  2) (up, 2)  (up)  "

	b := split(string(data))

	h := false
	// m := ""
	for i := 0; i < len(b); i++ {
		if b[i] == "(hex)" {
			if i == 0 {
				b[i] = ""

				b = cln(b)
				continue
			}
			dec, err := strconv.ParseInt(b[i-1], 16, 64)
			b[i] = ""
			b[i-1] = strconv.Itoa(int(dec))
			// b[i] = clean2(b[i])
			if err != nil && h {
				b[i] = ""
				b[i-1] = ""
			} else if err != nil {
				fmt.Println("Error:", err)
				return
			}
			h = true
			// fmt.Println(dec)
			b = cln(b)
			i--
		} else if i > 0 && b[i] == "(bin)" {
			if i == 0 {
				b[i] = ""

				b = cln(b)
				continue
			}
			dec, err := strconv.ParseInt(b[i-1], 2, 64)
			b[i] = ""
			b[i-1] = strconv.Itoa(int(dec))
			// fmt.Println(b[i+1])
			if err != nil && h {
				b[i] = ""
				b[i-1] = ""
			} else if err != nil {
				fmt.Println("Error:", err)
				return
			}
			h = true
			// fmt.Println(dec)
			b = cln(b)
			i--
		} else if b[i] == "(up)" {
			if i == 0 {
				b[i] = ""

				b = cln(b)
				continue
			}
			// fmt.Println(b[i-1])
			b[i] = ""
			b[i-1] = strings.ToUpper(b[i-1])
			// fmt.Println(b[i+1])

			h = true
			b = cln(b)
			i--
		} else if b[i] == "(low)" {
			if i == 0 {
				b[i] = ""

				b = cln(b)
				continue
			}
			// fmt.Println(b)
			b[i] = ""
			b[i-1] = strings.ToLower(b[i-1])
			// fmt.Println(b)
			// fmt.Println(b[i+1])

			h = true
			// fmt.Println(b[i-1])
			b = cln(b)
			i--
		} else if b[i] == "(cap)" {
			if i == 0 {
				b[i] = ""

				b = cln(b)
				continue
			}
			b[i] = ""
			b[i-1] = cptlez(b[i-1])
			// fmt.Println(b[i+1])

			h = true
			// fmt.Println(b[i-1])
			b = cln(b)
			i--
		}
	}

	for i := 0; i < len(b); i++ {

		/*if i != len(b)-1 && b[i][0] == '(' && b[i+1] == "," {
			// fmt.Println(b[i+1][len(b[i+1])-1])
			b[i] = check2(b[i]) + b[i+1]
			b[i+1] = ""
			b = cln2(b)
			// fmt.Println(b[i])

		}*/

		fmt.Println(b[i])
		if i != len(b)-1 && !isvalisflag(b[i], b[i+1]) {
			continue
		}
		if b[i] == "(up," && i != len(b)-1 {
			if i == 0 {
				b[i] = ""
				b[i+1] = ""
				b = cln(b)
				continue
			}
			f := check(b[i] + b[i+1])
			if f <= 0 {
				continue
			}
			b[i] = ""
			b[i+1] = ""
			for j := 1; j <= f; j++ {
				if j > i {
					break
				}
				b[i-j] = strings.ToUpper(b[i-j])

			}

			// fmt.Println(f)

			h = true
			b = cln(b)
			i--
			// fmt.Println(b[i-1])
		} else if b[i] == "(low," && i != len(b)-1 {
			if i == 0 {
				b[i] = ""
				b[i+1] = ""
				b = cln(b)
				continue
			}
			f := check(b[i] + b[i+1])
			// fmt.Println(strings.Split(b[i+1], ")")[0])
			if f <= 0 {
				continue
			}
			b[i] = ""
			b[i+1] = (strings.Split(b[i+1], ")")[1])
			for j := 1; j <= f; j++ {
				if j > i {
					break
				}
				b[i-j] = strings.ToLower(b[i-j])
			}

			// fmt.Println(f)

			h = true
			b = cln(b)
			i--
			// fmt.Println(b[i-1])
		} else if b[i] == "(cap," && i != len(b)-1 {
			if i == 0 {
				b[i] = ""
				b[i+1] = ""
				b = cln(b)
				continue
			}
			f := check(b[i] + b[i+1])
			if f <= 0 {
				continue
			}
			b[i] = ""
			b[i+1] = (strings.Split(b[i+1], ")")[1])
			for j := 1; j <= f; j++ {
				if j > i {
					break
				}
				b[i-j] = cptlez(b[i-j])
			}
			// fmt.Println(f)
			h = true
			b = cln(b)
			i--
			// fmt.Println(b[i-1])
		}
	}
	fmt.Println(b[0])
	m := checkpun2(b)
	fmt.Println(b[0])
	fmt.Println(m)
	/*m := ""
	for i := 0; i < len(b); i++ {
		if b[i] != "" {
			m += b[i]
			if i != len(b)-1 {
				m += " "
			}
		}
	}*/
	f := mark(m)
	str4 := []byte(clean(check2(f)))
	err = os.WriteFile(os.Args[2], str4, 0o644)
	if err != nil {
		fmt.Println("Error2")
	}
	// fmt.Print((str4))
}
