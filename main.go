package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// creates buffer to store input and temp byte of slice to read the chunks of input
	var buffer bytes.Buffer
	tempBuf := make([]byte, 1024)

	//loops and write the input by chunks into the buffer until there is no more input
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

	//convert the buffer to string and then check the length of it.
	inputB := buffer.Bytes()
	inputS := string(inputB)
	inputSLen := len(inputB)

	if inputS == "" {
		os.Exit(0)
	}

	//count the columns and the rows
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

	// checks the predefined pattern if the input pattern matches, i.e. QuadA have o's dashes and poles. The code repeats for the rest aswell from QuadB -  QuadE, last part is just if its not a quad function, it will print that.
	// ** If numRes == 0 and result = result + in quad D and E is because D-E share the same characters as C as thats the first one we check it doesn't need it.
	result := ""
	numRes := 0
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

// Turns int to string
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

// generate a string reprensentation of a box with the given dimensions and characters.
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

// generate a single line of the box.
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
