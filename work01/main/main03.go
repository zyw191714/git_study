package main

import "fmt"

func main() {
	a := []string{"I", "am", "stupid", "and", "weak"}
	for i, value := range a {
		fmt.Printf("key=%v;value=%v\n", i, value)
	}
}
