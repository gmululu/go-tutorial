package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter a string : ")

	scanner := bufio.NewScanner(os.Stdin)
	// Scans a line from Stdin(Console)
	scanner.Scan()
	// Holds the string that scanned
	s := scanner.Text()

	// breaks at space
	//_, err := fmt.Scanf("%s", &s)

	s = strings.ToLower(s)
	if 'i' == s[0] && 'n' == s[len(s)-1] && strings.Contains(s, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
