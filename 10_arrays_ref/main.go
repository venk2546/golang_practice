package main

import "fmt"

func main() {
	a := [...]string{"venkatesh", "satish", "sai"} //lighiting operator

	b := a
	a[2] = "kumar"
	fmt.Println("array a output ..:", a)
	fmt.Println("array b output ..:", b)

	refb := &a // copy of array by reference
	a[1] = "venkat"
	refb[2] = "venky"
	fmt.Println("array b output ..:", b)

	x := a[:len(a)-1] //doubt
	fmt.Println(x)

	y := a[2:]
	fmt.Println(y)

	z := a[:] //total array copy
	fmt.Print(z)

	m := a[1:2]
	fmt.Println(m)

	//names := printArray(a)

	//fmt.Println(names)

	//var s []int = make([]int, 3, 5) //slice another way of declaration
	var s []int = []int{1, 2, 3}
	var q []int = s
	q[0] = 10
	fmt.Println("Values of S: ", s)
	fmt.Println("Values of Q: ", q)

	var g = []int{1, 2, 3, 4}
	var h = append(g, 5, 6, 7)
	//fmt.Println(g, h)

	var i = append(h, g...)
	fmt.Println("values in i : ", i)

}
