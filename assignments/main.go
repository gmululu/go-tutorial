package main

import "fmt"

type nslice []int

func createSlice() nslice {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return numbers
}

func (n nslice) checkEvenOrOdd() {
	for _, num := range n {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
}

func main() {
	n := createSlice()
	n.checkEvenOrOdd()
}
