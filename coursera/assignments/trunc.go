package main

import (
	"fmt"
)

func main() {

	var y float32
	fmt.Println("Enter a floating number :")
	_, err := fmt.Scan(&y)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(int(y))
}
