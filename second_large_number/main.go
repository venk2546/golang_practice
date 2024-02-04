package main

import "fmt"

func main() {
	var value1 int = 0
	var value2 int = 0
	arr := [3]int{4, 5, 6}

	value1 = arr[0]
	for i := 1; i <= 2; i++ {
		if value1 < arr[i] {
			value2 = value1
			value1 = arr[i]
		} else if value2 < arr[i] {
			value2 = arr[i]
		}
	}
	fmt.Println("Second largest element is: ", value2)
}
