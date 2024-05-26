## QUADCHECKER ##
== Grit:Lab May Picsine Go 2024==<br>
<br>
The purpose of this program is to check if an input matches any of the quad outputs (from the first raid).
There are 5 quad programs:<br>

<code>
// A)<br>
o---o<br>
|   |<br>
o---o<br>
<br>
// B)<br>
/***\<br>
*   *<br>
\***/<br>
<br>
// C)<br>
ABBBA<br>
B   B<br>
CBBBC<br>
<br>
// D)<br>
ABBBC<br>
B   B<br>
ABBBC<br>
<br>
// E)<br>
ABBBC<br>
B   B<br>
CBBBA<br>
</code>

Each quad function should take 2 int arguments: one for the row and one for the column of the quad generated.
The command to run the quad generatoirs are ./[program name] [row] [col]
Example: ./quadA 5 3, which will print out the examples above.<br>


The quadchecker(main.go) is intended to be run with the quad generator and a | command. This will feed the output of the quad generator into the quadchecker and tells the user which are the possible quads. Echo command can be used to feed incorrect output int the quadchecker.<br>


<code>
Correct Example:<br>
$ ./quadA 123 456 | ./quadchecker <br>
[quadA] [123] [456]<br>

Fail Example:<br>
$ echo -n "ABBo\n|  B\nB  |\nC--C\n"<br>
ABBo<br>
|  B<br>
B  |<br>
C--C<br>
$ echo -n "ABBo\n|  B\nB  |\nC--C\n" | ./quadchecker<br>
Not a quad function<br>
</code>

Made by:
Allen Lee 'ylee' Yuan Neng<br>

Tobias 'tjerner' Jernér<br>

Sean 'skipina' Kipinä


