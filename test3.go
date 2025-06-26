package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkpun2(b []string) string {
	m := ""
	for i := 0; i < len(b); i++ {
		m += (clean(b[i]))
		if i != len(b)-1 && b[i] != "\n" && b[i+1] != "\n" {
			m += " "
		}

	}

	return (m)
}

func split(f string) [][]string {
	sn := false
	s := []rune((f))
	str := ""
	a := []string{}
	m := [][]string{}
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {

			a = append(a, str)
			// a = append(a, "\n")
			m = append(m, a)
			str = ""
			a = []string{}
			sn = false

		} else if s[i] != ' ' && s[i] != '\n' {
			str += string(s[i])
			sn = true
		} else if sn && (s[i] == ' ') {
			a = append(a, str)
			str = ""
			sn = false
		}
	}
	if sn {
		a = append(a, str)
		m = append(m, a)
	}
	return m
}

func clean(f string) string {
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
	file, eror := os.Open(os.Args[1])
	if eror != nil {
		fmt.Println("Error can not open file")
	}
	defer file.Close()
	var lines string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += (scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
	}

	// fmt.Println(lines)
	//s := ""
	/*for i := 0; i < len(lines); i++ {

		s += lines[i]
		if i != len(lines) {
			s += "\n"
		}

	}*/

	// a := "(up, 2) (low,  2) (up, 2)  (up)  "

	b := split((lines))

	h := false
	// m := ""
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == "(hex)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = cln(b[i])
					continue
				}
				dec, err := strconv.ParseInt(b[i][j-1], 16, 64)
				b[i][j] = ""
				b[i][j-1] = strconv.Itoa(int(dec))
				// b[i][j] = clean2(b[i][j])
				if err != nil && h {
					b[i][j] = ""
					b[i][j-1] = ""
				} else if err != nil {
					fmt.Println("Error:", err)
					return
				}
				h = true
				// fmt.Println(dec)
				b[i] = cln(b[i])
				j--
			} else if b[i][j] == "(bin)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = cln(b[i])
					continue
				}
				dec, err := strconv.ParseInt(b[i][j-1], 2, 64)
				b[i][j] = ""
				b[i][j-1] = strconv.Itoa(int(dec))
				// fmt.Println(b[i+1])
				if err != nil && h {
					b[i][j] = ""
					b[i][j-1] = ""
				} else if err != nil {
					fmt.Println("Error:", err)
					return
				}
				h = true
				// fmt.Println(dec)
				b[i] = cln(b[i])
				j--
			} else if b[i][j] == "(up)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = cln(b[i])
					continue
				}
				// fmt.Println(b[i][j-1])
				b[i][j] = ""
				b[i][j-1] = strings.ToUpper(b[i][j-1])
				// fmt.Println(b[i+1])

				h = true
				b[i] = cln(b[i])
				j--
			} else if b[i][j] == "(low)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = cln(b[i])
					continue
				}
				// fmt.Println(b)
				b[i][j] = ""
				b[i][j-1] = strings.ToLower(b[i][j-1])
				// fmt.Println(b)
				// fmt.Println(b[i+1])

				h = true
				// fmt.Println(b[i][j-1])
				b[i] = cln(b[i])
				j--
			} else if b[i][j] == "(cap)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = cln(b[i])
					continue
				}
				b[i][j] = ""
				b[i][j-1] = cptlez(b[i][j-1])
				// fmt.Println(b[i+1])

				h = true
				// fmt.Println(b[i][j-1])
				b[i] = cln(b[i])
				j--
			}
		}
	}

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			/*if i != len(b)-1 && b[i][j][0] == '(' && b[i+1] == "," {
				// fmt.Println(b[i+1][len(b[i+1])-1])
				b[i][j] = check2(b[i][j]) + b[i+1]
				b[i+1] = ""
				b = cln2(b)
				// fmt.Println(b[i][j])

			}*/

			fmt.Println(b[i][j])
			if j != len(b[i])-1 && !isvalisflag(b[i][j], b[i][j+1]) {
				continue
			}
			if j != len(b[i])-1 && b[i][j] == "(up," {
				if j == 0 {
					b[i][j] = ""
					b[i][j+1] = ""
					b[i] = cln(b[i])
					continue
				}
				f := check(b[i][j] + b[i][j+1])
				if f <= 0 {
					continue
				}
				b[i][j] = ""
				b[i][j+1] = ""
				for k := 1; k <= f; k++ {
					if k > j {
						break
					}
					b[i][j-k] = strings.ToUpper(b[i][j-k])

				}

				// fmt.Println(f)

				h = true
				b[i] = cln(b[i])
				j--
				// fmt.Println(b[i][j-1])
			} else if j != len(b[i])-1 && b[i][j] == "(low," {
				if j == 0 {
					b[i][j] = ""
					b[i][j+1] = ""
					b[i] = cln(b[i])
					continue
				}
				f := check(b[i][j] + b[i][j+1])
				// fmt.Println(strings.Split(b[i][j+1], ")")[0])
				if f <= 0 {
					continue
				}
				b[i][j] = ""
				b[i][j+1] = (strings.Split(b[i][j+1], ")")[1])
				for k := 1; k <= f; k++ {
					if k > j {
						break
					}
					b[i][j-k] = strings.ToLower(b[i][j-k])

				}

				// fmt.Println(f)

				h = true
				b[i] = cln(b[i])
				j--
				// fmt.Println(b[i][j-1])
			} else if j != len(b[i])-1 && b[i][j] == "(cap," {
				if i == 0 {
					b[i][j] = ""
					b[i][j+1] = ""
					b[i] = cln(b[i])
					continue
				}
				f := check(b[i][j] + b[i][j+1])
				if f <= 0 {
					continue
				}
				b[i][j] = ""
				b[i][j+1] = (strings.Split(b[i][j+1], ")")[1])
				for k := 1; k <= f; k++ {
					if k > j {
						break
					}
					b[i][j-k] = cptlez(b[i][j-k])

				}
				// fmt.Println(f)
				h = true
				b[i] = cln(b[i])
				j--
				// fmt.Println(b[i][j-1])
			}
		}
	}
	// fmt.Println(b[1])
	m := ""
	for i := 0; i < len(b); i++ {
		m += checkpun2(b[i])
		if i != len(b)-1 {
			m += "\n"
		}
	}

	//fmt.Println(b[0])
	//fmt.Println(m)
	/*m := ""
	for i := 0; i < len(b); i++ {
		if b[i][j] != "" {
			m += b[i]
			if i != len(b)-1 {
				m += " "
			}
		}
	}*/
	fam := (mark(m))
	fam = clean((m))
	str4 := []byte((check2(fam)))
	err := os.WriteFile(os.Args[2], str4, 0o644)
	if err != nil {
		fmt.Println("Error2")
	}
	// fmt.Print((str4))
}
