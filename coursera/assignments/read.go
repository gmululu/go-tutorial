package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main()  {
	sli := make([] name,0)
	fmt.Println("Enter full path of file, name and the extension. Eg 'D:/Trainning/go/coursera/assignments/names.txt': ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	file := scanner.Text()
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		fline := fscanner.Text()
		names := strings.Fields(fline)
		sli = append(sli, name{
			fname:names[0],
			lname:names[1],
		})
	}
	for _, n := range sli {
		fmt.Printf("First Name: %s Last Name : %s\n", n.fname, n.lname)
	}
}
