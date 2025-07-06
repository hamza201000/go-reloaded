package goreld

import (
	"strconv"
	"strings"
	
)

func Flags(b [][]string) ([][]string, error) {
	var err error
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == "(hex)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				dec, err := strconv.ParseInt(b[i][j-1], 16, 64)
				b[i][j] = ""
				b[i][j-1] = strconv.Itoa(int(dec))

				if err != nil {
					continue
				}

				b[i] = Clnslc(b[i])
				j--
			} else if b[i][j] == "(bin)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				dec, err := strconv.ParseInt(b[i][j-1], 2, 64)
				b[i][j] = ""
				b[i][j-1] = strconv.Itoa(int(dec))
				if err != nil {
					continue
				}
				b[i] = Clnslc(b[i])
				j--
			} else if b[i][j] == "(up)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				b[i][j] = ""
				b[i][j-1] = strings.ToUpper(b[i][j-1])

				b[i] = Clnslc(b[i])
				j--
			} else if b[i][j] == "(low)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				b[i][j] = ""
				b[i][j-1] = strings.ToLower(b[i][j-1])
				b[i] = Clnslc(b[i])
				
				b[i] = Clnslc(b[i])
				j--
				
			} else if b[i][j] == "(cap)" {
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				b[i][j] = ""
				b[i][j-1] = Cptlez(b[i][j-1])

				b[i] = Clnslc(b[i])
				j--
			} else if len(b[i][j])>=5&&b[i][j][0:5]== "(up, " {
				if j == 0 {
					b[i][j] = ""
					b[i] = Clnslc(b[i])
					continue
				}
				f, nr,vld := Check(b[i][j])
				if vld!=""{
					continue
				}
				if nr != nil {
					continue

				}

				if f <= 0 {
					b[i][j] = ""
					b[i] = Clnslc(b[i])
					continue
				}
				b[i][j] = ""
				for k := 1; k <= f; k++ {
					if k > j {
						break
					}
					b[i][j-k] = strings.ToUpper(b[i][j-k])
				}
				b[i] = Clnslc(b[i])
				j--
			} else if len(b[i][j])>=6&& b[i][j][0:6]== "(low, "{
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				f, nr,vld := Check(b[i][j])
				if vld!=""{
					continue
				}
				if nr != nil {
					continue

				}
				if f <= 0 {
					b[i][j] = ""
					b[i] = Clnslc(b[i])
					continue
				}
				b[i][j] = ""
				for k := 1; k <= f; k++ {
					if k > j {
						break
					}
					b[i][j-k] = strings.ToLower(b[i][j-k])
				}
				b[i] = Clnslc(b[i])
				j--
			} else if len(b[i][j])>=6&& b[i][j][0:6]== "(cap, " {
				if j == 0 {
					b[i][j] = ""

					b[i] = Clnslc(b[i])
					continue
				}
				f, nr,vld := Check(b[i][j])
				if vld!=""{
					continue
				}
				if nr != nil {
					continue

				}
				if f <= 0 {
					b[i][j] = ""
					b[i] = Clnslc(b[i])
					continue
				}
				b[i][j] = ""

				for k := 1; k <= f; k++ {
					if k > j {
						break
					}
					b[i][j-k] = Cptlez(b[i][j-k])

				}
				b[i] = Clnslc(b[i])
				j--
			}
		}
	}
	return b, err
}
