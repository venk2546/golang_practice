package main

import "fmt"

//func main() {
//	fmt.Println("..loops..")

//	for i := 0; i <= 100; i++ {
//		if i%5 == 0 {
//			fmt.Println(i)
//		}

//	}
//}
func main() {
	i := 1
	for i < 100 {
		if i%5 == 0 {
			fmt.Println(i)
			break
		}
		i++
	}
}
