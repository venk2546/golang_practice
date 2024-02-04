package main

import "fmt"

// template, class
type contact struct {
	greeting string
	id       int
	name     string
}

func swithontype(x interface{}) {
	switch x.(type) { // this is assert, asserting
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float")
		//fallthrough
	case contact:
		fmt.Println("Our Own structure contact")
		//fallthrough
	default:
		fmt.Println("Unknown Type")
	}

}
func main() {

	swithontype(10)
	swithontype("HaiVenkat")
	swithontype(2.2)
	var t = contact{"Good to see you", 1234, "Kumar"}
	swithontype(t)
	swithontype(t.greeting)
	fmt.Println("structure member...:", t.greeting, t.id, t.name)
	swithontype(t.id)

}
