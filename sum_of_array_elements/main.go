package main

import "fmt"

func main() {
	Arr := [3]int{4, 5, 6}
	Sum := 0
	for i := 0; i < len(Arr); i++ {
		Sum = Sum + Arr[i]
	}
	fmt.Println("The sum of array items = ", Sum)
}
