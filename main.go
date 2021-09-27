package main

import (
	"fmt"
	"germandefrancesco/result"
)

func main() {
	success := false
	fmt.Println("In order to format correctly, you need to pass: 2 digits as type, 2 digits as value length and NN digits as value (it will take the length you previously specified)")
	for success == false {
		fmt.Print("Enter text to format (e.g. NN03ASD): ")
		var input string
		fmt.Scan(&input)
		r, err := result.NewResult(input)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			success = true
			fmt.Println("Your result is:", r)
		}
	}
}
