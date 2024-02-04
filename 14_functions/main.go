package main

import (
	"fmt"
)

func total(principal, interest, years int) int {

	si := (principal * interest * years) / 100

	return si
}

func main() {

	var principal, interest, years, si int

	principal = 150000
	interest = 18
	years = 2

	si = total(principal, interest, years)
	fmt.Println("Total amount to be paid:", si)
}
