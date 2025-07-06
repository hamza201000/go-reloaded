package goreld

import (
	"strconv"
	"unicode"
)

func Cptlez(s string) string {
	slice := []rune(s)
	
	for i := 0; i < len(slice); i++ {
		if  i==0 {
			slice[i] = unicode.ToUpper(slice[i])
			
		} else  {
			slice[i] = unicode.ToLower(slice[i])
		}
	}
	return string(slice)
}

func Check(s string) (int, error,string) {
	f := ""
	var err error
	if s[0] != '(' || s[len(s)-1] != ')' {
		return 0,err, "invalid flag"
	} else {
		f = (s[1 : len(s)-1])

	}
	for i := 0; i < len(s); i++ {
		if f[i] == ',' {

			m, err := strconv.Atoi(f[i+2:])
			if err != nil {
				return 0, err,""
			}
			return m, err,""
		}
	}
	return 0, err,""
}
