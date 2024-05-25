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
	QuadD(row, col)
	os.Exit(0)
}

func atoi(s string) int {
	num := 0
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		num = num*10 + int(s[i]-48)
	}
	return num
}

func QuadD(x, y int) {
	if x > 0 && y > 0 {

		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if (row == 1 && col == 1) || (row == y && col == 1) {
					z01.PrintRune('A')
				} else if (row == 1 && col == x) || (row == y && col == x) {
					z01.PrintRune('C')
				} else if (row == 1 || row == y) && (col > 1 && col < x) {
					z01.PrintRune('B')
				} else if (col == 1 || col == x) && (row > 1 && row < y) {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
