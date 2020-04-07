package main

import (
	"fmt"
	"strings"
)

func main() {
	var y string
	var l string
	fmt.Println("Enter quoted string eg. : ")
	_, err := fmt.Scanf("%q", &y)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	l = strings.ToLower(y)
	if strings.HasPrefix(l, "i") && strings.Contains(l, "a") && strings.HasSuffix(l, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
