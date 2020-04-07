package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main()  {
	p := make(map[string] string)
	fmt.Println("Enter name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p["name"] = scanner.Text()
	fmt.Println("Enter address:")
	scanner.Scan()
	p["address"] = scanner.Text()
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	print(string(b))

}
