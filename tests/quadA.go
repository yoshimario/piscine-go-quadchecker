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
 QuadA(col, row)
 os.Exit(0)
}

func atoi(s string) int {
 num := 0
 isMinus := false

 if s[0] == '-' {
 isMinus = true
 s = s[1:]
 }

 sLen := len(s)
 for i := 0; i < sLen; i++ {
 num = num*10 + int(s[i]-48)
 }

 if isMinus {
 num = -num
 }
 return num
}

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