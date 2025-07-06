package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	helper "goreld/help"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: missing required arguments.")
		return
	}
	file, eror := os.Open(os.Args[1])
	if eror != nil {
		fmt.Println("Error can not open file")
		return
	}
	defer file.Close()
	var lines string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += (scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error :", err)
		return
	}
	// filee := strings.HasSuffix(os.Args[2], ".txt")
	if !strings.HasSuffix(os.Args[1], ".txt") || os.Args[1] == ".txt"||!strings.HasSuffix(os.Args[2], ".txt") || os.Args[2] == ".txt" {
		fmt.Println("Error: file name must be like this (<file name>.txt)")
		return
	}
	b := helper.Split((lines))
	b, err2 := helper.Flags(b)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	m := ""
	for i := 0; i < len(b); i++ {
		m += (helper.Clean(helper.Mark(helper.Punct(helper.Slctostr(b[i])))))
		if i != len(b)-1 {
			m += "\n"
		}
	}
	str4 := []byte(m)

	err3 := os.WriteFile(os.Args[2], str4, 0o644)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
}
