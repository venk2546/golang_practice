package main

import "fmt"

func main() {
	num := 20
	teacher := false
	student := true
	var x bool
	a := true
	rating := 5
	var movie string

	if student {
		fmt.Println("he is a student ")
	}

	if !student {
		fmt.Println("he is teacher")
		if teacher {
			fmt.Println("he is a teacher ")
		}
		if !teacher {
			fmt.Println("he is a student ")
		}
	}
	if x {
		fmt.Println("this is false")
	}
	if !x {
		fmt.Println("this is a true")
	}

	if name := "venkat"; a {
		fmt.Println("this is ", name)
	}
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
		if rating > 5 {
			movie = "hit"
		} else if rating < 5 {
			movie = "flop"
		} else {
			movie = "under rated"
		}

		fmt.Println(movie)
	}

}
