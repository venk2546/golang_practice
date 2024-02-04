package main

import "fmt"

const num = 3
const name = "venkat"

const (
	name1 = "satish"
	age   = 30
)

const (
	c3 = iota
	c4
	c5
	c6
)

func main() {

	fmt.Println("the number is:", num)
	fmt.Println("the name is:", name)
	fmt.Println(name1, age)
	fmt.Println(c3, c4, c5, c6)
}
