package main

import (
	"fmt"

	addr "github.com/venk2546/modules/address"
	total "github.com/venk2546/modules/loans"
)

var employee = "venkat"

func main() {

	var principal , rateOfInterest , time float64

	principal = 20000
	rateOfInterest = 18%
	time = 12



	fmt.Println(employee)
	fmt.Println(Bank)
	Emp_details("personal banker")
	fmt.Println(addr.Address_details)
	fmt.Println("total amount:",total.Loan_amount)
}
