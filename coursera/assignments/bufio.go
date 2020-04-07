package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 3)

	i := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nEnter an integer to be added to the slice: ")
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		if strings.ToLower(scanner.Text()) == "x" {
			break
		}
		nl := i + 1
		if len(sli) == nl {
			sli = append(sli, nl)
		}
		sli[i] = v
		sort.Ints(sli)
		fmt.Printf("\n%d", sli)
		i++
		fmt.Println("\nEnter another integer to be added to the slice or X to exit: ")
	}

}
