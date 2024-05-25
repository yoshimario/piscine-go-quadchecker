package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	tempBuf := make([]byte, 1024)
	for {
		count, err := os.Stdin.Read(tempBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Exit(1)
		}
		buffer.Write(tempBuf[:count])

	}

	inputB := buffer.Bytes()

	/* inputStat, _ := os.Stdin.Stat()          // grab info of output of previous command
	inputB := make([]byte, inputStat.Size()) // create array of correct size for output of previous command
	os.Stdin.Read(inputB)                    // grab output into array inputB */

	inputS := string(inputB) // convert input byte to input string
	inputSLen := len(inputB) // get length of input

	// check the number of rows and col
	col := 0
	row := 0
	for i := 0; i < inputSLen; i++ {
		if inputS[i] != '\n' {
			col++
		} else {
			break
		}
	}

	for i := 0; i < inputSLen; i++ {
		if inputS[i] == '\n' {
			row++
		}
	}

	// generate quadABCDE and compare with input to see if it matches any
	result := "" // empty string to store result
	numRes := 0  // count number of results
	if inputS == prntBox(col, row, 'o', 'o', 'o', 'o', '-', '|') {
		result = "[quadA] [" + itoa(col) + "] [" + itoa(row) + "]"
		numRes++
	}
	if inputS == prntBox(col, row, '/', '\\', '\\', '/', '*', '*') {
		result = "[quadB] [" + itoa(col) + "] [" + itoa(row) + "]"
		numRes++
	}
	if inputS == prntBox(col, row, 'A', 'A', 'C', 'C', 'B', 'B') {
		result = "[quadC] [" + itoa(col) + "] [" + itoa(row) + "]"
		numRes++
	}
	if inputS == prntBox(col, row, 'A', 'C', 'A', 'C', 'B', 'B') {
		if numRes == 0 {
			result = "[quadD] [" + itoa(col) + "] [" + itoa(row) + "]"
		} else {
			result = result + " || " + "[quadD] [" + itoa(col) + "] [" + itoa(row) + "]"
		}
		numRes++
	}
	if inputS == prntBox(col, row, 'A', 'C', 'C', 'A', 'B', 'B') {
		if numRes == 0 {
			result = "[quadE] [" + itoa(col) + "] [" + itoa(row) + "]"
		} else {
			result = result + " || " + "[quadE] [" + itoa(col) + "] [" + itoa(row) + "]"
		}
		numRes++
	}

	if numRes == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(result)
	}
	os.Exit(0)
}

func itoa(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}
	for num > 0 {
		s = string(rune(num%10+48)) + s
		num = num / 10
	}

	return s
}

func prntBox(x, y int, c1, c2, c3, c4, top, side rune) string {
	s := ""

	for i := 0; i < y; i++ {
		if i == 0 {
			s = s + prntLine(x, c1, c2, top)
		} else if i == y-1 {
			s = s + prntLine(x, c3, c4, top)
		} else {
			s = s + prntLine(x, side, side, ' ')
		}
	}

	return s
}

func prntLine(x int, c1 rune, c2 rune, mid rune) string {
	s := ""
	for i := 0; i < x; i++ {
		if i == 0 {
			s = s + string(c1)
		} else if i == x-1 {
			s = s + string(c2)
		} else {
			s = s + string(mid)
		}
	}
	s = s + string('\n')

	return s
}
