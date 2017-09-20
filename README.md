#Arithmetic Library
By Marko Mrdja - SENG 560

The library is simple add, subtract, multiply, delete, convert

To see usage, look at the main class.

The functions that take and return strings, return a string in the format of the base that you are converting to or doing arithmetic with.  Leading zeros are trimmed. I opted to use and return strings for ease of development.

The functions are:
~~~
arith.AddFloat(a float64, b float64) returns float64

arith.SubtractFloat(a float64, b float64) returns float64


arith.DivideFloat(a float64, b float64) returns float64


arith.MultiplyFloat(a float64, b float64) returns float64


arith.SqrtFloat(a float64) returns float64


arith.ExpFloat(a float64, b float64) returns float64


arith.AddN(nA string, nB string, base int) returns string


arith.SubtractN(nA string, nB string, base int) returns string


arith.MultiplyN(nA string, nB string, base int) returns string


arith.DivideN(nA string, nB string, base int) returns string


arith.SquareRootN(nA string, base int) returns string


arith.ExpN(nA string, nB string, base int) returns string


arith.NToN(nA string, from int, to int) returns string
~~~


To use library just add it to your src folder and import it using proper Go syntax.

Only steps i took to make this reusable is to put it in its own go file, allowed it to take strings and tried to make the functions as generic as possible.