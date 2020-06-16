package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in []byte
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter: ")
	if scanner.Scan() {
		in = []byte(scanner.Text())
	}
	//fmt.Printf("Expression: %s\n", in)
	//in = []byte("(\\abc.b(abc))(\\sz.z)")
	in = []byte("\\bc.b((\\sz.z)bc)")
	b := resolve(in)
	fmt.Printf("%s\n", b)
}

// (\a.ab)(c) = (c)b
// ((\ab.ab)(c))(d) = (\b.cb)(d) = cd
func resolve(expr []byte) []byte {
	// find where the function ends
	rp, lp, endInd := 0, 0, 0
	if expr[0] != '(' {
		return expr
	}
	for i, v := range expr {
		if v == '(' {
			rp++
		} else if v == ')' {
			lp++
		}
		if rp == lp {
			endInd = i
			break
		}
	}
	var argument []byte
	function := resolve(expr[1:endInd])

	if endInd != len(expr)-1 {
		argument = append([]byte{'('}, append(resolve(expr[endInd+1:]), ')')...)

		fmt.Println(string(argument))
		// find the \ and the argument after
		lambdaPos, i := 0, 0
		for function[lambdaPos] != '\\' { // LATER: lambdapos should be 1, but just to be sure
			lambdaPos++
		}
		a := function[lambdaPos+1]
		i = lambdaPos + 2
		for i < len(function) && function[i] != a { // find where the var is in body
			i++
		}
		fmt.Println(string(function), i)
		fmt.Println(i)
		if i == len(function) {
			function = append(function[:lambdaPos+1], function[lambdaPos+2:]...)
		} else {
			cop := make([]byte, len(function[i+1:]))
			copy(cop, function[i+1:])
			function = append(append(function[:lambdaPos+1], append(function[lambdaPos+2:i], argument...)...), cop...)
			fmt.Println(string(function[i+1:]))
		}

		if function[lambdaPos+1] == '.' {
			function = function[lambdaPos+2:]
		}

	}
	return function
}
