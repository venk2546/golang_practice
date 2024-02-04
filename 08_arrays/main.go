package main

import "fmt"

func main() {
	var account [4]string

	account[0] = "venkat"
	account[1] = "satish"
	account[2] = "kumar"
	account[3] = "sai"

	loans := [...]string{"personal loan", "auto loan", "jumbo loan", "insta loan", "bike loan"}

	//loan_account := [...]string{}

	fmt.Println("list of loans: \n", loans, len(loans))
	fmt.Println("list of all accounts", account)
	fmt.Println("list of all account", len(account), account[2])

	//for _, v := range loans {
	//fmt.Println("vegetable in new line :  \n", v)
	for x := 0; x < len(loans); x++ {
		fmt.Println("loans in new line\n", loans[x])
		//if x <= 3 {
		//	loan_account[x] = loans[x]
		//}

	}
}
