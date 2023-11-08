package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func ReadInput() {
	data := make([]byte, 1024)
	
	for {
		size, err := os.Stdin.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			PrintStr("ERROR: reading standard input")
			return
		}
		PrintStr(string(data[:size]))
	}
}

func PrintStr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 1 {
		ReadInput()
	} else {
		args := os.Args[1:]
		for _, filename := range args {
			ReadFile(filename)
		}
	}
}

func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		PrintStr("ERROR: open " + filename + ": no such file or directory")
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	for {
		size, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			PrintStr("ERROR: reading " + filename)
			return
		}
		PrintStr(string(data[:size]))
	}
}
