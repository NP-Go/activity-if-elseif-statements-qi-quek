package main

import (
	"fmt"
)

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }

	var testNum int

	fmt.Printf("Input Number:")
	fmt.Scan(&testNum)

	if testNum%2 == 0 {
		fmt.Println(testNum, " is an even number")
		if testNum <= 9 {
			fmt.Printf("%v is single digit\n", testNum)
		} else {
			fmt.Printf("%v is double digit\n", testNum)
		}
	} else {
		fmt.Printf("%v is an odd number\n", testNum)
		if testNum <= 9 {
			fmt.Printf("%v is single digit\n", testNum)
		} else {
			fmt.Printf("%v is double digit\n", testNum)
		}

	}

}
