** QUADCHECKER **
== Grit:Lab May Picsine Go 2024==

The purpose of this program is to check if an input matches any of the quad outputs (from the first raid).
There are 5 quad programs:

A)
o---o
|   |
o---o

B)
/***\
*   *
\***/

C)
ABBBA
B   B
CBBBC

D)
ABBBC
B   B
ABBBC

E)
ABBBC
B   B
CBBBA

Each quad function should take 2 int arguments: one for the row and one for the column of the quad generated.
The command to run the quad generatoirs are ./[program name] [row] [col]
Example: ./quadA 5 3, which will print out the examples above.

The quadchecker(main.go) is intended to be run with the quad generator and a | command. This will feed the output of the quad generator into the quadchecker and tells the user which are the possible quads. Echo command can be used to feed incorrect output int the quadchecker.

Correct Example:
$ ./quadA 123 456 | ./quadchecker 
[quadA] [123] [456]

Fail Example:
$ echo -n "ABBo\n|  B\nB  |\nC--C\n"
ABBo
|  B
B  |
C--C
$ echo -n "ABBo\n|  B\nB  |\nC--C\n" | ./quadchecker
Not a quad function

Made by:
Allen Lee 'ylee' Yuan Neng
Tobias 'tjerner' Jernér
Sean 'skipina' Kipinä


