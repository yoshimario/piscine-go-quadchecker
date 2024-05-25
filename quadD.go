package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen != 2 {
		return
	}

	col := atoi(args[0])
	row := atoi(args[1])
	printBox(col, row, 'A', 'C', 'A', 'C', 'B', 'B')
}

func atoi(s string) int {
	num := 0
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		num = num*10 + int(s[i]-48)
	}
	return num
}

func printBox(x, y int, c1, c2, c3, c4, top, side rune) {

	if x < 1 || y < 1 {
		return
	}

	for i := 0; i < y; i++ {
		if i == 0 {
			printLine(x, c1, c2, top)
		} else if i == y-1 {
			printLine(x, c3, c4, top)
		} else {
			printLine(x, side, side, ' ')
		}
	}
	z01.PrintRune('\n')
}

func printLine(x int, c1 rune, c2 rune, mid rune) {
	for i := 0; i < x; i++ {
		if i == 0 {
			z01.PrintRune(c1)
		} else if i == x-1 {
			z01.PrintRune(c2)
		} else {
			z01.PrintRune(mid)
		}
	}
	z01.PrintRune('\n')
}
