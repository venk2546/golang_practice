package main

import "fmt"

func main() {
	x := 20
	y := 30
	z := &x
	add := x + y
	sub := x - y
	div := x / y
	mul := x * y
	mod := x % y
	eql := x == y
	neqll := x != y
	great := x < y
	less := x > y
	greateq := x <= y
	lesseq := x >= y
	bitwise1 := x & y
	bitwise2 := x | y
	bitwise3 := x ^ y
	bitwise4 := x << y
	bitwise5 := x >> y
	bitwise6 := x &^ y

	if (x != y) && (x >= y) {
		fmt.Println("True")
	}
	if (x == y) || (x >= y) {
		fmt.Println("True")
	}
	if !(x == y) {
		fmt.Println("True")
	}

	fmt.Printf("Addtitional value is : %v\n", add)
	fmt.Printf("Subtraction value is : %v\n", sub)
	fmt.Printf("Division value is : %v\n", div)
	fmt.Printf("Multiplied value is : %v\n", mul)
	fmt.Printf("Modulus value is : %v\n", mod)
	fmt.Printf("result is : %t\n", eql)
	fmt.Printf("Result is : %t\n", neqll)
	fmt.Printf("Result is : %t\n", great)
	fmt.Printf("Result is : %t\n", less)
	fmt.Printf("Result is : %t\n", greateq)
	fmt.Printf("Result is : %t\n", lesseq)
	fmt.Printf("\nResult of x & y = %d", bitwise1)
	fmt.Printf("\nResult of x | y = %d", bitwise2)
	fmt.Printf("\nResult of x ^ y = %d", bitwise3)
	fmt.Printf("\nResult of x << y = %d", bitwise4)
	fmt.Printf("\nResult of x >> y = %d", bitwise5)
	fmt.Printf("\nResult of x &^ y = %d", bitwise6)
	fmt.Println(*z)
	*z = 7
	fmt.Println(x)
}
