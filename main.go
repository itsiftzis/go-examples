package main

import (
	"fmt"
)

func main() {
	checker := 5
	for {
		fmt.Printf("\nUer name: ")
		var userName string
		scan, err := fmt.Scan(&userName)
		if err != nil {
			fmt.Printf("Result: %d, Error:%f\n", scan, err)
			return
		}
		var stores []int
		fmt.Print(userName)
		checker--
		stores = append(stores, checker)
		fmt.Println(stores)
		if checker == 0 {
			break
		}
	}
}
