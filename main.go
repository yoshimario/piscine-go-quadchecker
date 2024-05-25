package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return // if number is negative, it will exit and not execute
	}
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == 1 && col == x) || (row == y && col == 1) || (row == y && col == x) {
				z01.PrintRune('o')
			} else if row == 1 || row == y {
				z01.PrintRune('-')
			} else if col == 1 || col == x {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadB(x, y int) {
	if x > 0 && y > 0 {

		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 && col == 1 {
					z01.PrintRune('/')
				} else if (row == 1 && col == x) || (row == y && col == 1) {
					z01.PrintRune('\\')
				} else if row == y && col == x {
					z01.PrintRune('/')
				} else if (row == 1 || row == y) && (col > 1 && col < x) {
					z01.PrintRune('*')
				} else if (col == 1 || col == x) && (row > 1 && row < y) {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func QuadC(x, y int) {
	if x > 0 && y > 0 {

		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 && col == 1 {
					z01.PrintRune('A')
				} else if row == 1 && col == x {
					z01.PrintRune('A')
				} else if row == y && col == x || (row == y && col == 1) {
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

func QuadE(x, y int) {
	if x > 0 && y > 0 {

		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 && col == 1 {
					z01.PrintRune('A')
				} else if row == 1 && col == x {
					z01.PrintRune('C')
				} else if row == y && col == 1 {
					z01.PrintRune('C')
				} else if row == y && col == x {
					z01.PrintRune('A')
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
